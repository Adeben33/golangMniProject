package initializers

import "github.com/adeben33/golangMiniProject/Todo/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Todo{})
}
