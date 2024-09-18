package models

import "gorm.io/gorm"

type Merchant struct {
	gorm.Model
	Name          string            `json:"name"`
	Email         string            `gorm:"unique" json:"email"`
	Location      string            `json:"location"` // Coordenadas
	Phone         string            `json:"phone"`
	Address       string            `json:"address"`
	UserID        uint              `json:"user_id"`       // FK a Usuario
	Products      []MerchantProduct `json:"products"`      // Relación 1:N con ComercioProducto
	Subscriptions []Subscription    `json:"subscriptions"` // Relación 1:N con Subscripciones
}
