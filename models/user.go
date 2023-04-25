package models


type User struct {
	UserID  int64  `gorm:"primary_key;column:userID" json:"userID"`
	Phone string `json:"phone"`
	Profile  Profile `gorm:"foreignKey:UserID" json:"profile"`
    CreatedAt string `json:"createdAt,omitempty"`
    UpdatedAt string `json:"updatedAt,omitempty"`
}


