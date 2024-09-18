package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Price       float64           `json:"price"`
	CategoryID  uint              `json:"category_id"` // FK a Categoria
	Category    Category          `json:"category"`
	Merchants   []MerchantProduct `json:"merchants"` // Relación 1:N con ComercioProducto
}
