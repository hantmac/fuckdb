package main

import (
	"fmt"
	"fuckdb/routers/middleware"
	"fuckdb/config"
	"fuckdb/routers"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"io/ioutil"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	g := gin.New()
	// LoggerWithFormatter middleware
	// By default gin.DefaultWriter = os.Stdout
	g.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		b, _ := ioutil.ReadAll(param.Request.Body)
		// your custom format
		return fmt.Sprintf("StatusCode:%d--%s-ClientIp:%s - TimeStamp:[%s] \"ReqMethod:%s--"+
			"API:%s--ReqProto%s--CostTime:%s \"UserAgent:%s-\" -Error:%s\"\n",
			param.StatusCode,
			string(b),
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	//Add middleware in this slice
	var middlewares = []gin.HandlerFunc{middleware.NoCache, middleware.Secure}

	routers.Load(
		g,
		middlewares...,
	)

	// init config
	if err := config.InitConfig(""); err != nil {
		log.Error("init config error:%s", err)
		panic(err)
	}
	log.Info("config init success")

	// run server
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	if err := g.Run(host + ":" + port); err != nil {
		log.Error("run server error:", err)
		panic(err)
	}

}
