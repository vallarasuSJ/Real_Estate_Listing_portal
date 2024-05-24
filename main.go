package main

import (
	"fmt"
	"log"
	"os"
	"real_estate/src/utils/db"

	_ "github.com/lib/pq"
	"real_estate/src/config"
	"real_estate/src/routes"
	
)

const (
	ENVDev = "dev"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = ENVDev
	}
	file, err := os.Open(env + ".json")
	if err != nil {
		log.Println("Unable to get env file.Err:", err)
		os.Exit(1)
	}
	err = config.Parse(config.TypeJSON, file)
	if err != nil {
		log.Println("Unable to parse json env file.Err:", err)
		os.Exit(1)
	} 
	db.Init() 

	r:=routes.GetRouter()
	r.Run(":"+config.Conf.Port)
	fmt.Printf("%+v", config.Conf)
}
