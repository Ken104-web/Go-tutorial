package models

import (
	models "_/home/ken/dev/Go-tutorial/one-to-many"

	"github.com/go-gorm/gorm"
)

type User struct {
    gorm.Model
    Name string
    Posts []Post  `gorm:"foreignKey:UserID"` // one to many relationships
}

type Post struct {
    gorm.Model
    userID uint
    User  User // one to one with User
}

user := models.User{
    Name: "Alice",
    Posts: []models.Post{
        {Content: "First post content"},
        {Content: "second post content"},
    },
}
db.Create(&user)

var user models.User
db.Preload("Posts").First(&user, 1)


