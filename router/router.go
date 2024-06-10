package router

import (
	"shops/handler"
	"shops/repository"
	"shops/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Setup() *gin.Engine {
	router := gin.Default()

	var r = *repository.NewItemRepository(sqlite.Open("items.db"), &gorm.Config{})
	var s = services.NewItemService(r)
	var h = handler.NewItemHandler(*s)

	router.POST("/items", h.CreateItem)

	return router
}
