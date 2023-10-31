package main

import (
	"fmt"
	"inventory/pkg/config"
	"inventory/pkg/db"
	"inventory/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	c, err := config.LoadConfig()
	fmt.Println("port", c.Port)

	if err != nil {
		log.Fatalln("Failed at config.......", err)
	}
	db.Init(c.DBUrl)
	route := gin.Default()
	routes.Routes(route)
	route.Run(c.Port)

}
