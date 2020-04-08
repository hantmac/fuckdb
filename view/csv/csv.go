package csv

import (
	"encoding/csv"
	"fmt"
	"fuckdb/bases/model"
	"fuckdb/view"
	"io"
)

func init() {
	view.Register(Name, NewView())
}

const (
	// Name 视图名称
	Name = "csv"
)

// View CSV视图
type View struct {
}

// NewView 返回CSV视图实例
func NewView() view.Viewer {
	return new(View)
}

// Do 将数据库元数据以CSV视图形式输出。
func (v *View) Do(items []model.DB, out io.Writer) (err error) {
	for i := range items {
		if err = v.renderDB(&items[i], out); err != nil {
			return err
		}
		fmt.Fprintf(out, "\n\n\n")
		for j := range items[i].Tables {
			fmt.Fprintf(out, "TABLE: %s,%s\n",
				items[i].Tables[j].Name,
				items[i].Tables[j].Comment,
			)
			if err = v.renderTable(&items[i].Tables[j], out); err != nil {
				return err
			}
			fmt.Fprintf(out, "\n\n\n")
		}
	}
	return nil
}

func (v *View) renderDB(db *model.DB, out io.Writer) error {
	rows := [][]string{
		[]string{"Database", "Character Set", "Collation"},
		[]string{db.Name, db.CharSet, db.Collation},
	}
	return csv.NewWriter(out).WriteAll(rows)
}

func (v *View) renderTable(table *model.Table, out io.Writer) error {
	rows := make([][]string, 0, len(table.Columns)+1)
	rows = append(rows, []string{"Column", "Data Type", "Nullable", "Key", "Default", "Character Set", "Collation", "Comment"})
	for i := range table.Columns {
		rows = append(rows, []string{
			table.Columns[i].Name,
			table.Columns[i].DataType,
			table.Columns[i].Nullable,
			table.Columns[i].Key,
			table.Columns[i].Default,
			table.Columns[i].CharSet,
			table.Columns[i].Collation,
			table.Columns[i].Comment,
		})
	}
	return csv.NewWriter(out).WriteAll(rows)
}
