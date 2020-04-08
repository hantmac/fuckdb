package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MysqlInfo is mysql的元数据信息
type MysqlInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Addr     string `json:"addr"`
	DbName   string `json:"db_name"`
}

//FormatMysqlStr is %s:%s@tcp(%s)/%s?charset=utf8
func FormatMysqlStr(UserName, Password, Addr, DbName string) string {
	mysqlStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", UserName, Password, Addr, DbName)
	return mysqlStr
}

// FormatMysqlCmdStr generate mysql -h xx -u xx -p xx db
func FormatMysqlCmdStr(UserName, Password, Addr, DbName string) string {
	mysqlCmdStr := fmt.Sprintf("mysql -h %s -u %s -p %s %s", Addr, UserName, Password, DbName)

	return mysqlCmdStr
}

// FormatMysql to format sql str
func FormatMysql(c *gin.Context) {
	sqlStrMap := map[string]string{}
	var mysqlInfo MysqlInfo
	err := c.ShouldBindJSON(&mysqlInfo)
	if err != nil {
		log.Fatal("Get req data failed")
		HandleError(http.StatusBadRequest, c, err)
		return
	}

	mysqlStr := FormatMysqlStr(mysqlInfo.UserName, mysqlInfo.Password, mysqlInfo.Addr, mysqlInfo.DbName)
	mysqlCmdStr := FormatMysqlCmdStr(mysqlInfo.UserName, mysqlInfo.Password, mysqlInfo.Addr, mysqlInfo.DbName)
	sqlStrMap["mysqlStr"] = mysqlStr
	sqlStrMap["mysqlCmdStr"] = mysqlCmdStr

	c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "Get sql str success",
		Data:    sqlStrMap,
	})
}
