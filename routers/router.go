package routers

import (
	"fuckdb/routers/middleware"
	"fuckdb/services"
	"github.com/gin-gonic/gin"
	"net/http"
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
	db2s := g.Group("/api")
	{
		db2s.GET("/health", services.HealthCheck)
		db2s.POST("/db2struct", DbToGoStruct)
		db2s.POST("/download", DownloadStructFile)
		db2s.GET("/sql_str", services.FormatMysql)
	}
	return g
}
