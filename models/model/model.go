package model

import ( 
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	*gorm.DB
}

var sharedInstance *Model = NewModel()

func NewModel() *Model {
	model := new(Model)

	// データベースの初期化
	dbSetting := USER_NAME + ":" + PASSWORD + "@tcp(" + IP_ADDRESS + ":" + PORT + ")/" + DB_NAME
	db, err := gorm.Open("mysql", dbSetting)

	if err != nil {
		panic(err.Error())
	}

	model.DB = db

	return model
}

func GetInstance() *Model {
	return sharedInstance
}

