package routers

import (
	"bufio"
	"fuckdb/services"
	"fuckdb/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path/filepath"
)

type DownLoadReqInfo struct {
	StructName string `json:"struct_name"`
	Content    string `json:"content"`
}

func DownloadStructFile(c *gin.Context) {
	var downLoadInfo DownLoadReqInfo
	err := c.ShouldBindJSON(&downLoadInfo)
	if err != nil {
		services.HandleError(http.StatusBadRequest, c, err)
		return
	}

	if downLoadInfo.StructName == "" {
		downLoadInfo.StructName = "new_struct"
	}
	if downLoadInfo.Content == "" {
		services.HandleError(http.StatusBadRequest, c, err)
		return
	}

	var filename = filepath.Join(viper.GetString("genFilePath"), downLoadInfo.StructName) + ".go"
	if !utils.CheckFileIsExist(filename) {
		file, err := os.Create(filename) //创建文件
		defer file.Close()
		if err != nil {
			services.HandleError(http.StatusInternalServerError, c, err)
			return
		}
		buf := bufio.NewWriter(file) //创建新的 Writer 对象
		buf.WriteString(downLoadInfo.Content)
		buf.Flush()

	}
	//返回文件路径
	c.JSON(http.StatusOK, services.Response{
		Status:  "0",
		Message: "ok",
		Data:    filename,
	})
}
