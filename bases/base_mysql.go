package bases

import (
	"bytes"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func ConnectToDB(user, password, host string, port int, dbname string) (*sql.DB, error) {
	var err error
	var db *sql.DB
	if password != "" {
		db, err = sql.Open("mysql", user+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/"+dbname+"?&parseTime=True")
	} else {
		db, err = sql.Open("mysql", user+"@tcp("+host+":"+strconv.Itoa(port)+")/"+dbname+"?&parseTime=True")
	}

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// GetTableInfo query table details from information schema
func GetTableInfo(user string, password string, host string, port int, dbname string, tableName string) (*Table, error) {

	db, err := ConnectToDB(user, password, host, port, dbname)

	// Check for error in db, note this does not check connectivity but does check uri
	if err != nil {
		logrus.Errorln("Error opening mysql db: " + err.Error())
		return nil, err
	}

	defer db.Close()

	// Select column data from INFORMATION_SCHEMA
	columnDataTypeQuery := "SELECT `COLUMN_NAME`, COLUMN_KEY, DATA_TYPE,COLUMN_DEFAULT, IS_NULLABLE,CHARACTER_MAXIMUM_LENGTH,EXTRA,COLUMN_COMMENT FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = ? AND table_name = ?"

	rows, err := db.Query(columnDataTypeQuery, dbname, tableName)
	if err != nil {
		logrus.Errorln("Error selecting from db: " + err.Error())
		return nil, err
	}

	table := &Table{
		Name:    tableName,
		Columns: make([]Column, 0),
	}
	for rows.Next() {
		var columnName string
		var columnKey string
		var dataType string
		var defaultValue sql.NullString
		var nullable string
		var maxLength sql.NullInt64
		var extra string
		var comment string

		if err := rows.Scan(&columnName, &columnKey, &dataType, &defaultValue, &nullable, &maxLength, &extra, &comment); err != nil {
			logrus.Errorln("scan rows error:", err)
			return nil, err
		}

		table.Columns = append(table.Columns, Column{
			Name:       columnName,
			Comment:    comment,
			Type:       dataType,
			Size:       int(maxLength.Int64),
			NotNull:    !(nullable == "YES"),
			Default:    defaultValue.String,
			AutoInc:    isAutoInc(extra),
			Unique:     isUnique(columnKey),
			PrimaryKey: isPrimaryKey(columnKey),
		})
	}

	return table, nil
}

func isAutoInc(extra string) bool {
	return strings.Contains(extra, "auto_increment")
}

func isUnique(columnKey string) bool {
	return isPrimaryKey(columnKey) || strings.Contains(columnKey, "UNI")
}

func isPrimaryKey(columnKey string) bool {
	return strings.Contains(columnKey, "PRI")
}

// generateMysqlTypes go struct entries for a map[string]interface{} structure
func generateModelFields(table *Table, depth int, option *GenerateOption) string {

	structure := ""

	for _, column := range table.Columns {

		// Get the corresponding go value type for this mysql type
		valueType := mysqlTypeToGoType(column.Type, !column.NotNull, option.WithGureguTypes)

		fieldName := fmtFieldName(stringifyFirstChar(column.Name))
		var annotations []string
		if option.WithGormAnnotation {
			annotations = append(annotations, generateGormAnnotation(column))
		}
		if option.WithJsonAnnotation {
			annotations = append(annotations, generateJSONAnnotation(column))
		}
		if option.WithDBAnnotation {
			annotations = append(annotations, generateDBAnnotation(column))
		}
		if option.WithXmlAnnotation {
			annotations = append(annotations, generateXMLAnnotation(column))
		}
		if option.WithXormAnnotation {
			annotations = append(annotations, generateXormAnnotation(column))
		}
		if option.WithFakerAnnotation {
			annotations = append(annotations, generateFakerAnnotation(column))
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

func generateGormAnnotation(col Column) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf("column:%s", col.Name))
	if col.PrimaryKey {
		buf.WriteString(";primaryKey")
	}
	if col.Unique {
		buf.WriteString(";unique")
	}
	if col.NotNull {
		buf.WriteString(";not null")
	}
	if col.AutoInc {
		buf.WriteString(";autoIncrement")
	}
	return fmt.Sprintf(`gorm:"%s"`, buf.String())
}
func generateJSONAnnotation(col Column) string {
	return fmt.Sprintf(`json:"%s"`, col.Name)
}
func generateDBAnnotation(col Column) string {
	return fmt.Sprintf(`db:"%s"`, col.Name)
}
func generateXMLAnnotation(col Column) string {
	return fmt.Sprintf(`xml:"%s"`, col.Name)
}
func generateXormAnnotation(col Column) string {
	return fmt.Sprintf(`xorm:"%s"`, col.Name)
}
func generateFakerAnnotation(col Column) string {
	return fmt.Sprintf(`faker:"%s"`, col.Name)
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
