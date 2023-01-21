package main

import (
	"fmt"
	"github.com/Cali0707/palette_pal/db"
	"github.com/Cali0707/palette_pal/pkg/config"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
)

func main() {
	//postgresUser := os.Getenv()
	port := ":3000"

	conf, err := config.LoadConfig()

	if err != nil {
		println("Something went wrong!")
		os.Exit(1)
	}

	err = db.RunMigrationsClean(conf)
	if err != nil {
		fmt.Printf("Failed to run migrations, exiting...")
		os.Exit(1)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})

	err = r.Run(port)
	if err != nil {
		println("Failed to run app")
	}
}
