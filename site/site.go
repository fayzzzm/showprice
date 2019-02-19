package site

// Imported Packages
import (
	"billing/site/configs"
	"billing/site/db"
	"billing/site/router"
	"billing/site/server"
)

func Run() {
	// Configure Configs
	if err := configs.Init(); err != nil {
		panic(err)
	}

	// Connect Db
	if err := db.Init(); err != nil {
		panic(err)
	}

	defer db.Exit()
	// Connect Router
	if err := router.Init(); err != nil {
		panic(err)
	}
	// Configure Server
	if err := server.Init(); err != nil {
		panic(err)
	}
}
