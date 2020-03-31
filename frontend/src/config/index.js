let APIdb2struct
if(process.env.NODE_ENV === "development"){
  APIdb2struct = "http://0.0.0.0:8000"
}else{
  APIdb2struct = "http://0.0.0.0:8000"
}

export default {
  APIdb2struct
}
