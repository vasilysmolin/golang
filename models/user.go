package models


type User struct {
	UserID  int64  `gorm:"primary_key;column:userID" json:"userID"`
	Phone string `gorm:"column:phone" json:"phone"`
	Profile  Profile `gorm:"foreignKey:UserID" json:"profile"`
// 	Profile Profile `gorm:"many2many:user_studio;references:UserID" json:"studioProfile"`
    CreatedAt string `json:"createdAt,omitempty"`
    UpdatedAt string `json:"updatedAt,omitempty"`
}

func (User) TableName() string {
    return "users"
}
