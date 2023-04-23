package restaurant

import (
	"gin-gorm-go/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateRestaurant(ctx *gin.Context) {

	var restaurant Restaurant
	err := ctx.ShouldBind(&restaurant)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid Restaurant"})
		return
	}

	err = common.DB.Create(&restaurant).Error
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Success"})

}
func GetRestaurant(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 0 {
		ctx.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	restaurant := Restaurant{ID: uint(idInt)}
	common.DB.First(&restaurant)

	ctx.JSON(200, gin.H{"restaurant": restaurant})

}
func GetRestaurantAll(ctx *gin.Context) {

	var restaurants []Restaurant

	err := common.DB.Find(&restaurants).Error
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	ctx.JSON(200, gin.H{"restaurants": restaurants})

}
func UpdateRestaurant(ctx *gin.Context) {
	var restaurant Restaurant
	err := ctx.ShouldBind(&restaurant)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Bad request"})
		return
	}
	err = common.DB.Save(&restaurant).Error
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Success"})
}

func DeleteRestaurant(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 0 {
		ctx.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	restaurant := Restaurant{ID: uint(idInt)}

	err = common.DB.Delete(&restaurant).Error
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Success"})

}

func AddLike(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 0 {
		ctx.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	restaurant := Restaurant{ID: uint(idInt)}
	err = common.DB.First(&restaurant).Error
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal server error"})
	}

	restaurant.Likes += 1

	err = common.DB.Save(&restaurant).Error
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal server error"})
	}

	ctx.JSON(200, gin.H{"message": "Success"})

}
func AddDislike(ctx *gin.Context) {

	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 0 {
		ctx.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	restaurant := Restaurant{ID: uint(idInt)}
	err = common.DB.First(&restaurant).Error
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal server error"})
	}

	restaurant.Dislikes += 1

	err = common.DB.Save(&restaurant).Error
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal server error"})
	}

	ctx.JSON(200, gin.H{"message": "Success"})

}
