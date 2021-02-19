package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm-practice/models"
	"os"
)

var (
	DBConn *gorm.DB
)

func AddItem() {
	var (
		Name        string
		Description string
		Price       float64
		Quantity    int64
	)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Item Name : ")
	Name, _ = reader.ReadString('\n')
	fmt.Print("Enter Item Description : ")
	Description, _ = reader.ReadString('\n')
	fmt.Print("Enter Item Price : ")
	fmt.Scan(&Price)
	fmt.Print("Enter Item Quantity : ")
	fmt.Scan(&Quantity)
	var item models.Item
	item.Name = Name
	item.Description = Description
	item.Price = Price
	item.Quantity = Quantity
	db := DBConn
	db.Create(&item)
	itemInJson, _ := json.MarshalIndent(item, "", "    ")
	fmt.Println("--- Item Details in JSON ---")
	fmt.Println(string(itemInJson))
}

func AllItems() {
	var items []models.Item
	db := DBConn
	db.Find(&items)
	itemsToJson, _ := json.MarshalIndent(items, "", "    ")
	fmt.Println("--- Items Details in Json ---")
	fmt.Println(string(itemsToJson))
}

func FindItem() {

}

func DeleteItem() {

}

func AllItemsToJsonFile() {

}
