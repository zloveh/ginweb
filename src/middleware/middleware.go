package middleware

import (
	"ginweb/src/util"
	"github.com/gin-gonic/gin"
)

func Middleware1(c *gin.Context) {
	//逻辑代码
	util.Infof("method:%s", c.Request.Method)

	// 执行该中间件之后的逻辑
	c.Next()

	// 再回过来执行
	util.Infof("finall exec")
}
