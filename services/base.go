package services

import (
	"errors"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

type ListResponse struct {
	Response
	Total int `json:"total"`
}

//封装错误处理，避免冗余
func HandleError(statusCode int, c *gin.Context, err error) {
	logrus.Errorln("handle error:", err)
	c.Error(errors.New(string(debug.Stack())))
	debug.PrintStack()
	c.AbortWithStatusJSON(statusCode, Response{
		Status:  "-1",
		Message: "error",
		Error:   err.Error(),
	})
}
