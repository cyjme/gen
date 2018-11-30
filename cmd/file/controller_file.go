package file

import (
	"gen/cmd/vars"
	"gen/util"
	"os"
	"strings"
)

type ControllerFile struct {
	ModelName   string
	Tpl         string
	Content     string
}

func NewControllerFile(modelName string) *ControllerFile {
	controllerFile := ControllerFile{}

	controllerFile.ModelName = modelName
	controllerFile.Tpl = getDefaultControllerTpl()

	return &controllerFile
}

func (m *ControllerFile) Write() error {
	filePath := "./controller/" + m.ModelName + "_controller.go"
	f, err := os.Create(filePath)
	m.Generate()

	f.Write([]byte(m.Tpl))

	util.FormatSourceCode(filePath)

	return err
}

func (m *ControllerFile) Generate() {
	m.Tpl = strings.Replace(m.Tpl, "{{modelName}}", m.ModelName, -1)
	m.Tpl = strings.Replace(m.Tpl, "{{ModelName}}", strings.Title(m.ModelName), -1)
	m.Tpl = strings.Replace(m.Tpl, "{{projectName}}", vars.ProjectName, -1)
}

func getDefaultControllerTpl() string {
	return `//generate by gen
package controller

import (
	"{{projectName}}/model"
	"{{projectName}}/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type {{ModelName}}Controller struct {
}

// @Summary Create
// @Tags    {{ModelName}}
// @Param body body model.{{ModelName}} true "{{ModelName}}"
// @Success 200 {string} json ""
// @Router /{{modelName}}s [post]
func (ctl *{{ModelName}}Controller) Create(c *gin.Context) {
	{{modelName}} := model.{{ModelName}}{
	}

	if err := pkg.ParseRequest(c, &{{modelName}}); err != nil {
		return
	}

	if err := {{modelName}}.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}

	{{modelName}}.Get()

	c.JSON(http.StatusOK, {{modelName}})
}

// @Summary  Delete
// @Tags     {{ModelName}}
// @Param  {{modelName}}Id  path string true "{{modelName}}Id"
// @Success 200 {string} json ""
// @Router /{{modelName}}s/{{{modelName}}Id} [delete]
func (ctl *{{ModelName}}Controller) Delete(c *gin.Context) {
	{{modelName}} := model.{{ModelName}}{}
	{{modelName}}.Id = c.Param("{{modelName}}Id")
	err := {{modelName}}.Delete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

// @Summary Put
// @Tags    {{ModelName}}
// @Param body body model.{{ModelName}} true "{{modelName}}"
// @Param  {{modelName}}Id path string true "{{modelName}}Id"
// @Success 200 {string} json ""
// @Router /{{modelName}}s/{{{modelName}}Id} [put]
func (ctl *{{ModelName}}Controller) Put(c *gin.Context) {
	{{modelName}} := model.{{ModelName}}{}
	{{modelName}}.Id = c.Param("{{modelName}}Id")

	if err := pkg.ParseRequest(c, &{{modelName}}); err != nil {
		return
	}

	err := {{modelName}}.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, {{modelName}})
}

// @Summary Patch
// @Tags    {{ModelName}}
// @Param body body model.{{ModelName}} true "{{modelName}}"
// @Param  {{modelName}}Id path string true "{{modelName}}Id"
// @Success 200 {string} json ""
// @Router /{{modelName}}s/{{{modelName}}Id} [patch]
func (ctl *{{ModelName}}Controller) Patch(c *gin.Context) {
	{{modelName}} := model.{{ModelName}}{}
	{{modelName}}.Id = c.Param("{{modelName}}Id")

	if err := pkg.ParseRequest(c, &{{modelName}}); err != nil {
		return
	}

	err := {{modelName}}.Patch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, {{modelName}})
}

// @Summary List
// @Tags    {{ModelName}}
// @Param query query string false "query, ?query=age:>:21,name:like:%jason%"
// @Param order query string false "order, ?order=age:desc,created_at:asc"
// @Param page query int false "page"
// @Param pageSize query int false "pageSize"
// @Success 200 {array} model.{{ModelName}} "{{modelName}} array"
// @Router /{{modelName}}s [get]
func (ctl *{{ModelName}}Controller) List(c *gin.Context) {
	{{modelName}} := &model.{{ModelName}}{}
	{{modelName}}.Id = c.Param("{{modelName}}Id")
	var err error

	pageParam := c.DefaultQuery("page", "-1")
	pageSizeParam := c.DefaultQuery("pageSize", "-1")
	rawQuery := c.DefaultQuery("query", "")
	rawOrder := c.DefaultQuery("order", "")

	pageInt, err := strconv.Atoi(pageParam)
	pageSizeInt, err := strconv.Atoi(pageSizeParam)

	offset := pageInt*pageSizeInt - pageSizeInt
	limit := pageSizeInt

	if pageInt < 0 || pageSizeInt < 0 {
		limit = -1
	}

	{{modelName}}s, err := {{modelName}}.List(rawQuery, rawOrder, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, {{modelName}}s)
}


// @Summary Get
// @Tags    {{ModelName}}
// @Param  {{modelName}}Id path string true "{{modelName}}Id"
// @Success 200 {object} model.{{ModelName}} "{{modelName}} object"
// @Router /{{modelName}}s/{{modelName}}Id [get]
func (ctl *{{ModelName}}Controller) Get(c *gin.Context) {
	{{modelName}} := &model.{{ModelName}}{}
	{{modelName}}.Id = c.Param("{{modelName}}Id")

	{{modelName}}, err := {{modelName}}.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, {{modelName}})
}
`
}
