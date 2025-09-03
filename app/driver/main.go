package driver

import (
 	fiber "github.com/gofiber/fiber/v2"

 	environment "app/app/driver/core/environment"
 	migration "app/app/driver/migration"
  	handler "app/app/driver/handler"
)

func Init() {
	environment.Init()
	migration.Common()
	migration.JDDriverLoginIdType()
	migration.JDDriver()
}

func Main(app_router *fiber.App) {
	app := app_router
	handler.MainApp(app)
}
