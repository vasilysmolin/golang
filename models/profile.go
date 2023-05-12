package models

import (
    "gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID   int64 `gorm:"column:user_id;uniqueIndex:idx_user_id,unique" json:"user_id"`
}



