package models

import "gorm.io/gorm"

// Definición del tipo de usuario (enum)

type UserType string

const (
	CustomerUser UserType = "customer"
	MerchantUser UserType = "merchant"
)

type User struct {
	gorm.Model
	Name     string   `json:"name"`
	Email    string   `gorm:"unique" json:"email"`
	Password string   `json:"password"`
	Location string   `json:"location"`
	Phone    string   `json:"phone"`
	Address  string   `json:"address"`
	UserType UserType `json:"user_type"`
	Merchant Merchant `json:"merchant"`
}
