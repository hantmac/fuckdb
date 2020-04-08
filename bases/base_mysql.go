package bases

import (
	"database/sql"
	"errors"
	"fmt"
	"fuckdb/bases/model"
	"sort"
	"strings"
)

// newDB connect mysql
func newDB(c *model.MysqlConfig) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", model.GenDataSource(c, "charset=utf8mb4&parseTime=true&loc=Local"))
	// Check for error in db, note this does not check connectivity but does check uri
	if err != nil {
		fmt.Println("Error opening mysql db: " + err.Error())
		return
	}
	return
}

// GetColumnsFromMysqlTable Select column details from information schema and return map of map
func GetColumnsFromMysqlTable(c *model.MysqlConfig) (*map[string]map[string]string, error) {
	db, err := newDB(c)
	if err != nil {
		fmt.Println("Error opening mysql db: " + err.Error())
		return nil, err
	}

	defer db.Close()

	// Store colum as map of maps
	columnDataTypes := make(map[string]map[string]string)
	// Select columnd data from INFORMATION_SCHEMA
	columnDataTypeQuery := "SELECT COLUMN_NAME, COLUMN_KEY, DATA_TYPE, IS_NULLABLE, COLUMN_DEFAULT, COLUMN_COMMENT, CHARACTER_SET_NAME, COLLATION_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?"

	if Debug {
		fmt.Println("running: " + columnDataTypeQuery)
	}

	rows, err := db.Query(columnDataTypeQuery, c.DB, c.Table)

	if err != nil {
		fmt.Println("Error selecting from db: " + err.Error())
		return nil, err
	}
	if rows != nil {
		defer rows.Close()
	} else {
		return nil, errors.New("No results returned for table ")
	}

	for rows.Next() {
		var column string
		var columnKey string
		var dataType string
		var nullable string
		var defaultVal string
		var comment string
		var charSet string
		var collection string

		rows.Scan(&column, &columnKey, &dataType, &nullable, &defaultVal, &comment, &charSet, &collection)

		columnDataTypes[column] = map[string]string{
			"name":       column,
			"value":      dataType,
			"nullable":   nullable,
			"primary":    columnKey,
			"default":    defaultVal,
			"charset":    charSet,
			"collection": collection,
			"comment":    comment,
		}
	}

	return &columnDataTypes, err
}

// Generate go struct entries for a map[string]interface{} structure
func generateMysqlTypes(obj map[string]map[string]string, depth int, jsonAnnotation bool, gormAnnotation bool, xmlAnnotation bool, xormAnnotation bool, fakerAnnotation bool, gureguTypes bool) string {
	structure := "struct {"

	keys := make([]string, 0, len(obj))
	for key := range obj {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		mysqlType := obj[key]
		nullAble := false
		if mysqlType["nullable"] == "YES" {
			nullAble = true
		}

		primary := ""
		if mysqlType["primary"] == "PRI" {
			primary = ";primary_key"
		}

		// Get the corresponding go value type for this mysql type
		valueType := mysqlTypeToGoType(mysqlType["value"], nullAble, gureguTypes)

		fieldName := fmtFieldName(stringifyFirstChar(key))
		var annotations []string
		if gormAnnotation {
			annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s%s\"", key, primary))
		}
		if jsonAnnotation {
			annotations = append(annotations, fmt.Sprintf("json:\"%s\"", key))
		}

		if xmlAnnotation {
			annotations = append(annotations, fmt.Sprintf("xml:\"%s\"", key))
		}
		if xormAnnotation {
			annotations = append(annotations, fmt.Sprintf("xorm:\"%s%s\"", key, primary))
		}
		if fakerAnnotation {
			annotations = append(annotations, fmt.Sprintf("faker:\"%s\"", key))
		}
		if len(annotations) > 0 {
			structure += fmt.Sprintf("\n%s %s `%s`",
				fieldName,
				valueType,
				strings.Join(annotations, " "))

		} else {
			structure += fmt.Sprintf("\n%s %s",
				fieldName,
				valueType)
		}
	}
	return structure
}

// mysqlTypeToGoType converts the mysql types to go compatible sql.NullAble (https://golang.org/pkg/database/sql/) types
func mysqlTypeToGoType(mysqlType string, nullable bool, gureguTypes bool) string {
	switch mysqlType {
	case "tinyint", "int", "smallint", "mediumint":
		if nullable {
			if gureguTypes {
				return gureguNullInt
			}
			return sqlNullInt
		}
		return golangInt
	case "bigint":
		if nullable {
			if gureguTypes {
				return gureguNullInt
			}
			return sqlNullInt
		}
		return golangInt64
	case "char", "enum", "varchar", "longtext", "mediumtext", "text", "tinytext", "json":
		if nullable {
			if gureguTypes {
				return gureguNullString
			}
			return sqlNullString
		}
		return "string"
	case "date", "datetime", "time", "timestamp":
		if nullable && gureguTypes {
			return gureguNullTime
		}
		return golangTime
	case "decimal", "double":
		if nullable {
			if gureguTypes {
				return gureguNullFloat
			}
			return sqlNullFloat
		}
		return golangFloat64
	case "float":
		if nullable {
			if gureguTypes {
				return gureguNullFloat
			}
			return sqlNullFloat
		}
		return golangFloat32
	case "binary", "blob", "longblob", "mediumblob", "varbinary":
		return golangByteArray
	}
	return ""
}
