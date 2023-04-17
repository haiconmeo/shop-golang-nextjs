package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProductMedia struct {
	gorm.Model
	Image     string `json:"image" gorm:"image"`
	Position  int    `json:"position" gorm:"position"`
	ProductId int    `json:"productId" gorm:"product_id"`
}
type Product struct {
	gorm.Model
	Name        string `json:"name" gorm:"name"`
	Description string `json:"description" gorm:"description"`
	Price       uint   `json:"price" gorm:"price"`
	Media       []ProductMedia
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("err db")
	}
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&ProductMedia{})
	r := gin.Default()
	r.GET("/product", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "oke",
		})
	})
	r.Run()
}
