FROM golang:1.13 AS backend-build

WORKDIR /go/src/app
COPY . .

ENV GO111MODULE on
ENV GOPROXY https://goproxy.io

RUN go install
RUN make
# copy backend files

FROM node:8.16.0-alpine AS frontend-build

ADD . /app
WORKDIR /app/frontend

# install frontend
RUN npm config set unsafe-perm true
RUN npm install -g yarn && yarn install --registry=https://registry.npm.taobao.org

RUN npm run build

FROM ubuntu:latest
ADD . /app
RUN apt-get update \
	&& apt-get install -y curl

# set as non-interactive
ENV DEBIAN_FRONTEND noninteractive

# install nginx
RUN apt-get -y install nginx
COPY --from=backend-build /go/src/app/fuckdb /usr/local/bin
# copy frontend files
COPY --from=frontend-build /app/frontend/dist /app/frontend/dist
COPY --from=frontend-build /app/frontend/conf/fuckdb.conf /etc/nginx/conf.d
# working directory
WORKDIR /app
EXPOSE 8080
# backend port
EXPOSE 8000

# start backend
CMD ["/bin/sh", "/app/docker_init.sh"]
