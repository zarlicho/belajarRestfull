package main

import (
	"api/app"
	"api/controller"
	"api/helper"
	"api/repository"
	"api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDb()
	validator := validator.New()
	categoryRepository := repository.NewRepository()
	categoryService := service.NewService(categoryRepository, db, validator)
	categoryController := controller.NewController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories/findall", categoryController.FindAll)
	router.GET("/api/categories/findbyid/:id", categoryController.FindByid)
	router.DELETE("/api/categories/delete/:id", categoryController.Delete)
	router.POST("/api/categories/create", categoryController.Create)
	router.PUT("/api/categories/update/:id", categoryController.Update)
	router.POST("/api/categories/regis", categoryController.Register)
	router.POST("/api/categories/login", categoryController.Login)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicErrorIf(err)
}
