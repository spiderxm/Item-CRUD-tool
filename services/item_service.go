package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm-practice/models"
	"io/ioutil"
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
	var item models.Item
	db := DBConn
	var id int
	fmt.Print("Enter id to find Record : ")
	fmt.Scan(&id)
	result := db.Find(&item, id)
	if result.Error != nil {
		fmt.Println("--- ", result.Error.Error(), " ---")
	} else {
		itemInJson, _ := json.MarshalIndent(item, "", "    ")
		fmt.Println("--- Item Details in JSON ---")
		fmt.Println(string(itemInJson))
	}
}

func DeleteItem() {
	db := DBConn
	var id int
	fmt.Print("Enter id to find Record : ")
	fmt.Scan(&id)
	var item models.Item
	result := db.Find(&item, id)
	if result.Error != nil {
		fmt.Println("--- ", result.Error.Error(), " ---")
	} else {
		db.Delete(&item)
		itemInJson, _ := json.MarshalIndent(item, "", "    ")
		fmt.Println("--- Item Deleted Successfully ---")
		fmt.Println("--- Item Details in JSON ---")
		fmt.Println(string(itemInJson))
	}

}

func AllItemsToJsonFile() {
	var items []models.Item
	db := DBConn
	db.Find(&items)
	itemsToJson, _ := json.MarshalIndent(items, "", "    ")
	file, err := os.Create("items.json")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		file.Close()
		err = ioutil.WriteFile("items.json", itemsToJson, os.ModeAppend)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("--- Contents Written to File items.json---")
		}
	}
}

func UpdateItem() {
	var item models.Item
	db := DBConn
	var id int
	fmt.Print("Enter id to find Record : ")
	fmt.Scan(&id)
	result := db.Find(&item, id)
	if result.Error != nil {
		fmt.Println("--- ", result.Error.Error(), " ---")

	} else {
		itemInJson, _ := json.MarshalIndent(item, "", "    ")
		fmt.Println("--- Item Details in JSON ---")
		fmt.Println(string(itemInJson))
		var (
			Name        string
			Description string
			Price       float64
			Quantity    int64
		)
		fmt.Println("-- Enter New Details --")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Item Name : ")
		Name, _ = reader.ReadString('\n')
		fmt.Print("Enter Item Description : ")
		Description, _ = reader.ReadString('\n')
		fmt.Print("Enter Item Price : ")
		fmt.Scan(&Price)
		fmt.Print("Enter Item Quantity : ")
		fmt.Scan(&Quantity)
		item.Name = Name
		item.Description = Description
		item.Price = Price
		item.Quantity = Quantity
		db.Save(&item)
		fmt.Println("Item Updated")
		itemInJson, _ = json.MarshalIndent(item, "", "    ")
		fmt.Println("--- Updated Item Details in JSON ---")
		fmt.Println(string(itemInJson))
	}
}
