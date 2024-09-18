package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	MerchantProductID uint      `json:"merchant_product_id"` // FK a ComercioProducto
	Message           string    `json:"message"`
	SentAt            time.Time `json:"sent_at"` // Fecha y hora de envío
}
