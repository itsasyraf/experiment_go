package user

import (
	// fmt "fmt"
	time "time"

	fiber "github.com/gofiber/fiber/v2"

	auth "app/handler/user/auth"
	database "app/core/database"
)
type mdlReqData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type mdlResStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type mdlResData struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	LoginID      string `json:"login_id"`
	// Do not expose password in real system!
	LoginPassword string `json:"login_password,omitempty"`
	CurrentTime   string `json:"current_time"`
}
type mdlRes struct {
	Status mdlResStatus `json:"status"`
	Data   mdlResData   `json:"data"`
}

func MainApp(app_router fiber.Router) {
	app := app_router
	app.Post("/user/login", Login)
	// Register route groups
    // routes.UserRoutes(api)
}

// @Summary      User Login
// @Description
// @Tags       	 /user
// @Param        payload  body      mdlReqData  true  "Login payload"
// @Accept       json
// @Produce      json
// @Router       /user/login [post]
func Login(c *fiber.Ctx) error {
	var req mdlReqData
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}

	rows, err := database.MainFetch("SELECT (now() at time zone 'UTC')::varchar as current_time")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if len(rows) == 0 {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}

	currentTime, ok := rows[0]["current_time"].(string)
	// var currentTime string
	if !ok {
		return c.Status(500).JSON(fiber.Map{"error": "cannot parse current_time"})
	}

	token, err := auth.CreateAccessToken("user-123", "admin", 15*time.Minute)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create token"})
	}

	return c.JSON(
		mdlRes{
			Status: mdlResStatus{
				Code:    200,
				Message: "Success",
			},
			Data: mdlResData{
				AccessToken:   token,
				TokenType:     "bearer",
				ExpiresIn:     900,
				LoginID:       req.Email,
				LoginPassword: req.Password,
				CurrentTime:   currentTime, // ✅ now is string
			},
		},
	)
}
