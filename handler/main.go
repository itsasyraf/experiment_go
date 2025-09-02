package handler

import (
	fiber "github.com/gofiber/fiber/v2"

	user "app/handler/user"
)

type mdlResStatus struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}
type mdlResData struct {
    UIMessage   string    `json:"ui_message"`
}
type mdlRes struct {
	Status mdlResStatus `json:"status"`
    Data mdlResData `json:"data"`
}

func MainApp(app_router fiber.Router) {
	app := app_router
	app.Get("/", Root)
	// Register route groups
    user.MainApp(app)
}

// @Summary      Welcome message
// @Description  This endpoint used to welcome users who are visiting the http://domain:port/ directly
// @Tags       	 /
// @Accept       json
// @Produce      json
// @Router       / [get]
func Root(c *fiber.Ctx) error {
	return c.JSON(mdlRes{
		Status: mdlResStatus{
			Code:    200,
			Message: "Success",
		},
		Data: mdlResData{
			UIMessage: "Welcome to the 'experiment' project developed and deployed in Go language!",
		},
	})
}
