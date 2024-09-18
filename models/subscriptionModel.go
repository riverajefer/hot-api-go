package models

import (
	"gorm.io/gorm"
)

// Subscription struct
type Subscription struct {
	gorm.Model
	UserID     uint `json:"user_id"`
	MerchantID uint `json:"merchant_id"` // FK a Comercio
	// todo revisar este campo
	// 	Notifications datatypes.JSON `json:"notifications"`  // JSON con las preferencias de notificación

	Notifications []Notification `json:"notifications"` // Relación 1:N con Notificaciones
}
