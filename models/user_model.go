package models

type User struct {
	CommonModelFields
	FullName    string        `json:"full_name" gorm:"type:varchar(100)"`
	Email       string        `json:"email,omitempty" gorm:"type:varchar(100);unique_index"`
	PhoneNumber string        `json:"phone_number" gorm:"type:varchar(100);unique_index;unique"`
	Addresses   []UserAddress `json:"addresses" gorm:"foreignKey:UserID"`
}
