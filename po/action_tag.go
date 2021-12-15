package po

import "gorm.io/gorm"

type ActionTag struct {
	gorm.Model
	ActionID uint `gorm:"unique_index:action_tag_ind"`
	TagID    uint `gorm:"unique_index:action_tag_ind"`
}
