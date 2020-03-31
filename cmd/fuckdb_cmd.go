package main

import (
	"encoding/json"
	"fmt"
	"fuckdb/bases"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

type DBConf struct {
	Db struct {
		Host           string `json:"host"`
		Port           int    `json:"port"`
		Password       string `json:"password"`
		User           string `json:"user"`
		Table          string `json:"table"`
		Database       string `json:"database"`
		PackageName    string `json:"packageName"`
		StructName     string `json:"structName"`
		JSONAnnotation bool   `json:"jsonAnnotation"`
		GormAnnotation bool   `json:"gormAnnotation"`
	} `json:"db"`
}

func GetDbConf() (*DBConf, error) {
	mysqlInfo := DBConf{}
	jsonFile, err := ioutil.ReadFile("./fuckdb.json")
	if err != nil {
		log.Fatal(err)
		return &mysqlInfo, err
	}

	err = json.Unmarshal(jsonFile, &mysqlInfo)

	if err != nil {
		log.Fatal(err)
		return &mysqlInfo, err
	}

	return &mysqlInfo, nil
}

func main() {

	mysqlInfo, err := GetDbConf()

	if mysqlInfo.Db.Host == "" || mysqlInfo.Db.Database == "" || mysqlInfo.Db.Table == "" {
		log.Fatal("the host, database, table cant be empty")
		return
	}
	if mysqlInfo.Db.Port == 0 {
		mysqlInfo.Db.Port = 3306
	}
	columnDataTypes, err := bases.GetColumnsFromMysqlTable(mysqlInfo.Db.User, mysqlInfo.Db.Password,
		mysqlInfo.Db.Host, mysqlInfo.Db.Port, mysqlInfo.Db.Database, mysqlInfo.Db.Table)
	if err != nil {
		fmt.Println("Error in selecting column data information from mysql information schema")
		log.Fatal("Error in selecting column data information from mysql information schema")
		return
	}

	if mysqlInfo.Db.StructName == "" {
		mysqlInfo.Db.StructName = "MyNewStruct"
	}
	if mysqlInfo.Db.PackageName == "" {
		mysqlInfo.Db.PackageName = "my_new_package"
	}

	structInfo, err := bases.Generate(*columnDataTypes, mysqlInfo.Db.Table, mysqlInfo.Db.StructName,
		mysqlInfo.Db.PackageName, mysqlInfo.Db.JSONAnnotation, mysqlInfo.Db.GormAnnotation,
		false, false, false, true)

	fmt.Println(string(structInfo))
}
