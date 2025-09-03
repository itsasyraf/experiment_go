package bcrypt

import (
	fiber "github.com/gofiber/fiber/v2"
)

func MainApp(app_router *fiber.App, parent_slug, parent_group string) {
	app := app_router
	slug := parent_slug + "/tool/bcrypt"
	group := parent_group + "/tool"

	GenerateApp(app, slug, group)
    VerifyApp(app, slug, group)
}
