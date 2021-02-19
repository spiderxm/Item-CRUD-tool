package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm-practice/models"
	"gorm-practice/services"
)

func initDB() {
	var err error
	services.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection successfully opened.")

	services.DBConn.AutoMigrate(&models.Item{})
	fmt.Println("Database Migrated")
}

func menu() {
	heading := "** Welcome to Items CRUD functionality ** "
	option1 := "1.) Add Item"
	option2 := "2.) Delete Item"
	option3 := "3.) Retrieve Item"
	option4 := "4.) Retrieve All Items"
	option5 := "5.) Dump All Items in items.json file"

	fmt.Println(heading)
	fmt.Println(option1)
	fmt.Println(option2)
	fmt.Println(option3)
	fmt.Println(option4)
	fmt.Println(option5)
}
func main() {
	// Initialize Database
	initDB()

}
