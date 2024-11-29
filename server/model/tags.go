package model

import "gorm.io/gorm"

type Tags struct {
	gorm.Model
	Name            string `gorm:"type:varchar(255)"`
	TextColor       string `gorm:"type:varchar(7)"`
	BackgroundColor string `gorm:"type:varchar(7)"`
	// Tasks           []Tasks `gorm:"many2many:task_tags;foreignKey:ID;joinForeignKey:TagID;References:ID;joinReferences:TaskID;onDelete:CASCADE"`
	Tasks []Tasks `gorm:"many2many:task_tags"`
}
