package file

import (
	"os"
	"strings"

	"github.com/cyjme/gen/cmd/vars"
	"github.com/cyjme/gen/util"
)

type ModelFile struct {
	ProjectName string
	Name        string
	Fields      []string
	Tpl         string
	Content     string
	NeedCommon  bool
}

func NewModelFile(name string, fields []string, needCommon bool) *ModelFile {
	modelFile := ModelFile{}

	modelFile.Name = name
	modelFile.Fields = fields
	modelFile.NeedCommon = needCommon
	modelFile.Tpl = getDefaultModelTpl()

	return &modelFile
}

func (m *ModelFile) Write() error {
	filePath := "./model/" + m.Name + ".go"
	f, err := os.Create(filePath)
	m.Generate()

	f.Write([]byte(m.Tpl))

	util.FormatSourceCode(filePath)

	return err
}

func (m *ModelFile) Generate() {
	structText := ""
	if m.NeedCommon {
		structText += `	Common
`
	}

	for _, field := range m.Fields {
		fieldSplit := strings.Split(field, ":")
		fieldName := fieldSplit[0]
		fieldType := fieldSplit[1]
		oneLineText := "	" + strings.Title(fieldName) + "    " + fieldType + "    " + "`json:" + "\"" + fieldName + "\"" + "`"
		structText += oneLineText
		structText += `
`
	}

	m.Tpl = strings.Replace(m.Tpl, "{{modelStruct}}", structText, -1)
	m.Tpl = strings.Replace(m.Tpl, "{{modelName}}", m.Name, -1)
	m.Tpl = strings.Replace(m.Tpl, "{{ModelName}}", strings.Title(m.Name), -1)
	m.Tpl = strings.Replace(m.Tpl, "{{projectName}}", vars.ProjectName, -1)
}

func getDefaultModelTpl() string {
	return `//generate by gen
package model

import (
	"{{projectName}}/app"
)

type {{ModelName}} struct {
{{modelStruct}}
}

func ({{modelName}} *{{ModelName}}) Insert() error {
	err := app.DB.Create({{modelName}}).Error

	return err
}

func ({{modelName}} *{{ModelName}}) Patch() error {
	err := app.DB.Model({{modelName}}).Updates({{modelName}}).Error

	return err
}

func ({{modelName}} *{{ModelName}}) Update() error {
	err := app.DB.Save({{modelName}}).Error

	return err
}

func ({{modelName}} *{{ModelName}}) Delete() error {
	return app.DB.Delete({{modelName}}).Error
}

func ({{modelName}} *{{ModelName}}) List(rawQuery string, rawOrder string, offset int, limit int) (*[]{{ModelName}}, int, error) {
	{{modelName}}s := []{{ModelName}}{}
	total := 0

	db := app.DB.Model({{modelName}})

	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return &{{modelName}}s, total, err
	}

	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return &{{modelName}}s, total, err
	}

	db.Offset(offset).
		Limit(limit).
		Find(&{{modelName}}s).
		Count(&total)

	err = db.Error

	return &{{modelName}}s, total, err
}

func ({{modelName}} *{{ModelName}}) Get() (*{{ModelName}}, error) {
	err := app.DB.Find(&{{modelName}}).Error

	return {{modelName}}, err
}

`
}
