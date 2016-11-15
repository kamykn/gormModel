package main

import ( 
	"fmt"
	"./models"
	"./migration"
)

func main() {
	migrationDown()
	migrationUp()
	addAddress()
	addUser()

}

func addAddress () {
	addressModel := models.NewAddressModel()
	addressModel.Add("Tokyo", "Setagaya")
	fmt.Println(addressModel.First())
}

func addUser() {
	userModel := models.NewUserModel()
	userModel.Add("kmszk", 27)
}

func migrationUp () {
	migration.UpAddressTable()
	migration.UpUserTable()
}

func migrationDown () {
	migration.DownAddressTable()
	migration.DownUserTable()
}

