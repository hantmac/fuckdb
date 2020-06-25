package md

import (
	"fmt"
	"fuckdb/bases/model"
	"fuckdb/view"
	"io"

	"github.com/olekukonko/tablewriter"
)

func init() {
	view.Register(Name, NewView())
}

const (
	// Name 视图名称
	Name = "md"
)

// View Markdown视图
type View struct {
}

// NewView 返回Markdown视图实例
func NewView() view.Viewer {
	return new(View)
}

// Do 将数据库元数据以Markdown视图形式输出。
func (v *View) Do(items []model.DB, out io.Writer) error {
	for i := range items {
		// v.renderDB(&items[i], out)
		// fmt.Fprintln(out)

		for j := range items[i].Tables {
			fmt.Fprintf(out, "---\ntitle: %s - %s\nweight: 100\n---\n\n",
				items[i].Tables[j].Name,
				items[i].Tables[j].Comment,
			)
			fmt.Fprint(out, "## 结构\n\n")
			v.renderTable(&items[i].Tables[j], out)
			fmt.Fprintln(out)
			fmt.Fprint(out, "## SQL\n\n```sql\n"+items[i].Tables[j].CreateSQL+"\n```\n\n")
			fmt.Fprint(out, "## updateSQL\n\n```sql\n```\n\n")
		}
	}

	return nil
}

func (v *View) renderDB(db *model.DB, out io.Writer) {
	rows := [][]string{[]string{db.Name, db.CharSet, db.Collation}}

	t := tablewriter.NewWriter(out)
	t.SetHeader([]string{"数据库", "字符集", "比对方法"})
	t.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	t.SetCenterSeparator("|")
	t.AppendBulk(rows)
	t.Render()
}

func (v *View) renderTable(table *model.Table, out io.Writer) {
	rows := make([][]string, 0, len(table.Columns))
	for i := range table.Columns {
		rows = append(rows, []string{
			table.Columns[i].Name,
			table.Columns[i].ColumnType,
			// table.Columns[i].Nullable,
			table.Columns[i].Key,
			table.Columns[i].Default,
			// table.Columns[i].CharSet,
			// table.Columns[i].Collation,
			table.Columns[i].Comment,
		})
	}

	t := tablewriter.NewWriter(out)
	t.SetHeader([]string{"字段", "数据类型", "索引", "默认值", "备注"})
	t.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	t.SetCenterSeparator("|")
	t.AppendBulk(rows)
	t.Render()
}
