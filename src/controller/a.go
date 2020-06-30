package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func A1(c *gin.Context) {
	c.String(http.StatusOK, "Welcome Gin Server")
}

func GetParam(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Welcome %s", name)

}

func GetQuery(c *gin.Context) {
	name := c.Query("name")
	c.String(http.StatusOK, "Welcome %s", name)
}

// post 提交
func GetPostValue(c *gin.Context) {
	name := c.PostForm("name")
	c.JSON(200, gin.H{
		"name": name,
	})
}

type Person struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}

// 绑定查询字符串或表单数据
func BuildData(c *gin.Context) {
	var person Person
	if c.ShouldBind(&person) == nil {
		c.JSON(200, gin.H{
			"name": person.Name,
			"age":  person.Age,
		})
	} else {
		c.JSON(400, gin.H{
			"name": "",
			"age":  "",
		})
	}
}
