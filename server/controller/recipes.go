package controller

import (
	"github.com/Cali0707/palette_pal/db"
	"github.com/Cali0707/palette_pal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRecipeHandler(db db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json struct {
			Recipe     models.Recipe `json:"recipe"`
			User       models.User   `json:"user"`
			Post       models.Post   `json:"post"`
			PhotoLinks []string      `json:"photoLinks"`
		}

		err := c.Bind(&json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Please send a valid request body"})
			return
		}

		postId, err := models.CreateRecipeWithPhotosAndPost(json.Recipe, json.User, json.Post, json.PhotoLinks, db.DB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong on our end, please try again later!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully posted new recipe!",
			"postId":  postId,
		})
	}
}
