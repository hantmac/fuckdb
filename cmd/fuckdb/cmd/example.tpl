package {{ .PackageName }}

type {{ .ModelName }} struct {
	{{ .Fields }}
}

{{ if .Option.WithGormAnnotation }}
	// TableName sets the insert table name for this struct type
    func (model *{{ .ModelName }})TableName() string {
    	return "{{ .Table.Name }}"
    }
{{ end }}

{{ if .Option.WithGormAnnotation }}
 	func Add{{.ModelName}}(m *{{.ModelName}}) error {
    	return db.Save(m).Error
    }

    func Delete{{.ModelName}}ByID(id int) (bool, error) {
    	if err := db.Delete(&{{.ModelName}}{}, id).Error; err != nil {
    		return false, err
    	}
    	return db.RowsAffected > 0, nil
    }

    func Delete{{.ModelName}}(condition string, args ...interface{}) (int64, error) {
    	if err := db.Where(condition, args...).Delete(&{{.ModelName}}{}).Error; err != nil {
    		return 0, err
    	}
    	return db.RowsAffected, nil
    }

    func Update{{.ModelName}}(m *{{.ModelName}}) error {
    	return db.Save(m).Error
    }

    func Get{{.ModelName}}ByID(id int) (*{{.ModelName}}, error) {
    	var m {{.ModelName}}
    	if err := db.First(&m, id).Error; err != nil {
    		return nil, err
    	}
    	return &m, nil
    }

    func Get{{.ModelName}}s(condition string, args ...interface{}) ([]*{{.ModelName}}, error) {
    	res := make([]*{{.ModelName}}, 0)
    	if err := db.Where(condition, args...).Find(&res).Error; err != nil {
    		return nil, err
    	}
    	return res, nil
    }

{{ end }}



