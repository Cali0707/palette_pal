package controller

import (
	"fmt"
	"github.com/Cali0707/palette_pal/db"
	"github.com/Cali0707/palette_pal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserHandler(db db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//fmt.Printf("/users/create")
		var newUser models.User
		//
		if c.Bind(&newUser) == nil {
			err := models.CreateUser(newUser, db.DB)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong on our end, please try again later!"})
			}

			c.JSON(http.StatusOK, gin.H{"message": "Welcome to Palette Pals!"})
		}
	}
}

func GetUserHandler(db db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Please provide a username!"})
			return
		}

		fmt.Printf(username)

		user, err := models.GetUserByUsername(username, db.DB)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Sorry, we could not find a user with that username!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"username": user.UserName,
			"email":    user.Email,
		})
	}
}
