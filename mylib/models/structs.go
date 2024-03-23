package models

import (
	"time"
)

// Item represents an item in the items table
type Item struct {
	// gorm.Model
	ItemID      uint   `gorm:"primaryKey" json:"itemID"`
	ItemCode    string `gorm:"type:int" json:"itemCode"`
	Description string `gorm:"type:varchar(300)" json:"description"`
	Quantity    int    `gorm:"type:int" json:"quantity"`
	OrderID     uint
}

// Order represents an order in the orders table
type Order struct {
	// gorm.Model
	OrderID      uint      `gorm:"primaryKey" json:"orderID"`
	CustomerName string    `gorm:"type:varchar(300)" json:"customerName"`
	OrderedAt    time.Time `gorm:"type:timestamp with time zone" json:"ordered_at"`
	Items        []Item    `json:"items"`
}
