package handler

import (
	fiber "github.com/gofiber/fiber/v2"

	toolbcrypt "app/app/driver/handler/tool/bcrypt"
	apidoc "app/apidoc"
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
    Data *mdlResData `json:"data"`
}

func MainApp(app_router *fiber.App) {
	app := app_router
	summary := "This is the first message when opening the driver app's API"
	slug := "/driver-app"
	group := "/driver-app"

	toolbcrypt.MainApp(app, slug, group)

	apidoc.AddWithModels(app, fiber.MethodGet, slug, summary, Root,
		nil, mdlRes{}, // Request Format, Request Format
		map[string]interface{}{ // Response Example
			"200": mdlRes{
				Status: mdlResStatus{
					Code:    200,
					Message: "Success",
				},
				Data: &mdlResData{
					UIMessage: "string",
				},
			},
		},
		group,
	)
}

func Root(c *fiber.Ctx) error {
	return c.JSON(mdlRes{
		Status: mdlResStatus{
			Code:    200,
			Message: "Success",
		},
		Data: &mdlResData{
			UIMessage: "Welcome to the 'experiment' project developed and deployed in Go language!",
		},
	})
}
