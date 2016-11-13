package main

import ( 
	"fmt"
	"./models"
	"./migration"
)

func main() {
	migrationDown()
	migrationUp()
	add()
}

func add () {
	addressModel := models.NewAddressModel()
	addressModel.Add("Tokyo", "Setagaya")
	fmt.Println(addressModel.First())
}

func migrationUp () {
	migration.Up()
}

func migrationDown () {
	migration.Down()
}


