package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/laughing-nerd/goformix"
)

// Just a sample struct
// Specify goformix types in the struct itself
type Employee struct {
	FirstName string `json:"firstName" goformix:"string"`
	LastName  string `json:"lastName" goformix:"string"`
	Email     string `json:"email" goformix:"email|gmail.com"`
	Phone     string `json:"phone" goformix:"phone"`
}

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		employee := Employee{}
		c.BodyParser(&employee)

		// Returns a map. Marshal it into JSON for better readability
		checks := goformix.Validate(employee)
		j, _ := json.Marshal(checks)
    fmt.Println(string(j))
    return nil
	})

	log.Fatal(app.Listen(":3000"))
}
