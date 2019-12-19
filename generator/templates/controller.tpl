package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"strconv"
	"weibang/components"
	"weibang/http/form"
	"weibang/models"
	"weibang/services"
)
var (
    {{.Name}} {{.CreatorName}}
)

type {{.CreatorName}} struct {
}
func (i *{{.CreatorName}}) Post(c *gin.Context) {
	var form{{.Name}} form.ApiForm{{.Name}}
	err := c.Bind(&form{{.Name}})
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	model := services.{{.Name}}.Find(form{{.Name}}.ID)
	if model.ID > 0 {
		c.JSON(201, gin.H{
			"message": "",
		})
		return
	} else {
		var model models.{{.Name}}

		copier.Copy(&model, &form{{.Name}})

		model = services.{{.Name}}.Save(model)
		if model.ID > 0 {
			c.JSON(201, gin.H{
				"message": "",
			})
			return

		} else {
			c.JSON(500, gin.H{
				"message": "server error",
			})
			return
		}
	}
}

func (i *{{.CreatorName}}) Delete(c *gin.Context) {

	id,_ := strconv.ParseUint(c.Param("id"),10,32)
	succ := services.{{.Name}}.Delete(uint(id))
	var code int
	if succ {
		code = 204
	} else {
		code = 404
	}
	c.JSON(code, gin.H{
		"message": "",
	})
}

func (i *{{.CreatorName}}) Put(c *gin.Context) {

	var form{{.Name}} form.ApiForm{{.Name}}
	err := c.Bind(&form{{.Name}})
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	model := services.{{.Name}}.Find(form{{.Name}}.ID)
	if model.ID > 0 {
	    copier.Copy(&model, &form{{.Name}})
		model = services.{{.Name}}.Save(model)

		c.JSON(201, gin.H{
			"message": "",
		})
		return

	} else {
		c.JSON(500, gin.H{
			"message": "server error",
		})
		return
	}
}

func (controller *{{.CreatorName}}) GET(c *gin.Context) {
    idstr := c.Param("id")
    id,_:=	strconv.ParseUint(idstr,10,32)
	var out dto.{{.Name}}
	model := services.{{.Name}}.Find(uint(id))

	copier.Copy(&out, &model)

	c.JSON(200, gin.H{
		"message": "",
		"data":    out,
	})
}

func (i *{{.CreatorName}}) Page(c *gin.Context) {
	pageNo := c.DefaultQuery("page", "1")
	var page components.Page
	pageNoInt, _ := strconv.Atoi(pageNo)
	page.CurrentPage = pageNoInt
	lists := services.{{.Name}}.FindAll(page)
	var out []dto.{{.Name}}
	copier.Copy(&out, &lists)
	c.JSON(200, gin.H{
		"message":     "",
		"result":      out,
		"currentPage": pageNoInt,
		"total":       services.{{.Name}}.Count(),
	})
}



