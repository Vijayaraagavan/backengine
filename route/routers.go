package route

import (
	// "github.com/valyala/fasthttp"
	controllers "backengine/controllers"
	fiber "github.com/gofiber/fiber/v2"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
	// csrf "github.com/gofiber/fiber/v2/middleware/csrf"
	// utils "github.com/gofiber/fiber/v2/utils"
	// "time"
)

func Routes() *(fiber.App) {
	var app = fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",	//you can use Origin, Content-Type, Accept also
	}))
	// app.Use(csrf.New(csrf.Config{
	// 	KeyLookup:      "header:X-Csrf-Token",
	// 	CookieName:     "csrf_",
	// 	CookieSameSite: "Strict",
	// 	Expiration:     48 * time.Hour,
	// 	KeyGenerator:   utils.UUID,
	// }))
	// Extractor:      func(c *fiber.Ctx) (string, error) { ... },
	var salePipe = app.Group("/new", NakamaChecker)

	salePipe.Get("/", controllers.DefaultRouting)
	salePipe.Post("/register", controllers.NewFriend)
	salePipe.Get("/user/:name", controllers.GetUser)
	salePipe.Post("/user/login", controllers.Login)

	var loginHandler = app.Group("/", LoginAuth)
	loginHandler.Get("/", controllers.Home)
	// return func(ctx *fiber.Ctx) {
	// 	switch(string(ctx.Path())) {
	// 	case "/dashboard":
	// 		controllers.AddMember(ctx)
	// 	default: 
	// 		ctx.Error("not found for this shit", fasthttp.StatusNotFound)
	// 	}
	// }
	return app
}