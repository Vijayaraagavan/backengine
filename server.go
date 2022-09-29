package main

import (
	"fmt"
	"flag"
	"strconv"
	route "github.com/Vijayaraagavan/backengine/route"
	config "github.com/Vijayaraagavan/backengine/config"
	db "github.com/Vijayaraagavan/backengine/db"
	scripts "github.com/Vijayaraagavan/backengine/scripts"
	"log"
	"os"
	// "github.com/valyala/fasthttp"
)

// to run app -> 	go run server.go --config dev 
func main() {
	fmt.Println("success")
	// fasthttp.ListenAndServe(":8086", route.Routes())
	configFlag := flag.String("config", "", "give path of config file")
	configFla := "dev"
	flag.Parse()
	log.Println(configFlag)
	cfg := config.Parse("config/" + configFla + ".json")
	
	logFile, _ := os.OpenFile("./logs/records.txt", os.O_RDWR|os.O_CREATE, 0644)
	log.SetOutput(logFile)
	pgdb := db.Connect()
	scripts.Run(pgdb)
	log.Println("server listens on port: 8086")
	route.Routes().Listen(":" + strconv.Itoa(cfg.Port))
}