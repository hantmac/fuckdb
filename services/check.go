package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary for health check
// @Description health check
// @Accept  json
// @Produce  json
// @Success 200 {string} string    "ok"
// @Failure 500 {string} string "server is not healthy"
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	message := "healthy"

	c.JSON(http.StatusOK, Response{
		Status:  "0",
		Message: message,
	})

}
