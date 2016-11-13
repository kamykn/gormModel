package models

import ( 
	"./model"
)

type Address struct {
	ID       int
	Address1 string         `gorm:"type:varchar(100);not null;unique"` // Set field as not nullable and unique
	Address2 string         `gorm:"type:varchar(100);unique"`
}

type AddressModel struct {
	*model.Model
}

func NewAddressModel () *AddressModel {
	addressModel := new(AddressModel)
	addressModel.Model = model.NewModel()
	return addressModel
}

func (model *AddressModel) Add (address1, address2 string) {
	address := Address{}
	address.Address1 = address1
	address.Address2 =  address2

	model.NewRecord(address)
	model.Create(&address)
}

func (model *AddressModel) First() Address {
	address := Address{}
	model.DB.First(&address)
	return address
}


