# gormModel

## Settings
Rename config.go.example to config.go.

```
$ cd models/model
$ mv config.go.example config.go
```

Then, change the database settings.

```
package model

const (
        USER_NAME = "myuser"
        PASSWORD = "**********"
        HOST = "111.111.111.111"
        PORT = "3306"
        DB_NAME = "mydatabase"
)
```

## Add Model

models/model/sampleUser.go
```
package models

import (
        "./model"
)

type User struct {
        ID      int
        Name    string  `gorm:"type:varchar(100);not null"`
        Age     int     `gorm:"not null"`
}

type UserModel struct {
        *model.Model
        User
}

func NewUserModel () *UserModel {
        userModel := new(UserModel)

        userModel.Model = model.GetInstance()
        return userModel
}
```


