package migration

import (
	"gin-gorm-go/common"
	"gin-gorm-go/restaurant"
	"gin-gorm-go/user"
	"log"
)

func RunMigrations() {
	err := common.DB.AutoMigrate(&user.User{}, &restaurant.Restaurant{})
	if err != nil {
		log.Fatal("Failed to run migrations")
	}

}
