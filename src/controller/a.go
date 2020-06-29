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
