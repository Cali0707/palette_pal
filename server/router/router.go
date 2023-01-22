package router

import (
	"github.com/Cali0707/palette_pal/controller"
	"github.com/Cali0707/palette_pal/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitializeRouter(db db.DB) (router *gin.Engine) {
	router = gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world!")
	})

	usersRouter := router.Group("/users")
	usersRouter.POST("/create", controller.CreateUserHandler(db))
	usersRouter.GET("", controller.GetUserHandler(db))

	recipesRouter := router.Group("/recipes")
	recipesRouter.POST("/create", controller.CreateRecipeHandler(db))

	postsRouter := router.Group("/posts")
	postsRouter.POST("/create", controller.CreatePostHandler(db))

	return
}
