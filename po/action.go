package po

import "gorm.io/gorm"

type Action struct {
	gorm.Model
	Name        string
	ProjectID   uint
	PlanedDate  string
	Desc        string
	IsCompleted bool
	State       int // 1 删除在回收站; 2 已完成;
}
