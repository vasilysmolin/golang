package models

import (
	"time"
)

type UserSocials struct {
	ID          uint64     `gorm:"primary_key;" json:"id"`
	SocialID    uint64     `gorm:"size:256;uniqueIndex:idx_social_id,unique;not null;" json:"social_id" validate:"required,min=3,max=256"`
	UserID      uint64     `gorm:"index:idx_user_id;not null;" json:"user_id"`
	Type        string     `gorm:"type:enum('vk', 'ok');check:provider IN ('vk', 'ok');not null" json:"type" validate:"oneof=vk ok google apple"`
	AccessToken string     `gorm:"size:3000;not null;" json:"access_token"`
	UpdatedAt   time.Time  `gorm:"autoCreateTime:nano;autoUpdateTime:nano" json:"updated_at,omitempty"`
	CreatedAt   time.Time  `gorm:"autoCreateTime:nano;" json:"created_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
