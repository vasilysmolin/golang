package models

type Studio struct {
    ItemID int64 `gorm:"primary_key;column:itemID" json:"itemID"`
	ProfileID int64 `gorm:"column:profileID" json:"profileID"`
	UserID int64 `gorm:"column:userID;foreignKey:UserID;ColumnName:id;RelationshipFKName:id" json:"userID"`
}

func (Studio) TableName() string {
    return "user_studio"
}


