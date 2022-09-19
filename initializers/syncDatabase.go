package initializers

import "Adam-backend-go/model"

func SyncDatabase() {
	DB.AutoMigrate(model.Accounts{})
}
