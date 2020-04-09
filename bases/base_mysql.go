package bases

import (
	"fmt"
	"fuckdb/bases/model"
	"strings"
)

func generateColumnTypes(columns []model.Column, depth int, jsonAnnotation bool, gormAnnotation bool, xmlAnnotation bool, xormAnnotation bool, fakerAnnotation bool, gureguTypes bool) string {
	structure := "struct {"
	for _, column := range columns {
		// mysqlType := column.Name
		nullAble := false
		if column.Nullable == "YES" {
			nullAble = true
		}

		primary := ""
		if column.Key == "PRI" {
			primary = ";primary_key"
		}

		// Get the corresponding go value type for this mysql type
		valueType := mysqlTypeToGoType(column.DataType, nullAble, gureguTypes)

		fieldName := fmtFieldName(stringifyFirstChar(column.Name))
		var annotations []string
		if gormAnnotation {
			annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s%s\"", column.Name, primary))
		}
		if jsonAnnotation {
			annotations = append(annotations, fmt.Sprintf("json:\"%s\"", column.Name))
		}

		if xmlAnnotation {
			annotations = append(annotations, fmt.Sprintf("xml:\"%s\"", column.Name))
		}
		if xormAnnotation {
			annotations = append(annotations, fmt.Sprintf("xorm:\"%s%s\"", column.Name, primary))
		}
		if fakerAnnotation {
			annotations = append(annotations, fmt.Sprintf("faker:\"%s\"", column.Name))
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
