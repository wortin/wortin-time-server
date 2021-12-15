package po

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name            string
	TargetID        uint
	PlanedStartDate string
	PlanedEndDate   string
	Desc            string
	State           int // 1 删除在回收站; 2 已完成;
}
