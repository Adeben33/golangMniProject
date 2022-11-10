package controllers

import (
	"github.com/adeben33/golangMiniProject/Todo/initializers"
	"github.com/adeben33/golangMiniProject/Todo/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func Signup(c *gin.Context) {
	//Bind the email & password to a struct
	var userStruct struct {
		Email    string
		Password string
	}
	err := c.Bind(&userStruct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	//	Harsh the password
	harsh, err := bcrypt.GenerateFromPassword([]byte(userStruct.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to harsh paswword",
		})
		return
	}
	//	Create the user in the database
	user := models.User{
		Email:    userStruct.Email,
		Password: string(harsh),
	}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Create user",
		})
		return
	}
	c.JSON(200, gin.H{})
}

func Login(c *gin.Context) {
	//Get the email and password
	var userStruct struct {
		Email    string
		Password string
	}
	err := c.Bind(&userStruct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	//	Look for requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", userStruct.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//	compare sent in pass with saved user pass hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userStruct.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect Password",
		})
	}
	//	generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Token not generated",
		})
	}
	//	save it in a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(200, gin.H{})
}

func Logout(c *gin.Context) {
	_, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user not logged in",
		})
	}
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(200, gin.H{
		"info": "logout successful",
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(200, gin.H{
		"user":  user,
		"Hello": "Hello",
	})
}
