package models

type User struct {
	CommonModelFields
	UUID        string `gorm:"type:varchar(100);unique_index;not_null;unique" json:"uuid"`
	FullName    string `json:"full_name" gorm:"type:varchar(100)"`
	Email       string `json:"email,omitempty" gorm:"type:varchar(100);unique_index"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(100);unique_index;unique"`
}
