package main

import (
	"gin-gorm-go/common"
	"gin-gorm-go/migration"
	"gin-gorm-go/restaurant"
	"gin-gorm-go/user"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {

	common.InitDB()
	migration.RunMigrations()

}

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello World"})
	})

	r.GET("/template1", func(ctx *gin.Context) {
		ctx.HTML(200, "template1.html", gin.H{})
	})

	r.GET("/template2", func(ctx *gin.Context) {
		ctx.HTML(200, "template2.tmpl", gin.H{"name": "Kevin"})
	})

	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/", user.GetUserAll)
		userRoutes.GET("/:id", user.GetUser)
		userRoutes.POST("/create", user.CreateUser)
		userRoutes.PATCH("/update/:id", user.UpdateUser)
		userRoutes.DELETE("/delete/:id", user.DeleteUser)
	}

	restaurantRoutes := r.Group("/restaurant")
	{
		restaurantRoutes.GET("/", restaurant.GetRestaurantAll)
		restaurantRoutes.GET("/:id", restaurant.GetRestaurant)
		restaurantRoutes.POST("/create", restaurant.CreateRestaurant)
		restaurantRoutes.PATCH("/update/:id", restaurant.UpdateRestaurant)
		restaurantRoutes.DELETE("/delete/:id", restaurant.DeleteRestaurant)

		restaurantRoutes.GET("/addLike/:id", restaurant.AddLike)
		restaurantRoutes.GET("/addDislike/:id", restaurant.AddDislike)
	}

	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal("Failed to start server")
	}

}
