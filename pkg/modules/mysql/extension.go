package mysql

import "software-theory-final/pkg/modules/model"

func MigrateTables(client *Client) {
	client.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")
	if err := client.DB.Debug().AutoMigrate(&model.Algorithm{}, model.Dataset{}, &model.Prediction{}, &model.User{}).Error; err != nil {
		panic(err)
	}
}
