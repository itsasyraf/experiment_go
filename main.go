package main

import (
	fmt "fmt"
	sync "sync"

 	fiber "github.com/gofiber/fiber/v2"
 	swagger "github.com/gofiber/swagger"

 	// database "app/core/database"
 	environment "app/core/environment"
   	_ "app/docs"
  	handler "app/handler"
)

var once sync.Once

func main() {
    once.Do(func() {
    	environment.Init()
     	// database.MainInit()
	    // core.InitDB()
    })
	port := environment.Mandatory("PROJECT_PORT")

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	handler.MainApp(app)

	fmt.Println("🚀 Server running on :"+port)
	app.Listen(":"+port)
}
