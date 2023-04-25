package models

type Address struct {
	AddressID uint `gorm:"primary_key;column:addressID" json:"addressID"`
	ProfileID int64 `gorm:"column:profileID" json:"profileID"`
    CreatedAt string `json:"createdAt,omitempty"`
    UpdatedAt string `json:"updatedAt,omitempty"`
}


