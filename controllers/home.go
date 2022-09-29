package controllers

import(
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	// pg "github.com/go-pg/pg/v10"
	// "errors"
	// "github.com/Vijayaraagavan/backengine/models"
	// "github.com/Vijayaraagavan/backengine/util"
	// "github.com/Vijayaraagavan/backengine/db"
	// "github.com/Vijayaraagavan/backengine/auth/core"
	// "log"
)

func Home(c *fiber.Ctx) error {
	fmt.Println(c.Body)
	return c.SendString("welcome")
}