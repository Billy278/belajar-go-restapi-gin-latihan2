package main

import (
	"belajar-go-restapi-gin-latihan2/app"
	"belajar-go-restapi-gin-latihan2/middleware"
	"belajar-go-restapi-gin-latihan2/repository"
	"belajar-go-restapi-gin-latihan2/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	db := app.NewDB()

	repositoryGuru := repository.NewGuruRepositoryImpl()
	serviceGuru := service.NewGuruServiceImpl(db, repositoryGuru)
	router.Use(middleware.AuthMiddleware)
	router.GET("/api/guru", serviceGuru.FindAll)
	router.POST("/api/guru", serviceGuru.Create)
	router.GET("/api/guru/:IdGuru", serviceGuru.FindById)
	router.PUT("/api/guru/:IdGuru", serviceGuru.Update)
	router.DELETE("/api/guru/:IdGuru", serviceGuru.Delete)

	router.Run(":9000")
}
