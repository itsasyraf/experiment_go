package user

import (
	"log"
	"os"
	"time"

	auth "app/handler/user/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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

func MainApp(app_router fiber.Router) {
	app := app_router
	app.Post("/user/login", GetRoot)
	// Register route groups
    // routes.UserRoutes(api)
}

func Login(c *fiber.Ctx) error {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}

	// TODO: validate user from DB; for demo assume ok
	token, err := auth.CreateAccessToken("user-123", "admin", 15*time.Minute)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create token"})
	}
	return c.JSON(fiber.Map{"access_token": token, "token_type": "bearer", "expires_in": 900})
}
