package routers

import (
	"errors"
	"fmt"
	"fuckdb/bases"
	"fuckdb/services"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"net/http"
)

type MysqlInfoReqData struct {
	MysqlHost       string `json:"mysql_host"`
	MysqlPort       int    `json:"mysql_port"`
	MysqlDB         string `json:"mysql_db"`
	MysqlTable      string `json:"mysql_table"`
	MysqlPasswd     string `json:"mysql_passwd"`
	MysqlUser       string `json:"mysql_user"`
	PackageName     string `json:"package_name"`
	StructName      string `json:"struct_name"`
	XmlAnnotation   bool   `json:"xml_annotation"`
	JsonAnnotation  bool   `json:"json_annotation"`
	GormAnnotation  bool   `json:"gorm_annotation"`
	XormAnnotation  bool   `json:"xorm_annotation"`
	FakerAnnotation bool   `json:"faker_annotation"`
	GureGuTypes     bool   `json:"gure_gu_types"`
}

func DbToGoStruct(c *gin.Context) {
	var mysqlInfo MysqlInfoReqData
	err := c.ShouldBindJSON(&mysqlInfo)
	if err != nil {
		services.HandleError(http.StatusBadRequest, c, err)
		log.Error("Bind Err", err)
		return
	}

	if mysqlInfo.MysqlHost == "" || mysqlInfo.MysqlDB == "" || mysqlInfo.MysqlTable == "" {
		services.HandleError(http.StatusBadRequest, c, errors.New("Need Mysql Info "))
		return
	}
	if mysqlInfo.MysqlPort == 0 {
		mysqlInfo.MysqlPort = 3306
	}
	columnDataTypes, err := bases.GetColumnsFromMysqlTable(mysqlInfo.MysqlUser, mysqlInfo.MysqlPasswd,
		mysqlInfo.MysqlHost, mysqlInfo.MysqlPort, mysqlInfo.MysqlDB, mysqlInfo.MysqlTable)
	if err != nil {
		fmt.Println("Error in selecting column data information from mysql information schema")
		services.HandleError(http.StatusInternalServerError, c, err)
		return
	}

	if mysqlInfo.StructName == "" {
		mysqlInfo.StructName = "MyNewStruct"
	}
	if mysqlInfo.PackageName == "" {
		mysqlInfo.PackageName = "my_new_package"
	}

	structInfo, err := bases.Generate(*columnDataTypes, mysqlInfo.MysqlTable, mysqlInfo.StructName,
		mysqlInfo.PackageName, mysqlInfo.JsonAnnotation, mysqlInfo.GormAnnotation,
		mysqlInfo.XmlAnnotation, mysqlInfo.XormAnnotation, mysqlInfo.FakerAnnotation, mysqlInfo.GureGuTypes)

	if err != nil {
		fmt.Println("Error in creating struct from json: " + err.Error())
		services.HandleError(http.StatusInternalServerError, c, err)
		return
	}

	fmt.Println(string(structInfo))
	c.JSON(http.StatusOK, services.Response{
		Status:  "0",
		Message: "Ok",
		Data:    string(structInfo),
	})
}
