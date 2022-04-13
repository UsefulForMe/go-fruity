package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	CommonModelFields
	TextID string    `json:"text_id" gorm:"type:varchar(255);unique_index"`
	UserID uuid.UUID `json:"user_id"`
	User   User      `json:"user"`

	Status        string      `json:"status" gorm:"type:varchar(255); default:processing"`
	SellerID      uuid.UUID   `json:"seller_id"`
	Seller        Seller      `json:"seller"`
	OrderItems    []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
	PaymentID     uuid.UUID   `json:"payment_id"`
	Payment       Payment     `json:"payment"`
	Tracks        []Track     `json:"tracks" gorm:"foreignKey:OrderID"`
	ReceivedAt    time.Time   `json:"received_at"`
	UserAddressID uuid.UUID   `json:"user_address_id" `
	UserAddress   UserAddress `json:"user_address"`
	Note          string      `json:"note"`
	TotalPrice    float32     `json:"total_price,omitempty"  gorm:"<-:false;->;-:migration" `
}

func (o *Order) BeforeCreate(db *gorm.DB) (err error) {
	o.TextID = "ORD-" + time.Now().Format("20060102150405")
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
	return
}

type OrderItem struct {
	CommonModelFields
	OrderID   uuid.UUID `json:"order_id" gorm:"type:uuid"`
	ProductID uuid.UUID `json:"product_id" gorm:"type:uuid"`
	Product   Product   `json:"product"`
	Quantity  int       `json:"quantity" gorm:"type:int"`
	Price     float32   `json:"price" gorm:"type:numeric"`
	Note      string    `json:"note" gorm:"type:varchar(100)"`
}

type Track struct {
	CommonModelFields
	OrderID uuid.UUID `json:"order_id" gorm:"type:uuid"`
	Status  string    `json:"status" gorm:"type:varchar(100);default:processing"`
	Time    time.Time `json:"time" `
	Note    string    `json:"note" gorm:"type:varchar(100)"`
}
