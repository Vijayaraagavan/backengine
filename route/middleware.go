package route

import (
	fiber "github.com/gofiber/fiber/v2"
	"fmt"
	"time"
	
	// authentication "backengine/auth/core"
)

type nakamaChecker struct {
	Gender string `json:"gender"`
}
type jwtValid struct {
	Id int64
}
func (j jwtValid) Valid() error {
	return nil
}

func NakamaChecker(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	// cook := new(fiber.Cookie)
	var cook = &fiber.Cookie{
		Name: "zoro",
		Value: "1.119",
		Expires: time.Now().Add(24 * time.Hour),
		Domain: "http://192.168.0.106:8080/",
	}
	c.Cookie(cook)
	// fmt.Println(c.Cookies("zoro"))
	var a = &nakamaChecker{}
	c.BodyParser(a)
	fmt.Println(a.Gender)
	// fmt.Println(c.BaseURL())
	if(a.Gender == "male") {
		return fiber.NewError(400, "You are not allowed my friend");
		// return nil
	}
	fmt.Println(c.Params("name"))
	c.Set("type", "straw Hat")
	return c.Next()
}

// func LoginAuth(c *fiber.Ctx) error {
// 	headers := c.GetReqHeaders()
// 	token := headers["Authorization"]
// 	fmt.Println(token)
// 	tokenbuf := new(bytes.Buffer)
// 	io.Copy(tokenbuf, []byte(token))
// 	tokenString := strings.TrimSpace(tokenbuf.String())

// 	token, err := jwt.ParseWithClaims(tokenString, &jwtValid{}, authentication.GetAuthMethod)
// 	fmt.Println(token.Valid)
// 	fmt.Printf("%+v", token)
// 	return c.Next()
// }

func SetCookies(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name: "zoro",
		Value: "1.119",
		Expires: time.Now().Add(60000 * time.Second),
	})
	return nil
}