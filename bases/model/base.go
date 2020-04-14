package model

import (
	"errors"
	"fmt"

	"database/sql"
)

// Repo MySQL的model.IRepo接口实现
type Repo struct {
	db *sql.DB
}

var (
	// ErrDBNotFound 数据库不存在
	ErrDBNotFound = errors.New("db not found")
)

// DB 数据库实例元信息
type DB struct {
	Name      string            `json:"name,omitempty"`
	CharSet   string            `json:"charset,omitempty"`
	Collation string            `json:"collation,omitempty"`
	Tables    []Table           `json:"tables,omitempty"`
	Extra     map[string]string `json:"extra,omitempty"`
}

// Table 表元信息
type Table struct {
	DB        string            `json:"-"`
	Name      string            `json:"name,omitempty"`
	Collation string            `json:"collation,omitempty"`
	Comment   string            `json:"comment,omitempty"`
	Columns   []Column          `json:"columns,omitempty"`
	Extra     map[string]string `json:"extra,omitempty"`
	CreateSQL string            `json:"create_sql"`
}

// Column 列元信息
type Column struct {
	DB         string            `json:"-"`
	Table      string            `json:"-"`
	Name       string            `json:"name,omitempty"`
	Default    string            `json:"default,omitempty"`
	Nullable   string            `json:"nullable,omitempty"`
	DataType   string            `json:"data_type,omitempty"`
	ColumnType string            `json:"column_type,omitempty"`
	Key        string            `json:"key,omitempty"`
	CharSet    string            `json:"charset,omitempty"`
	Collation  string            `json:"collation,omitempty"`
	Comment    string            `json:"comment,omitempty"`
	Extra      map[string]string `json:"extra,omitempty"`
}

// IRepo 数据库元信息查询接口
type IRepo interface {
	// GetDBs 查询数据库元信息
	GetDBs(cond *DB, lazy bool) ([]DB, error)
	// GetTables 查询表元信息
	GetTables(cond *Table) ([]Table, error)
	// GetColumns 查询列元信息
	GetColumns(cond *Column) ([]Column, error)

	GetCreateTableSQL(tableName string) (string, error)
}

// MysqlConfig mysql 配置
type MysqlConfig struct {
	Host     string
	Port     int
	Socket   string
	Username string
	Password string
	DB       string
	Table    string
	Viewer   string
	Output   string
	Debug    bool
}

// GenDataSource generator data connect source
func GenDataSource(c *MysqlConfig, params string) (dataSource string) {
	if c.Socket == "" {
		// use tcp protocol
		if c.Password == "" {
			return fmt.Sprintf("%s@tcp(%s:%d)/%s?%s", c.Username, c.Host, c.Port, c.DB, params)
		}
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.Username, c.Password, c.Host, c.Port, c.DB, params)
	}
	// use unix domain socket protocol
	if c.Password == "" {
		return fmt.Sprintf("%s@unix(%s)/%s?%s", c.Username, c.Socket, c.DB, params)
	}
	return fmt.Sprintf("%s:%s@unix(%s)/%s?%s", c.Username, c.Password, c.Socket, c.DB, params)
}

// NewRepo 实例化
func NewRepo(c *MysqlConfig) (IRepo, error) {
	db, err := sql.Open("mysql", GenDataSource(c, "charset=utf8mb4&parseTime=true&loc=Local"))
	// Check for error in db, note this does not check connectivity but does check uri
	if err != nil {
		fmt.Println("Error opening mysql db: " + err.Error())
		return nil, err
	}
	return &Repo{
		db: db,
	}, nil
}

// =========================================

type schema struct {
	Name      string `json:"SCHEMA_NAME"`
	CharSet   string `json:"DEFAULT_CHARACTER_SET_NAME"`
	Collation string `json:"DEFAULT_COLLATION_NAME"`
}

// getSchemas by schema_name
func (repo *Repo) getSchemas(cond *schema) (items []schema, err error) {
	query := "SELECT DEFAULT_CHARACTER_SET_NAME,DEFAULT_COLLATION_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = (?) "

	rows, err := repo.db.Query(query, cond.Name)
	defer rows.Close()

	if err != nil {
		fmt.Println("Error selecting from schema: " + err.Error())
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("No results returned for Schema ")
	}

	for rows.Next() {
		var charSet string
		var collection string
		rows.Scan(&charSet, &collection)
		items = append(items, schema{
			Name:      cond.Name,
			CharSet:   charSet,
			Collation: collection,
		})
	}

	// fmt.Println("Schema==================")
	// fmt.Println(items)
	return items, nil
}

type table struct {
	Schema    string `json:"TABLE_SCHEMA"`
	Name      string `json:"TABLE_NAME"`
	Collation string `json:"TABLE_COLLATION"`
	Comment   string `json:"TABLE_COMMENT"`
}

// getTables by table_schema
func (repo *Repo) getTables(cond *table) (items []table, err error) {
	query := "SELECT TABLE_SCHEMA, TABLE_NAME, TABLE_COLLATION, TABLE_COMMENT FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = (?) AND TABLE_NAME = (?)"

	rows, err := repo.db.Query(query, cond.Schema, cond.Name)
	defer rows.Close()

	if err != nil {
		fmt.Println("Error selecting from Tables: " + err.Error())
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("No results returned for Tables ")
	}

	for rows.Next() {
		var schema string
		var name string
		var collection string
		var comment string
		rows.Scan(&schema, &name, &collection, &comment)
		items = append(items, table{
			Schema:    cond.Schema,
			Name:      name,
			Comment:   comment,
			Collation: collection,
		})
	}

	// fmt.Println("Table==================")
	// fmt.Println(items)
	return items, nil
}

