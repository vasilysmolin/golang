package models

type Profile struct {
	ProfileID int64 `gorm:"primary_key;column:profileID" json:"profileID"`
	UserID   int64 `gorm:"column:userID" json:"userID"`
	Address  Address `gorm:"foreignKey:ProfileID" json:"address"`
    CreatedAt string `json:"createdAt,omitempty"`
    UpdatedAt string `json:"updatedAt,omitempty"`
}


