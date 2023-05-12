package models

import (
    "time"
)

type User struct {
    ID  uint64  `gorm:"primary_key;column:id" json:"id"`
	Phone string `gorm:"column:phone;size:256;uniqueIndex:idx_phone,unique" json:"phone" validate:"required,min=3,max=32"`
	Email string `gorm:"column:email;size:256;uniqueIndex:idx_email,unique;" json:"email" validate:"required,email,min=6,max=32"`
	Avatar string `gorm:"column:avatar;size:3000;" json:"avatar"`
    UpdatedAt time.Time `gorm:"autoCreateTime:nano;autoUpdateTime:nano" json:"updated_at,omitempty"`
    CreatedAt time.Time `gorm:"autoCreateTime:nano;" json:"created_at,omitempty"`
    DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

