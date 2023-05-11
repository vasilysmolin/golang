package models


type User struct {
	UserID  int64  `gorm:"primary_key;column:userID" json:"userID"`
	Phone string `gorm:"column:phone" json:"phone"`
	Avatar string `gorm:"column:avatar" json:"avatar"`
    CreatedAt string `json:"createdAt,omitempty"`
    UpdatedAt string `json:"updatedAt,omitempty"`
}

func (User) TableName() string {
    return "users"
}
