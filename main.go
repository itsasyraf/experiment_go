package main

import (
	fmt "fmt"
	core "app/core"
	handler "app/handler"
 	fiber "github.com/gofiber/fiber/v2"
)

type ResponseStatus struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}
type ResponseData struct {
    UIMessage   string    `json:"ui_message"`
}
type Response struct {
	Status ResponseStatus `json:"status"`
    Data ResponseData `json:"data"`
}

func main() {
	core.InitEnv()
	port := core.EnvMandatory("PROJECT_PORT")

	app := fiber.New()
	handler.main.app(app)

	fmt.Println("🚀 Server running on :"+port)
	app.Listen(":"+port)
}
