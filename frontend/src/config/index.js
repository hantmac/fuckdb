let APIdb2struct
if(process.env.NODE_ENV === "development"){
  APIdb2struct = "localhost:8088"
}else if(process.env.NODE_ENV === "production"){
  APIdb2struct = "localhost:8088"
}else{
  APIdb2struct = "localhost:8088"
}

export default {
  APIdb2struct
}
