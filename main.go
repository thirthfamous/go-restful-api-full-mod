package main

import (
	"net/http"
	"thirthfamous/golang-restful-api-clean-architecture/app"
	"thirthfamous/golang-restful-api-clean-architecture/controller"
	"thirthfamous/golang-restful-api-clean-architecture/helper"
	"thirthfamous/golang-restful-api-clean-architecture/middleware"
	"thirthfamous/golang-restful-api-clean-architecture/repository"
	"thirthfamous/golang-restful-api-clean-architecture/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
