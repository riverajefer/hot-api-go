package initializers

import "github.com/riverajer/hot-bread-api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Category{})
}
