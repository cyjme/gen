package file

import (
	"io/ioutil"
	"os"
	"strings"

	"gen/util"
)

type RouterFile struct {
	ProjectName string
	ModelName   string
	Tpl         string
	Content     string
}

func NewRouterFile(modelName string) *RouterFile {
	routerFile := RouterFile{}

	routerFile.ModelName = modelName
	routerFile.Tpl = getDefaultRouterTpl()

	return &routerFile
}

func (r *RouterFile) Write() error {
	filePath := "./router/" + "router.go"
	//f, err := os.Create(filePath)
	r.Generate()

	file, err := ioutil.ReadFile(filePath)
	newFile := strings.Replace(string(file), "//!!do not delete gen will generate router code at here", r.Tpl, -1)

	f, err := os.Create(filePath)
	f.Write([]byte(newFile))

	util.FormatSourceCode(filePath)

	return err
}

func (r *RouterFile) Generate() {
	r.Tpl = strings.Replace(r.Tpl, "{{modelName}}", r.ModelName, -1)
	r.Tpl = strings.Replace(r.Tpl, "{{ModelName}}", strings.Title(r.ModelName), -1)
}

func getDefaultRouterTpl() string {
	return `
	{{modelName}}Controller := controller.{{ModelName}}Controller{}
	{{modelName}}Group := r.Group("/{{modelName}}s")
	{
		{{modelName}}Group.GET("", {{modelName}}Controller.List)
		{{modelName}}Group.POST("", {{modelName}}Controller.Create)
		{{modelName}}Group.DELETE("/:{{modelName}}Id", {{modelName}}Controller.Delete)
		{{modelName}}Group.PUT("/:{{modelName}}Id", {{modelName}}Controller.Put)
		{{modelName}}Group.GET("/:{{modelName}}Id", {{modelName}}Controller.Get)
		{{modelName}}Group.PATCH("/:{{modelName}}Id", {{modelName}}Controller.Patch)
	}

	//!!do not delete gen will generate router code at here
`
}
