package main

import (
	"fmt"
	"flag"
	"strconv"
	route "backengine/route"
	config "backengine/config"
	db "backengine/db"
	scripts "backengine/scripts"
	"log"
	"os"
	// "github.com/valyala/fasthttp"
)

func main() {
	fmt.Println("success")
	// fasthttp.ListenAndServe(":8086", route.Routes())
	configFlag := flag.String("config", "", "give path of config file")
	flag.Parse()
	cfg := config.Parse("config/" + *configFlag + ".json")
	
	logFile, _ := os.OpenFile("./logs/records.txt", os.O_RDWR|os.O_CREATE, 0644)
	log.SetOutput(logFile)
	pgdb := db.Connect()
	scripts.Run(pgdb)
	log.Println("server listens on port: 8086")
	route.Routes().Listen(":" + strconv.Itoa(cfg.Port))
}