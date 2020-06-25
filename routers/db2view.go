package routers

import (
	"bytes"
	"errors"
	"fmt"
	"fuckdb/bases"
	"fuckdb/bases/model"
	"fuckdb/services"
	"fuckdb/view"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

type MysqlViewReqData struct {
	MysqlHost   string `json:"mysql_host"`
	MysqlPort   int    `json:"mysql_port"`
	MysqlDB     string `json:"mysql_db"`
	MysqlTable  string `json:"mysql_table"`
	MysqlPasswd string `json:"mysql_passwd"`
	MysqlUser   string `json:"mysql_user"`
	ViewType    string `json:"view_type"`
}

var out io.Writer = os.Stdout

// DbToView  DB2view view type:csv/md/json/txt
func DbToView(c *gin.Context) {
	var mysqlInfo MysqlViewReqData
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

	dbs, err := bases.GetMetadata(repo, mysqlInfo.MysqlDB, mysqlInfo.MysqlTable)
	if err != nil {
		fmt.Println("Error in get meta data")
		services.HandleError(http.StatusInternalServerError, c, err)
		return
	}
	// Output as target viewer
	v := view.SelectViewer(mysqlInfo.ViewType)
	if v == nil {
		fmt.Println("unsupported viewer error :" + err.Error())
		return
	}

	buf := new(bytes.Buffer)
	if err = v.Do(dbs, buf); err != nil {
		fmt.Println("dump viewer error :" + err.Error())
		services.HandleError(http.StatusInternalServerError, c, err)
		return
	}

	c.JSON(http.StatusOK, services.Response{
		Status:  "0",
		Message: "Ok",
		Data:    buf.String(),
	})
}
