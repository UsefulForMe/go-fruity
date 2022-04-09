package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Seller struct {
	CommonModelFields
	Name          string        `json:"name" gorm:"type:varchar(255);not null"`
	Logo          string        `json:"logo" gorm:"type:varchar(255);not null"`
	Banner        string        `json:"banner" gorm:"type:varchar(255);not null"`
	Type          string        `json:"type" gorm:"type:varchar(255);not null"`
	PhoneNumber   string        `json:"phone_number" gorm:"type:varchar(255)"`
	Description   string        `json:"description" gorm:"not null"`
	HeadQuarter   string        `json:"head_quarter" gorm:"not null"`
	Rating        float32       `json:"rating" gorm:"type:numeric;not null"`
	AvailableTime AvailableTime `json:"available_time" gorm:"type:jsonb"`
	Note          string        `json:"note" gorm:"type:varchar(255);"`
	Email         string        `json:"email" gorm:"type:varchar(255);"`
	TotalVote     int           `json:"total_vote" gorm:"type:int;not null; default: 0"`
	Products      []Product     `json:"products" gorm:"foreignKey:SellerID"`
}

type AvailableTime struct {
	OpenTime  string    `json:"open_time"`
	CloseTime string    `json:"close_time"`
	Holidays  []holiday `json:"holidays"`
}

func (a AvailableTime) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *AvailableTime) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

type holiday struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (a holiday) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *holiday) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
