package model

import ( 
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	USER_NAME = "root"
	PASSWORD = "qD8)BL(S"
	IP_ADDRESS = "192.168.99.100"
	PORT = "3306"
	DB_NAME = "revel-myapp"
)

type Model struct {
	*gorm.DB
}

func NewModel() *Model {
	model := new(Model)

	if model.DB != nil {
		return model
	}


	// データベースの初期化
	dbSetting := USER_NAME + ":" + PASSWORD + "@tcp(" + IP_ADDRESS + ":" + PORT + ")/" + DB_NAME
	db, err := gorm.Open("mysql", dbSetting)

	if err != nil {
		panic(err.Error())
	}

	model.DB = db

	return model
}


