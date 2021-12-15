package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"wortin-time-server/db"
)

func QueryTags(c *gin.Context) {
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

	tags, err := db.QueryTags(uint(pn), uint(ps))
	if err != nil {
		systemErr(c, err)
		return
	}
	success(c, tags)
	return
}
