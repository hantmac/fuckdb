package json

import (
	"encoding/json"
	"fuckdb/bases/model"
	"fuckdb/view"
	"io"
)

func init() {
	view.Register(Name, NewView())
}

const (
	// Name 视图名称
	Name = "json"
)

// View JSON视图
type View struct {
}

// NewView 返回JSON视图实例
func NewView() view.Viewer {
	return new(View)
}

// Do 将数据库元数据以JSON视图形式输出。
func (v *View) Do(items []model.DB, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("  ", "  ")
	return enc.Encode(items)
}
