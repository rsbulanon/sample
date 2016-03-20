package handlers

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
	m "test/sample/api/models"
	 "gopkg.in/mgo.v2"	
	//"github.com/dgrijalva/jwt-go"
	//"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	sess *mgo.Session
}

// NewAppoointment factory for AppointmentsController
func NewUserHandler(sess *mgo.Session) *UserHandler {
	return &UserHandler{sess}
}

// Index retrieves a list of users
func (handler UserHandler) Index(c *gin.Context) {
	start := -1
	max := 10

	//check if start exists in url parameters
	if c.Query("start") != ""  {
		i,_ := strconv.Atoi(c.Query("start"))
		start = i;
	} else {
		fmt.Println("cant read start query param")
	}

	if c.Query("max") != ""  {
		i,_ := strconv.Atoi(c.Query("max"))
		max = i;
	} 

	fmt.Printf("offset ---> %d max ---> %d\n", start, max)
	users := []m.User{}
	//handler.db.Where("deleted_at is null AND status = ?","active").Limit(max).Offset(start).Order("created_at desc").Find(&users)
	c.JSON(http.StatusOK, users)
}

// Show retrieves a user record with filters
func (handler UserHandler) Show(c *gin.Context) {
	//id := c.Param("id")
	user := []m.User{}
	//handler.db.Where("deleted_at is null AND status = ? AND ID = ?","active",id).Order("created_at desc").Limit(20).Find(&user)
	c.JSON(http.StatusOK, user)
}

// Create an appointment
func (handler UserHandler) Create(c *gin.Context) {
	user := m.User{}
	c.BindJSON(&user)

	//check if email is not existing
	handler.sess.DB("sampledb").C("users").Insert(user)
	c.JSON(http.StatusCreated,&user)
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.AbortWithStatus(code)
}







