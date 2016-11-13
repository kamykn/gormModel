package migration

import ( 
	"../model"
	"../address"
)


func Up() {
	address := address.NewAddress()
	model.DB.CreateTable(address)
	model.DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(address)
}

func Down() {
	address := address.NewAddress()
	model.DB.DropTable(address)
}


