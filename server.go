package main

import (
	"fmt"
	 "github.com/gin-gonic/gin"
	 "github.com/jinzhu/gorm"
	 _ "github.com/go-sql-driver/mysql"
	 m "test/sample/api/models"
	 h "test/sample/api/handlers"
)

func main() {
	db := *InitDB()
	router := gin.Default()
	LoadAPIRoutes(router, &db)
}

func LoadAPIRoutes(r *gin.Engine, db *gorm.DB) {
	public := r.Group("/api/v1")

	userHandler := h.NewUserHandler(db)
	public.GET("/users", userHandler.Index)
	public.POST("/users", userHandler.Create)
	r.Run(":8080")
}

func InitDB() *gorm.DB {
	_db, err := gorm.Open("mysql", "root:@/sampledb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database:  %s", err))
	}
	_db.DB()
	//_db.LogMode(true)
	//_db.DB().Ping()
	//_db.DB().SetMaxIdleConns(10)
	//_db.DB().SetMaxOpenConns(100)

	_db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&m.User{})

	return &_db
}