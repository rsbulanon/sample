package handlers

import (
	"net/http"
	"strconv"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	m "test/sample/api/models"
	"gopkg.in/mgo.v2"	
	"gopkg.in/mgo.v2/bson"	
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
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
	collection := handler.sess.DB("sampledb").C("users") 
	collection.Find(nil).All(&users)
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
	collection := handler.sess.DB("sampledb").C("users") 
	result := m.User{}
	err := collection.Find(bson.M{"email": user.Email}).One(&result)
	//check if email is not existing
	if fmt.Sprintf("%s", err) == "not found" {
		// generate hashed password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err == nil {
			user.ID = fmt.Sprintf("%s", uuid.NewV4())
			user.CreatedAt = time.Now().UTC()
			user.UpdatedAt = time.Now().UTC()
			user.Password = string(hashedPassword)
			collection.Insert(&user)
		    // Create the token
		    token := jwt.New(jwt.SigningMethodHS256)
		    token.Claims["id"] = user.ID
		    token.Claims["iat"] = time.Now().Unix()
		    token.Claims["exp"] = time.Now().Add(time.Second * 3600 * 24).Unix()
		    tokenString, err := token.SignedString([]byte("secret"))
		    if err == nil {
	    		resp := map[string]string{"token": tokenString}
				c.JSON(http.StatusCreated,resp)	
	    	} else {
    			fmt.Println("failed to create token --->",err)
				respondWithError(http.StatusBadRequest,"Failed to create account",c)
	    	}
		} else {
			fmt.Println("failed to encrypt password",err)
		}
	} else {
		fmt.Println("email existing")
		respondWithError(http.StatusBadRequest,"Email already taken",c)
	}
}




