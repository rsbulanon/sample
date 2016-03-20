package main

import (
	"os"
	"fmt"

	 "github.com/gin-gonic/gin"
	 _ "github.com/go-sql-driver/mysql"
	 h "test/sample/api/handlers"
	 "gopkg.in/mgo.v2"
	 //"test/sample/api/config"
)

func main() {
	db := *InitDB()
	router := gin.Default()
	LoadAPIRoutes(router, &db)
}

func LoadAPIRoutes(r *gin.Engine, db *mgo.Session) {
	public := r.Group("/api/v1")

	userHandler := h.NewUserHandler(db)
	public.GET("/users", userHandler.Index)
	public.POST("/users", userHandler.Create)
	public.POST("/auth", userHandler.Auth)
	var port = os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Println("PORT ---> ",port)
	r.Run(fmt.Sprintf(":%s", port))
}

func InitDB() *mgo.Session {
	//sess, err := mgo.Dial("mongodb://localhost/sampledb")
	sess, err := mgo.Dial("mongodb://rsbulanon:Passw0rd@ds019829.mlab.com:19829/sampledb")
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database:  %s", err))
	}
	sess.SetSafe(&mgo.Safe{})
	//_db.DB()
	//_db.LogMode(true)
	//_db.DB().Ping()
	//_db.DB().SetMaxIdleConns(10)
	//_db.DB().SetMaxOpenConns(100)

	//_db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&m.User{})

	return sess
}

func GetPort() string {
    var port = os.Getenv("PORT")
    // Set a default port if there is nothing in the environment
    if port == "" {
        port = "8000"
        fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
    }
    fmt.Println("port -----> ", port)
    return ":" + port
}