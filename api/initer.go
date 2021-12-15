package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Server *gin.Engine

func init() {
	Server = gin.Default()
	// 跨域
	Server.Use(Cors)
	// 路由
	Server.GET("tags", QueryTags)
	Server.PUT("actions/:actionID/tags/:tagID", ActionAddTag)
	Server.POST("project", CreateProject)
	Server.GET("projects", QueryProjects)
}

func Cors(c *gin.Context) {
	method := c.Request.Method
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,Access-Token")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Expose-Headers", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	c.Next()
}
