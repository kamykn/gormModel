package address

import ( 
	"../model"
)

type Address struct {
    ID       int
    Address1 string         `gorm:"type:varchar(100);not null;unique"` // Set field as not nullable and unique
    Address2 string         `gorm:"type:varchar(100);unique"`
}

func NewAddress () Address{
	address := Address{}
	return address
}

func (address *Address) Add (address1, address2 string) {
	address.Address1 = address1
	address.Address2 =  address2

	model.DB.NewRecord(address)
	model.DB.Create(&address)
}

func (address *Address) First() *Address {
	model.DB.First(&address)
	return address
}




