package models

type Image struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"type:varchar(300)" json:"title"`
}
