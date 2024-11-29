package model

import "gorm.io/gorm"

type Tasks struct {
	gorm.Model
	PlannedDateTime string `gorm:"type:varchar(30)"`
	ActualDateTime  string `gorm:"type:varchar(30)"`
	IsDone          bool   `gorm:"type:boolean;default:false"`
	TimeDiff        string `gorm:"type:varchar(50)"`
	Title           string `gorm:"type:varchar(100)"`
	Description     string `gorm:"type:text"`
	// Tags            []Tags `gorm:"many2many:task_tags;foreignKey:ID;joinForeignKey:TaskID;References:ID;joinReferences:TagID;onDelete:CASCADE"`
	Tags []Tags `gorm:"many2many:task_tags"`
}
