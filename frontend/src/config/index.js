let APIdb2struct
//本地开发
if(process.env.NODE_ENV === "development"){
  APIdb2struct = "http://0.0.0.0:8000"
}else{//生产环境
  APIdb2struct = "http://0.0.0.0:8000"
}

export default {
  APIdb2struct
}
