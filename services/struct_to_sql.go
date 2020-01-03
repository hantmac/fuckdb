package services

import (
	"go/ast"
	"go/parser"
	"go/token"
	"net/http"
	"path/filepath"
	"strings"

	"fuckdb/program"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

type StruStr struct {
	FileName   string `json:"file_name"`
	StructName string `json:"struct_name"`
	StructStr  string `json:"struct_str"`
}

func FromStructToSql(c *gin.Context) {
	var types []*ast.TypeSpec
	fset := token.NewFileSet()
	var structStr StruStr
	err := c.ShouldBindJSON(&structStr)
	if err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	}
	pattern := structStr.StructName
	matchFunc := func(structName string) bool {
		match, _ := filepath.Match(pattern, structName)
		return match
	}
	f, err := parser.ParseFile(fset, structStr.FileName, structStr.StructStr, parser.ParseComments)
	if err != nil {
		log.Infof("parse [file:%s] failed:%v", structStr.FileName, err)

	}

	types = program.FindMatchStruct([]*ast.File{f}, matchFunc)

	sqls := []string{}
	for _, typ := range types {
		ms, err := NewSqlGenerator(typ)
		if err != nil {
			log.Infof("create model struct failed:%v", err)
		}

		sql, err := ms.GetCreateTableSql()
		if err != nil {
			log.Infof("generate sql failed:%v", err)
		}
		sqls = append(sqls, sql)
	}

	resStr := strings.Join(sqls, "\n\n")

	c.JSON(http.StatusOK, Response{
		Status:  "0",
		Message: "Ok",
		Data:    resStr,
	})

}
