package models

type AddressesOnlineSettings struct {
	OnlineSettingID  int64  `gorm:"column:onlineSettingID"  json:"onlineSettingID"`
	AddressID int64 `gorm:"type:int64;column:addressID" json:"addressID"`
	RecordNotLater int64 `gorm:"column:recordNotLater" json:"recordNotLater"`
	RecordNotPrev int64 `gorm:"column:recordNotPrev" json:"recordNotPrev"`
	EditableUntil int64 `gorm:"column:editableUntil" json:"editableUntil"`
	ClientRegistrationRequired bool `gorm:"column:clientRegistrationRequired" json:"clientRegistrationRequired"`
	Active int64 `gorm:"column:active" json:"active"`
    CreatedAt string `json:"createdAt,omitempty"`
    UpdatedAt string `json:"updatedAt,omitempty"`
}

func (AddressesOnlineSettings) TableName() string {
    return "addresses_online_settings"
}


