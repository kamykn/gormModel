package models

import ( 
	"./model"
)

type User struct {
	ID	int
	Name	string	`gorm:"type:varchar(100);not null"`
	Age	int	`gorm:"not null"`
}

type UserModel struct {
	*model.Model
	User
}

func NewUserModel () *UserModel {
	userModel := new(UserModel)
	userModel.Model = model.NewModel()
	return userModel
}

func (model *UserModel) Add (name string, age int) {
	model.User.Name = name
	model.User.Age =  age

	model.NewRecord(model.User)
	model.Create(&model.User)
}

func (model *UserModel) First() User {
	model.DB.First(&model.User)
	return model.User
}



