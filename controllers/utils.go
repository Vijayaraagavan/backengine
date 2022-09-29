package controllers

import (
	fiber "github.com/gofiber/fiber/v2"
	"io"
	"encoding/json"
	"fmt"
	"backengine/errs"
)

func DefaultRouting(ctx *fiber.Ctx) error {
	// fmt.Fprintf(ctx, "Oops, go to grand line")
	return ctx.SendString("Oops, go to grand line")
}


func parseJSON(c *fiber.Ctx, body io.ReadCloser, target interface{}) bool {
	if body != nil {
		defer body.Close()
	}

	err := json.NewDecoder(body).Decode(target)
	if err == nil {
		return true;
	}
	// e := &errs.Error{
	// 	Code: 404,
	// 	Message: "Invalid data",
	// 	Err: err,
	// }
	// log.Printf("Error in get json: %s", e.Stack())
	// renderError(c, e)
	fmt.Println(fiber.StatusBadRequest)
	return false
}

func FailureResp(c *fiber.Ctx, err error, message string) (errs.Error) {
	c.Status(fiber.StatusBadRequest)
	
	er := fmt.Sprintf("%v", err)
	reqErr := errs.Error{Code: 400, Message: message, Err: er}
	return reqErr
}