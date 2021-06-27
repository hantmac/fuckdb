package routers

import (
	"fuckdb/routers/middleware"
	"fuckdb/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	// The health check handlers
	api := g.Group("/api")
	{
		api.GET("/health", services.HealthCheck)
		api.POST("/db2struct", DbToGoStruct)
		api.POST("/testconn",)
		api.POST("/download", DownloadStructFile)
		api.GET("/sql_str", services.FormatMysql)
		api.GET("struct_sql", services.FromStructToSql)
		api.POST("/tables", GetTables)
	}
	return g
}
