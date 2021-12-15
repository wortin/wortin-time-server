package po

import "gorm.io/gorm"

type Target struct {
	gorm.Model
	Name string
}
