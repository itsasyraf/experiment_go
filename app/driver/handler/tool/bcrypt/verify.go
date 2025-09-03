package bcrypt

import (
	fiber "github.com/gofiber/fiber/v2"

	security "app/app/driver/core/security"
	apidoc "app/apidoc"
)

type VerifyReqData struct {
     Password  string    `json:"password"`
     HashedPassword string `json:"hashed_password"`
}
type VerifyReq struct {
    Data VerifyReqData `json:"data"`
}

type VerifyResStatus struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message"`
}
type VerifyResData struct {
	IsMatch   bool `json:"is_match"`
}
type VerifyRes struct {
	Status VerifyResStatus `json:"status"`
	Data   *VerifyResData   `json:"data,omitempty"`
}

func VerifyApp(app_router *fiber.App, parent_slug, parent_group string) {
	app := app_router
	summary := "Verify a password against a hashed password"
	slug := parent_slug + "/verify"
	group := parent_group

	apidoc.AddWithModels(app, fiber.MethodPost, slug, summary, VerifyFlow,
		VerifyReq{}, VerifyRes{}, // Request Format, Request Format
		map[string]interface{}{ // Response Example
			"200": VerifyRes{
				Status: VerifyResStatus{
					Code:    200,
					Message: "Success",
				},
				Data: &VerifyResData{
					IsMatch: false,
				},
			},
			"400": VerifyRes{
				Status: VerifyResStatus{
					Code:    400,
					Message: "Something went wrong",
				}, Data: nil,
			},
			"401": VerifyRes{
				Status: VerifyResStatus{
					Code:    401,
					Message: "Something went wrong",
				}, Data: nil,
			},
		},
		group,
	)
}

func VerifyFlow(c *fiber.Ctx) error {
	var req VerifyReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}

	password := req.Data.Password
	hashedpassword := req.Data.HashedPassword
	ismatch := security.BcryptVerify(password, hashedpassword)

	return c.JSON(VerifyRes{
		Status: VerifyResStatus{
			Code:    200,
			Message: "Success",
		},
		Data: &VerifyResData{
			IsMatch: ismatch,
		},
	})
}
