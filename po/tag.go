package po

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name      string `gorm:"unique"`
	Color     string
	IsDefault bool
}