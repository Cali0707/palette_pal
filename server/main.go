package main

import (
	"fmt"
	"github.com/Cali0707/palette_pal/db"
	"github.com/Cali0707/palette_pal/pkg/config"
	"github.com/Cali0707/palette_pal/router"
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

	DB, err := db.Connect(conf)

	r := router.InitializeRouter(DB)

	err = r.Run(port)
	if err != nil {
		println("Failed to run app")
	}
}