type column struct {
	Schema     string `json:"TABLE_SCHEMA"`
	Table      string `json:"TABLE_NAME"`
	Name       string `json:"COLUMN_NAME"`
	Default    string `json:"COLUMN_DEFAULT"`
	Nullable   string `json:"IS_NULLABLE"`
	DataType   string `json:"DATA_TYPE"`
	ColumnType string `json:"COLUMN_TYPE"`
	Key        string `json:"COLUMN_KEY"`
	CharSet    string `json:"CHARACTER_SET_NAME"`
	Collation  string `json:"COLLATION_NAME"`
	Comment    string `json:"COLUMN_COMMENT"`
}

func (repo *Repo) getColumns(cond *column) (items []column, err error) {
	query := "SELECT COLUMN_NAME, COLUMN_KEY, DATA_TYPE, COLUMN_TYPE, IS_NULLABLE, COLUMN_COMMENT, COLUMN_DEFAULT, CHARACTER_SET_NAME, COLLATION_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = (?) AND TABLE_NAME = (?)"

	rows, err := repo.db.Query(query, cond.Schema, cond.Table)
	defer rows.Close()

	if err != nil {
		fmt.Println("Error selecting from db: " + err.Error())
		return nil, err
	}

	if rows == nil {
		return nil, errors.New("No results returned for table ")
	}

	for rows.Next() {
		var columnName string
		var columnKey string
		var dataType string
		var columnType string
		var nullable string
		var defaultVal string
		var comment string
		var charSet string
		var collection string

		rows.Scan(&columnName, &columnKey, &dataType, &columnType, &nullable, &comment, &defaultVal, &charSet, &collection)

		items = append(items, column{
			Schema:     cond.Schema,
			Table:      cond.Table,
			Name:       columnName,
			Key:        columnKey,
			DataType:   dataType,
			ColumnType: columnType,
			Nullable:   nullable,
			Default:    defaultVal,
			Comment:    comment,
			CharSet:    charSet,
			Collation:  collection,
		})

	}

	// fmt.Println("Column==================")
	// fmt.Printf("%#v", items)
	return items, nil
}

// =========================================

// GetDBs 按条件查询数据库信息
func (repo *Repo) GetDBs(cond *DB, lazy bool) (items []DB, err error) {
	var sCond schema
	if cond != nil {
		sCond.Name = cond.Name
		sCond.CharSet = cond.CharSet
		sCond.Collation = cond.Collation
	}
	schemas, err := repo.getSchemas(&sCond)

	if err != nil {
		return nil, err
	}

	if len(schemas) <= 0 {
		return nil, ErrDBNotFound
	}

	for i := range schemas {
		var tables []Table
		if !lazy {
			tables, err = repo.GetTables(&Table{
				DB: schemas[i].Name,
			})
			if err != nil {
				return nil, err
			}
		}
		items = append(items, DB{
			Name:      schemas[i].Name,
			CharSet:   schemas[i].CharSet,
			Collation: schemas[i].Collation,
			Tables:    tables,
		})
	}

	return items, nil
}

// GetTables 按条件查询表信息
func (repo *Repo) GetTables(cond *Table) (items []Table, err error) {
	var tCond table

	if cond != nil {
		tCond.Schema = cond.DB
		tCond.Name = cond.Name
		tCond.Collation = cond.Collation
		tCond.Comment = cond.Comment
	}

	tables, err := repo.getTables(&tCond)
	if err != nil {
		return nil, err
	}

	for i := range tables {
		cols, err := repo.GetColumns(&Column{
			DB:    tables[i].Schema,
			Table: tables[i].Name,
		})
		if err != nil {
			return nil, err
		}
		createSQL, _ := repo.GetCreateTableSQL(tables[i].Name)
		items = append(items, Table{
			DB:        tables[i].Schema,
			Name:      tables[i].Name,
			Collation: tables[i].Collation,
			Comment:   tables[i].Comment,
			Columns:   cols,
			CreateSQL: createSQL,
		})
	}
	return items, nil
}

// GetColumns 按条件查询列信息
func (repo *Repo) GetColumns(cond *Column) (items []Column, err error) {
	var cCond column
	if cond != nil {
		cCond.Schema = cond.DB
		cCond.Table = cond.Table
		cCond.Name = cond.Name
		cCond.Default = cond.Default
		cCond.Nullable = cond.Nullable
		cCond.CharSet = cond.CharSet
		cCond.Collation = cond.Collation
		cCond.DataType = cond.DataType
		cCond.ColumnType = cond.ColumnType
		cCond.Key = cond.Key
		cCond.Comment = cond.Comment
	}
	cols, err := repo.getColumns(&cCond)
	if err != nil {
		return nil, err
	}

	for i := range cols {
		items = append(items, Column{
			DB:         cols[i].Schema,
			Table:      cols[i].Table,
			Name:       cols[i].Name,
			Default:    cols[i].Default,
			Nullable:   cols[i].Nullable,
			DataType:   cols[i].DataType,
			ColumnType: cols[i].ColumnType,
			Key:        cols[i].Key,
			CharSet:    cols[i].CharSet,
			Collation:  cols[i].Collation,
			Comment:    cols[i].Comment,
		})
	}
	return items, nil
}

// GetCreateTableSQL by table_name
func (repo *Repo) GetCreateTableSQL(tableName string) (sql string, err error) {
	query := "SHOW CREATE TABLE `" + tableName + "`"
	rows, err := repo.db.Query(query)
	defer rows.Close()

	if err != nil {
		fmt.Println("Error show create table sql: " + err.Error())
		return
	}

	if rows == nil {
		return "", errors.New("No results returned for Tables ")
	}

	for rows.Next() {
		var table string
		var createTable string

		rows.Scan(&table, &createTable)
		sql = createTable
	}
	return
}
