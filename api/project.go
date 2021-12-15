package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"wortin-time-server/service"
)

// CreateProject 创建项目 只需要输入项目名，即可创建一个项目
func CreateProject(c *gin.Context) {
	var req CreateProjectReq
	err := c.BindJSON(&req)
	if err != nil {
		paramErr(c, fmt.Sprintf("fail to bind create project request body, please use json body like: {\"projectName\":\"xxx\"}. err [%v]", err))
		return
	}

	err = service.CreateProject(req.ProjectName)
	if err != nil {
		systemErr(c, err)
		return
	}

	successWithoutData(c)
	return
}

type CreateProjectReq struct {
	ProjectName string `json:"projectName"`
}

func QueryProjects(c *gin.Context) {
	pageNo := c.Query("pageNo")
	pn, err := strconv.Atoi(pageNo)
	if err != nil {
		paramErr(c, fmt.Sprintf("fail to parse pageNo [%s], because it is not a number. please input a number. the err is [%v]", pageNo, err))
		return
	}
	if pn <= 0 {
		paramErr(c, fmt.Sprintf("fail to check pageNo [%s], because it is <= 0. please input a positive number", pageNo))
		return
	}
	pageSize := c.Query("pageSize")
	ps, err := strconv.Atoi(pageSize)
	if err != nil {
		paramErr(c, fmt.Sprintf("fail to parse pageSize [%s], because it is not a number. please input a number. the err is [%v]", pageSize, err))
		return
	}
	if ps <= 0 {
		paramErr(c, fmt.Sprintf("fail to check pageSize [%s], because it is <= 0. please input a positive number", pageSize))
		return
	}

	projects, err := service.QueryProjects(uint(pn), uint(ps))
	if err != nil {
		systemErr(c, err)
		return
	}
	success(c, projects)
	return
}
