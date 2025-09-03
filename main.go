package main

import (
	fmt "fmt"
	os "os"
	sync "sync"
	json "encoding/json"

 	fiber "github.com/gofiber/fiber/v2"

    driver "app/app/driver"
    apidoc "app/apidoc"
)

var once sync.Once

func main() {
    once.Do(func() {
    	driver.Init()
    })
	port := os.Getenv("PROJECT_PORT")

	app := fiber.New()
	app.Get("/swagger.json", func(c *fiber.Ctx) error {
			spec := apidoc.BuildSwagger()
			b, _ := json.Marshal(spec)
			return c.Type("json").Send(b)
		})
	driver.Main(app)

	app.Static("/swagger", "./web/swagger", fiber.Static{
	    Browse: true,     // allows showing index.html by default
	})
	fmt.Println("🚀 Server running on :"+port)
	app.Listen(":"+port)
}
