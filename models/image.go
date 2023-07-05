package models

import (
	"time"
)

type Image struct {
	ID             uint64 `gorm:"primary_key" json:"id"`
	Name           string `gorm:"size:256;index:idx_image_name" json:"name" validate:"max=256"`
	MimeType       string `gorm:"size:100;" json:"mimeType" validate:"max=100"`
	CollectionName string `gorm:"size:100;" json:"collection_name" validate:"max=100"`
	Extension      string `gorm:"size:100;default:null;" json:"extension" validate:"max=100"`
	Disk           string `gorm:"size:100;default:null;" json:"disk" validate:"max=100"`
	Size           uint64 `gorm:"not null;" json:"size" validate:"max=999999"`
	OwnerID        uint
	OwnerType      string
	UpdatedAt      time.Time  `gorm:"autoCreateTime:nano;autoUpdateTime:nano" json:"updated_at,omitempty"`
	CreatedAt      time.Time  `gorm:"autoCreateTime:nano;" json:"created_at,omitempty"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
}
