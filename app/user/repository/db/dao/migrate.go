package dao

import "wait-to-do/app/user/repository/db/model"

func migration() error {
	return _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{})
}
