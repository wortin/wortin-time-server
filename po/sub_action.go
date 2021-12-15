package po

import "gorm.io/gorm"

type SubAction struct {
	gorm.Model
	Name        string
	ActionID    uint
	IsCompleted bool
	State       int // 1 删除在回收站; 2 已完成;
}
