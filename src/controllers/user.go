package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/tianmarillio/instant-note-backend/config"
	"github.com/tianmarillio/instant-note-backend/src/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	// get & bind request body
	var body *AuthBody
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// hash password
	passByte := []byte(body.Password)
	hash, err := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	// create user
	user := models.User{
		ID:       uuid.NewString(),
		Email:    body.Email,
		Password: string(hash),
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, result.Error)
		return
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"expired_at": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// return json token
	c.JSON(http.StatusCreated, gin.H{
		"userToken": tokenString,
	})
}

func Login(c *gin.Context) {
	// get & bind request body
	var body *AuthBody
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// get user by email
	var user models.User
	result := config.DB.
		Where("email = ?", body.Email).
		First(&user)
	if result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	// verify password
	// err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(body.Password),
	); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"expired_at": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Sign and get the complete encoded token as a string using the secret
	c.JSON(http.StatusAccepted, gin.H{
		"userToken": tokenString,
	})
}

func AuthTest(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusAccepted, gin.H{
		// "status": "ok",
		"user": user,
	})
}
