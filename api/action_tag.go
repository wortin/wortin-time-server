package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"wortin-time-server/db"
)

func ActionAddTag(c *gin.Context) {
	// 解析参数
	actionID := c.Param("actionID")
	a, err := strconv.Atoi(actionID)
	if err != nil {
		paramErr(c, fmt.Sprintf("fail to parse actionID [%s], because it is not a number. please input a number. the err is [%v]", actionID, err))
		return
	}
	if a <= 0 {
		paramErr(c, fmt.Sprintf("fail to check actionID [%s], because it is <= 0. please input a positive number", actionID))
		return
	}
	tagID := c.Param("tagID")
	t, err := strconv.Atoi(tagID)
	if err != nil {
		paramErr(c, fmt.Sprintf("fail to parse tagID [%s], because it is not a number. please input a number. the err is [%v]", tagID, err))
		return
	}
	if t <= 0 {
		paramErr(c, fmt.Sprintf("fail to check tagID [%s], because it is <= 0. please input a positive number", tagID))
		return
	}

	// 调用
	err = db.ActionAddTag(uint(a), uint(t))
	if err != nil {
		systemErr(c, err)
		return
	}
	successWithoutData(c)
	return
}
