package main

import ( 
	"fmt"
	"./model"
	"./address"
	"./migration"
)

func main() {
	//exists()
	add()
	// migrationUp()
	// migrationDown()
}

func exists() {
	fmt.Println(model.DB.HasTable("new_table"))
}

func add () {
	newAddress := address.NewAddress()
	newAddress.Add("Tokyo", "Setagaya")
	fmt.Println(newAddress.First())
}

func migrationUp () {
	migration.Up()
}

func migrationDown () {
	migration.Down()
}


