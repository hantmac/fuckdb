package bases

type Table struct {
	Name    string
	Comment string
	Columns []Column
}

type Column struct {
	Name       string
	Comment    string
	Type       string
	Size       int
	Default    interface{}
	NotNull    bool
	AutoInc    bool
	Unique     bool
	PrimaryKey bool
}
