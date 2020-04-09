package routers

import (
	"errors"
	"fmt"
	"fuckdb/bases"
	"fuckdb/bases/model"
	"fuckdb/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
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
	XMLAnnotation   bool   `json:"xml_annotation"`
	JSONAnnotation  bool   `json:"json_annotation"`
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

	mysqlConfig := &model.MysqlConfig{
		Host:     mysqlInfo.MysqlHost,
		Port:     mysqlInfo.MysqlPort,
		Socket:   "",
		Username: mysqlInfo.MysqlUser,
		Password: mysqlInfo.MysqlPasswd,
		DB:       mysqlInfo.MysqlDB,
		Table:    mysqlInfo.MysqlTable,
	}

	repo, err := model.NewRepo(mysqlConfig)
	if err != nil {
		fmt.Println("Error in create new repo")
		services.HandleError(http.StatusInternalServerError, c, err)
		return
	}

	structInfo, err := bases.GenStructInfo(repo, mysqlInfo.MysqlDB, mysqlInfo.MysqlTable, mysqlInfo.StructName,
		mysqlInfo.PackageName, mysqlInfo.JSONAnnotation, mysqlInfo.GormAnnotation,
		mysqlInfo.XMLAnnotation, mysqlInfo.XormAnnotation, mysqlInfo.FakerAnnotation,
		mysqlInfo.GureGuTypes)

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
