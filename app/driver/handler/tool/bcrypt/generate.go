package bcrypt

import (
	fiber "github.com/gofiber/fiber/v2"

	security "app/app/driver/core/security"
	apidoc "app/apidoc"
)

type GenerateReqData struct {
     Password  string    `json:"password"`
}
type GenerateReq struct {
    Data GenerateReqData `json:"data"`
}

type GenerateResStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type GenerateResData struct {
	HashedPassword   string `json:"hashed_password"`
}
type GenerateRes struct {
	Status GenerateResStatus `json:"status"`
	Data   *GenerateResData   `json:"data,omitempty"`
}

func GenerateApp(app_router *fiber.App, parent_slug, parent_group string) {
	app := app_router
	summary := "Generate a bcrypt hashed password"
	slug := parent_slug + "/generate"
	group := parent_group

	apidoc.AddWithModels(app, fiber.MethodPost, slug, summary, GenerateFlow,
		GenerateReq{}, GenerateRes{}, // Request Format, Request Format
		map[string]interface{}{ // Response Example
			"200": GenerateRes{
				Status: GenerateResStatus{
					Code:    200,
					Message: "Success",
				},
				Data: &GenerateResData{
					HashedPassword: "string",
				},
			},
			"400": GenerateRes{
				Status: GenerateResStatus{
					Code:    400,
					Message: "Something went wrong",
				}, Data: nil,
			},
			"401": GenerateRes{
				Status: GenerateResStatus{
					Code:    401,
					Message: "Something went wrong",
				}, Data: nil,
			},
		},
		group,
	)
}

func GenerateFlow(c *fiber.Ctx) error {
	var req GenerateReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}

	password := req.Data.Password
	hashedpassword, err := security.BcryptGenerate(password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to hash password"})
	}

	return c.JSON(GenerateRes{
		Status: GenerateResStatus{
			Code:    200,
			Message: "Success",
		},
		Data: &GenerateResData{
			HashedPassword: hashedpassword,
		},
	})
}
