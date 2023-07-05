package models

import (
	"time"
)

type User struct {
	ID          uint64      `gorm:"primary_key" json:"id"`
	Phone       string      `gorm:"size:256;default:null;uniqueIndex:idx_phone,unique" json:"phone" validate:"min=3,max=32"`
	Email       string      `gorm:"size:256;uniqueIndex:idx_email,unique;not null;" json:"email" validate:"required,email,min=6,max=32"`
	EmailVerify time.Time   `gorm:"size:256;default:null;" json:"deleted_at,omitempty"`
	PhoneVerify time.Time   `gorm:"size:256;default:null;" json:"deleted_at,omitempty"`
	Name        string      `gorm:"size:256;default:null;" json:"name" validate:"min=1,max=32"`
	LastName    string      `gorm:"size:256;default:null;" json:"last_name" validate:"min=1,max=32"`
	Surname     string      `gorm:"size:256;default:null;" json:"surname" validate:"min=1,max=32"`
	Password    string      `gorm:"size:256;default:null;" json:"-" validate:"min=1,max=32"`
	Secret      string      `gorm:"size:256;default:null;" json:"-" validate:"max=32"`
	Images      []Image     `gorm:"polymorphic:Owner" json:"images"`
	UserSocials UserSocials `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	UpdatedAt   time.Time   `gorm:"autoCreateTime:nano;autoUpdateTime:nano" json:"updated_at,omitempty"`
	CreatedAt   time.Time   `gorm:"autoCreateTime:nano;" json:"created_at,omitempty"`
	DeletedAt   *time.Time  `json:"deleted_at,omitempty"`
}
