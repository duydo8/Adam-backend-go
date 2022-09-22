package initializers

import "Adam-backend-go/model"

func SyncDatabase() {
	DB.AutoMigrate(model.Accounts{}, model.Address{}, model.CartItems{}, model.Category{}, model.Color{}, model.Comments{}, model.DetailOrder{},
		model.DetailProduct{}, model.DiscountOrder{}, model.District{}, model.Event{}, model.Favorite{}, model.HistoryOrder{}, model.Material{}, model.MaterialProduct{},
		model.Order{}, model.Product{}, model.Province{}, model.Size{}, model.TagProduct{}, model.Tag{}, model.Ward{})
}
