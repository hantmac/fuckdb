
# fuckdb
- [中文](https://github.com/hantmac/fuckdb/blob/master/README_zh.md)
- English

![Example to show fuckdb](http://static.vue2.net/db.gif)

`fuckdb` helps you fuck the db when you write go struct.
`fuckdb` generates a go compatible struct type with the required column names, data types, and annotations just fill in the database information in the web UI. Making go web develop very easy by saving a lot of time writing structure.`fuckdb`is based/inspired by the work of Seth Shelnutt's db2struct, and Db2Struct is based/inspired by the work of ChimeraCoder's gojson package gojson.
# Web UI
- Easy to use
Only a few clicks on the web UI can generate the corresponding golang struct with `ORM` or `json` or `xml` ... tags.
# How to use?

### Source code deploy
- git clone the source code
- cd fuckdb/
- make build-frontend
- make start
- Use `localhost:8000` you will get the web UI
- The same operation with docker deployment.

### Docker deploy
- modify your backend ip in  `frontend/src/config/index.js`
- docker-compose up -d
- Use `localhost:8000` you will get the next page
![](https://tva1.sinaimg.cn/large/006tNbRwgy1g9w1ru6tl4j31wb0u0aft.jpg)
- Just Put your mariaDB/mysql info into it and you will get your golang code.
- Click generate.
- Boom！Enjoy your coffee and copy your struct code.
![](https://tva1.sinaimg.cn/large/006tNbRwly1g9w531osobj31u90u0jzq.jpg)

### CMD - fuckdb Lite
#### install
- macos 
 ```
  brew tap hantmac/tap && brew install fuckdb
 ```
- Linux 

 ```
curl https://github.com/hantmac/fuckdb/releases/download/v0.2/fuckdb_linux.tar.gz
  ```
  
- windows 
  download the windows file in release

- First,  run `fuckdb generate` to generate `fuckdb.json` add your mysql info
```
{
  "db": {
    "host": "localhost",
    "port": 3306,
    "password": "password",
    "user": "root",
    "table": "cars",
    "database": "example",
    "packageName": "test",
    "structName": "test",
    "jsonAnnotation": true,
    "gormAnnotation": true
  }
}
```

- then run `fuckdb go` and get your code!

More info => 'Go_Official_Blog' on WeChat:

![](https://user-gold-cdn.xitu.io/2020/3/27/1711c6a37c75e3fb?w=142&h=142&f=jpeg&s=10725)

### Support

[![jetbrains](https://s1.ax1x.com/2020/03/26/G9uQoR.png)]( https://www.jetbrains.com/?from=fuckdb)
