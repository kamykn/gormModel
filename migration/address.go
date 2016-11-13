package migration

import ( 
	"../model"
)

type Address struct {
    ID       int
    Address1 string         `gorm:"type:varchar(100);not null;unique"` // Set field as not nullable and unique
    Address2 string         `gorm:"type:varchar(100);unique"`
}

func Up() {
	model.DB.CreateTable(&Address{})
	model.DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Address{})
}


