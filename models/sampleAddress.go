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
	Address
}

func NewAddressModel () *AddressModel {
	addressModel := new(AddressModel)
	addressModel.Model = model.NewModel()
	return addressModel
}

func (model *AddressModel) Add (address1, address2 string) {
	model.Address.Address1 = address1
	model.Address.Address2 =  address2

	model.NewRecord(model.Address)
	model.Create(&model.Address)
}

func (model *AddressModel) First() Address {
	model.DB.First(&model.Address)
	return model.Address
}


