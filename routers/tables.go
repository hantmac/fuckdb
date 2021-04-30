package routers

import (
	"database/sql"
	"errors"
	"fuckdb/bases"
	"fuckdb/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type config struct {
	Host   string `json:"mysql_host"`
	Port   int    `json:"mysql_port"`
	DB     string `json:"mysql_db"`
	Passwd string `json:"mysql_passwd"`
	User   string `json:"mysql_user"`
}

func DetectConnection(c *gin.Context) {
	var cfg config
	if err := c.ShouldBindJSON(&cfg); err != nil {
		services.HandleError(http.StatusBadRequest, c, err)
		logrus.Errorln("Bind Err", err)
		return
	}
	if cfg.Host == "" || cfg.DB == "" {
		services.HandleError(http.StatusBadRequest, c, errors.New("Need Mysql Info "))
		return
	}
	if cfg.Port == 0 {
		cfg.Port = 3306
	}

	db, err := connectToDB(cfg.Host, cfg.Port, cfg.User, cfg.Passwd, cfg.DB)
	if err != nil {
		logrus.Errorln("Error opening mysql db: " + err.Error())
		c.JSON(http.StatusOK, services.Response{
			Status:  "100",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		c.JSON(http.StatusOK, services.Response{
			Status:  "101",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, services.Response{
		Status:  "0",
		Message: "Ok",
		Data:    nil,
	})
}

func GetTables(c *gin.Context) {
	var cfg config
	if err := c.ShouldBindJSON(&cfg); err != nil {
		services.HandleError(http.StatusBadRequest, c, err)
		logrus.Errorln("Bind Err", err)
		return
	}
	if cfg.Host == "" || cfg.DB == "" {
		services.HandleError(http.StatusBadRequest, c, errors.New("Need Mysql Info "))
		return
	}
	if cfg.Port == 0 {
		cfg.Port = 3306
	}

	db, err := connectToDB(cfg.Host, cfg.Port, cfg.User, cfg.Passwd, cfg.DB)
	if err != nil {
		logrus.Errorln("Error opening mysql db: " + err.Error())
		c.JSON(http.StatusOK, services.Response{
			Status:  "100",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT TABLE_NAME,TABLE_COMMENT FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA= ?", cfg.DB)
	if err != nil {
		logrus.Errorln("query tables error:", err)
		c.JSONP(http.StatusInternalServerError, services.Response{
			Status:  "200",
			Message: err.Error(),
		})
		return
	}

	tables := make([]bases.Table, 0)
	for rows.Next() {
		var table bases.Table
		rows.Scan(&table.Name, &table.Comment)
		tables = append(tables, table)
	}

	c.JSONP(http.StatusOK, services.Response{
		Status:  "0",
		Message: "OK",
		Data:    tables,
	})

}

func connectToDB(host string, port int, user string, password string, dbname string) (*sql.DB, error) {
	if password != "" {
		return sql.Open("mysql", user+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/"+dbname+"?&parseTime=True")
	} else {
		return sql.Open("mysql", user+"@tcp("+host+":"+strconv.Itoa(port)+")/"+dbname+"?&parseTime=True")
	}
}
