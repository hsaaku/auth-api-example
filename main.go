package main

import (
	_ "github.com/lib/pq"
	"hsaaku.com/auth-api/app"
	"hsaaku.com/auth-api/config"
)

func main() {
	config := config.GetConfig()

	a := &app.App{}
	a.Initialize(config)

	// err := a.DB.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// a.Run(":4000")
}
