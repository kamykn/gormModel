package main

import ( 
	"fmt"
	"./address"
	_ "./migration"
)

func main() {
	//fmt.Println(model.DB.HasTable("new_table"))

	addObj := address.NewAddress()
	addObj.Add("chiba1", "ooami1")
	fmt.Println(addObj.First())

	//migration.Up()
}

