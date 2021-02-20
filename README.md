# Items Crud functionality tool
## Made with [GoLang](https://golang.org) and [Gorm](https://gorm.io)
### Database used `sqlite3`

Various Functionality Provided for `Item` table.
Definition for `Item` table is as follows:
```go
type Item struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
}
```

Various Functionality provided are as follows:
##### `1.) Add Item`
##### `2.) Update Item`
##### `3.) Delete Item`
##### `4.) List all Items`
##### `5.) Retrieve Item`
##### `6.) Retrieve All Items in JSON format(in items.json file in same directory)`

Made by [Mrigank Anand](https://github.com/spiderxm)