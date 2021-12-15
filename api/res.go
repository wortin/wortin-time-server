package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wortin-time-server/bo"
)

const (
	Success   = 50000
	SystemErr = 60001
	ParamErr  = 60002
	HttPErr   = 60003
)

func success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, bo.Response{
		Code:    Success,
		Message: "success",
		Data:    data,
	})
}

func successWithoutData(c *gin.Context) {
	c.JSON(http.StatusOK, bo.Response{
		Code:    Success,
		Message: "success",
	})
}

func systemErr(c *gin.Context, err error) {
	c.JSON(http.StatusOK, bo.Response{
		Code:    SystemErr,
		Message: err.Error(),
	})
}

func httpErr(c *gin.Context, err error) {
	c.JSON(http.StatusOK, bo.Response{
		Code:    HttPErr,
		Message: err.Error(),
	})
}

func paramErr(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, bo.Response{
		Code:    ParamErr,
		Message: msg,
	})
}
