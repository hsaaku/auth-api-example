package main

import (
	_ "github.com/lib/pq"
	"hsaaku.com/auth-api/app"
	"hsaaku.com/auth-api/config"
)

func main() {
	config := config.GetConfig()

	// messages := []model.Message{
	// 	model.Message{"Hai"},
	// 	model.Message{"Siewuir"},
	// 	model.Message{"fsdkjo safaisf sodjf"},
	// }
	// for _, msg := range messages {
	// 	fmt.Println(msg.Message)
	// }

	a := &app.App{}
	a.Initialize(config)

	a.Run(":4000")
}
