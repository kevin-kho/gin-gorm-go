package user

import (
	"gin-gorm-go/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid User"})
		return
	}

	err = common.DB.Create(&user).Error
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Success"})

}

func GetUser(ctx *gin.Context) {

	id := ctx.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 0 {
		ctx.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	var user User

	err = common.DB.First(&user, idInt).Error
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	ctx.JSON(200, gin.H{"user": user})

}

func GetUserAll(ctx *gin.Context) {

	var users []User

	err := common.DB.Find(&users).Error
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	ctx.JSON(200, gin.H{"users": users})

}

func UpdateUser(ctx *gin.Context) {
	var user User

	id := ctx.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 0 {
		ctx.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	err = ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	if user.ID != uint(idInt) {
		ctx.JSON(400, gin.H{"message": "path id mismatch"})
		return
	}

	err = common.DB.Save(&user).Error
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Success"})
}

func DeleteUser(ctx *gin.Context) {

	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 0 {
		ctx.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	user := User{ID: uint(idInt)}

	err = common.DB.Delete(&user).Error
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
	})
}
