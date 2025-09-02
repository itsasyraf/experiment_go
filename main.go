package main

import (
	fmt "fmt"
	core "app/core"
	sync "sync"
 	fiber "github.com/gofiber/fiber/v2"
 	swagger "github.com/gofiber/swagger"

  	handler "app/handler"
    _ "app/docs"

)

var once sync.Once

func main() {
    once.Do(func() {
	    core.InitDB()
		core.InitEnv()
    })
	port := core.EnvMandatory("PROJECT_PORT")

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	handler.MainApp(app)

	fmt.Println("🚀 Server running on :"+port)
	app.Listen(":"+port)
}
