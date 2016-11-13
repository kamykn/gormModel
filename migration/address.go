package migration

import ( 
	"../models"
)

func Up() {
	addressModel := models.NewAddressModel()
	addressModel.CreateTable(models.Address{})
	addressModel.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(models.Address{})
}

func Down() {
	address := models.Address{}

	addressModel := models.NewAddressModel()
	addressModel.DropTable(address)
}


