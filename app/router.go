package app

import (
	"thirthfamous/golang-restful-api-clean-architecture/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(categoryController controller.CategoryController) *gin.Engine {

	router := gin.Default()

	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories/", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	return router
}
