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
	option1 := "1.) Add Item. -- Press 1 -- "
	option2 := "2.) Delete Item. -- Press 2 -- "
	option3 := "3.) Retrieve Item. -- Press 3 -- "
	option4 := "4.) Retrieve All Items. -- Press 4 -- "
	option5 := "5.) Dump All Items in items.json file. -- Press 5 -- "
	option6 := "6.) Update Item. -- Press 6 -- "
	option7 := "7.) Press any other key to exit. -- Press any other key -- "

	fmt.Println(heading)
	fmt.Println(option1)
	fmt.Println(option2)
	fmt.Println(option3)
	fmt.Println(option4)
	fmt.Println(option5)
	fmt.Println(option6)
	fmt.Println(option7)
	fmt.Print("Enter Your Choice : ")
}
func main() {
	// Initialize Database
	initDB()
	var option string
	loop := true
	for loop == true {
		menu()
		fmt.Scan(&option)
		switch option {
		case "1":
			services.AddItem()
		case "2":
			services.DeleteItem()
		case "3":
			services.FindItem()
		case "4":
			services.AllItems()
		case "5":
			services.AllItemsToJsonFile()
		case "6":
			services.UpdateItem()
		default:
			fmt.Println("Thanks for using Items CRUD Functionality tool.")
			loop = false
		}
	}
}
