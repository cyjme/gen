package file

import (
	"io/ioutil"
	"os"
	"strings"

	"gen/util"
)

type MigrateFile struct {
	ProjectName string
	ModelName   string
	Tpl         string
	Content     string
}

func NewMigrateFile(modelName string) *MigrateFile {
	migrateFile := MigrateFile{}

	migrateFile.ModelName = modelName
	migrateFile.Tpl = getDefaultMigrateTpl()

	return &migrateFile
}

func (r *MigrateFile) Write() error {
	filePath := "./migrate/" + "create_table.go"
	//f, err := os.Create(filePath)
	r.Generate()

	file, err := ioutil.ReadFile(filePath)
	newFile := strings.Replace(string(file), "//!!do not delete the line, gen generate code at here", r.Tpl, -1)

	f, err := os.Create(filePath)
	f.Write([]byte(newFile))

	util.FormatSourceCode(filePath)

	return err
}

func (r *MigrateFile) Generate() {
	r.Tpl = strings.Replace(r.Tpl, "{{modelName}}", r.ModelName, -1)
	r.Tpl = strings.Replace(r.Tpl, "{{ModelName}}", strings.Title(r.ModelName), -1)
}

func getDefaultMigrateTpl() string {
	return `&model.{{ModelName}}{},
	//!!do not delete the line, gen generate code at here
`
}
