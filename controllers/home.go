package controllers

import(
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	// pg "github.com/go-pg/pg/v10"
	// "errors"
	// "backengine/models"
	// "backengine/util"
	// "backengine/db"
	// "backengine/auth/core"
	// "log"
)

func Home(c *fiber.Ctx) error {
	fmt.Println(c.Body)
	return c.SendString("welcome")
}