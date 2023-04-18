package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProductMedia struct {
	gorm.Model
	Image     string `json:"image" gorm:"image"`
	Position  int    `json:"position" gorm:"position"`
	ProductID uint   `json:"productId" gorm:"product_id"`
}

type Product struct {
	gorm.Model
	Name        string         `json:"name" gorm:"name"`
	Description string         `json:"description" gorm:"description"`
	Price       uint           `json:"price" gorm:"price"`
	Media       []ProductMedia `json:"media" gorm:"foreignKey:ProductID"`
}
type Page struct {
	Limit  int    `form:"limit"`
	Page   int    `form:"page"`
	Filter string `form:"filter"`
}

type ResponseTemplate struct {
	data   []interface{}
	count  int
	limit  int
	page   int
	filter string
}

func createProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data Product
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "imput error",
			})
			return
		}
		media := []ProductMedia{}
		for _, value := range data.Media {
			m := ProductMedia{Image: value.Image, Position: value.Position}
			if err := db.Create(&m).Error; err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "failed to create tag",
				})
				return
			}
			media = append(media, m)
		}
		data.Media = media
		db.Create(&data)
		ctx.JSON(200, gin.H{
			"data": data,
		})

	}
}
func getOneProductById(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product Product
		if err := db.Preload("Media").Where("id = ?", ctx.Param("id")).First(&product).Error; err != nil {
			ctx.JSON(400, gin.H{
				"message": "not found",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"data": product,
		})
	}
}
func getAllProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var page Page
		ctx.Bind(&page)
		// fmt.Println("manh", page.Limit)
		if page.Page == 0 {
			page.Page = 1
		}
		var products []Product
		repo := db.Limit(page.Limit).Offset(page.Page)
		if err := repo.Find(&products).Error; err != nil {
			ctx.JSON(400, gin.H{
				"message": "error",
			})
			return
		}
		var count int64
		repo.Count(&count)
		ctx.JSON(200, gin.H{
			"data":   products,
			"count":  count,
			"limit":  page.Limit,
			"page":   page.Page,
			"filter": page.Filter,
		})
		// ctx.JSON(200)

	}
}
func updateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func deleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("err db")
	}
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&ProductMedia{})
	r := gin.Default()
	r.GET("/products", getAllProduct(db))
	r.POST("/products", createProduct(db))
	r.GET("/products/:id", getOneProductById((db)))
	r.Run()
}
