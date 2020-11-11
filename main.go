package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Car struct {
	Gear  uint   `json:"gear"`
	Id    uint   `json:"id"`
	Model string `json:"model"`
}

func getAllCars(c *gin.Context) {
	var car_from_db []Car
	err := db.Find(&car_from_db).Error
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, car_from_db)
	}
}

func getCar(c *gin.Context) {
	var car Car
	id := c.Params.ByName("id")
	err := db.Where("id = ?", id).First(&car).Error
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)

	} else {
		c.JSON(200, car)
	}
}
func addCar(c *gin.Context) {
	var car Car
	c.BindJSON(&car)
	db.Create(&car)
	c.JSON(200, car)
	db.Save(&car)

}
func updateCar(c *gin.Context) {
	var car Car
	id := c.Params.ByName("id")
	err := db.Where("id = ?", id).First(&car).Error
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)

	} else {
		c.BindJSON(&car)
		db.Save(&car)
		c.JSON(200, car)
	}
}
func deleteCar(c *gin.Context) {
	id := c.Params.ByName("id")
	var car Car
	err := db.Where("id  = ? ", id).Delete(&car).Error
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{
			"Deleted": "Deleted with id " + string(id)})
	}
}
func main() {
	//
	//
	//
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Fatal("error")
	}
	defer db.Close()
	db.AutoMigrate(&Car{})
	r := gin.Default()
	r.GET("/", getAllCars)
	r.GET("/cars/:id", getCar)
	r.POST("/addingcar", addCar)
	r.PUT("/updatecar/:id", updateCar)
	r.DELETE("/delete/:id", deleteCar)
	r.Run(":8000")

	// var c3 Car
	// db.First(&c3)
	// fmt.Println(c1.Model)
	// fmt.Println(c2.Model)
	// fmt.Println(c3.Model)
}
