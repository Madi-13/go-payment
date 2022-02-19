package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type payment struct {
	ID         int64   `json:"id" gson:"primary key"`
	Name       string  `json:"name" `
	Price      float32 `json:"price"`
	Date       string  `json:"date"`
	Type       string  `json:"type"`
	Comment    string  `json:"comment"`
	CategoryID int64   `json:"category_id`
	Category   Category
}

type Category struct {
	ID   int64 `gson:"primary key"`
	Name string
}

func getPayments(c *gin.Context) {
	dsn := "host=localhost port=5432 user=madi dbname=postgres password=maditoto sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var payment []payment
	db.Preload("Category").Find(&payment)
	c.IndentedJSON(http.StatusOK, &payment)
}

func postPayments(c *gin.Context) {
	dsn := "host=localhost port=5432 user=madi dbname=postgres password=maditoto sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var newPayment payment

	if err := c.BindJSON(&newPayment); err != nil {
		return
	}

	db.Save(&newPayment)
	c.IndentedJSON(http.StatusCreated, newPayment)
}

func getPaymentByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}
	dsn := "host=localhost port=5432 user=madi dbname=postgres password=maditoto sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var p payment
	db.Preload("Category").Find(&p, id)
	c.IndentedJSON(http.StatusOK, &p)
}

func deletePayment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	dsn := "host=localhost port=5432 user=madi dbname=postgres password=maditoto sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Delete(&payment{}, id)
	c.IndentedJSON((http.StatusOK), gin.H{"message": "payment delete successfully"})
}

func main() {
	router := gin.Default()
	router.GET("/payments", getPayments)
	router.GET("/payments/:id", getPaymentByID)
	router.POST("/payments", postPayments)
	router.DELETE("/payments/:id", deletePayment)
	router.Run("localhost:8080")
}
