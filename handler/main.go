package handler

import (
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

func app(app_router fiber.Router) {
	app := app_router
	app.Get("/", root)
	// Register route groups
    // routes.UserRoutes(api)
}

func root(c *fiber.Ctx) error {
	rs := ResponseStatus{200, "Success"}
	rd := ResponseData{"Welcome to the 'experiment' project developed and deployed in Go language!"}
	result := Response{rs, rd}
	return c.JSON(result)
}
