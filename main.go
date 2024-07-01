package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"trainRestApi/Controllers"
	"trainRestApi/Database"
)

func main() {
	dsn := "user=musti password=1234 dbname=productdb host=localhost port=5432 sslmode=disable"
	Database.InitDb(dsn)

	router := gin.Default()

	router.GET("/products", Controllers.GetProducts)
	router.POST("/products", Controllers.CreateProduct)
	router.PUT("/products/:id", Controllers.UpdateProduct)
	router.DELETE("/products/:id", Controllers.DeleteProduct)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
