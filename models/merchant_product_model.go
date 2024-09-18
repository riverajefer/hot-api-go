package models

import "gorm.io/gorm"

type MerchantProduct struct {
	gorm.Model
	MerchantID    uint           `json:"merchant_id"`   // FK a Comercio
	ProductID     uint           `json:"product_id"`    // FK a Producto
	Notifications []Notification `json:"notifications"` // Relación 1:N con Notificaciones
}
