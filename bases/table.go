package bases

type Table struct {
	Name    string   `json:"name"`
	Comment string   `json:"comment"`
	Columns []Column `json:"columns,omitempty"`
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
