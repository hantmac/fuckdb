import axios from 'axios'


const service = axios.create({
  timeout: 80000
})

service.interceptors.request.use(config => {
  return config;
}, err => {
  return Promise.resolve(err);
})

service.interceptors.response.use(function (config) {
  return config
}, function (error) {
  captureException(error)
  return Promise.reject(error)
})

function captureException(error) {
  return window.$Raven && window.$Raven.captureException(error)
}


export default service
