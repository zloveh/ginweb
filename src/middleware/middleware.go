package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Middleware1(c *gin.Context) {
	log.Println("exec middleware1")

	//你可以写一些逻辑代码

	// 执行该中间件之后的逻辑
	c.Next()
}
