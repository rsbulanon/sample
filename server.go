package main

import (
 "github.com/gin-gonic/gin"
)

type User struct {
 	Id int64 `db:"id" json:"id"`
 	Firstname string `db:"firstname" json:"firstname"`
 	Lastname string `db:"lastname" json:"lastname"`
}

func main() {
	r := gin.Default()
	v1 := r.Group("api/v1")
	v1.GET("/users", GetUsers)
	r.Run(":8080")
}

func GetUsers(c *gin.Context) {
 	type Users []User
	var users = Users {
 		User {Id: 1, Firstname: "Oliver", Lastname: "Queen"},
 		User {Id: 2, Firstname: "Malcom", Lastname: "Merlyn"},
 	}
	c.JSON(200, users)
}