package controller

import (
	"github.com/Cali0707/palette_pal/db"
	"github.com/Cali0707/palette_pal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePostHandler(db db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json struct {
			Recipe     models.Recipe `json:"recipe"`
			User       models.User   `json:"user"`
			Post       models.Post   `json:"post"`
			PhotoLinks []string      `json:"photoLinks"`
		}

		err := c.Bind(&json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Please sent a properly formed request body with all the necessary fields!"})
			return
		}

		postId, err := models.CreatePostForRecipe(json.Recipe, json.User, json.Post, json.PhotoLinks, db.DB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong on our end, please try again later!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Success!", "postId": postId})
	}
}
