package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ssebs/ServerClip/api"
)

// Change this when debugging...
const DEBUG = false

func main() {
	fmt.Println("ServerClip")

	if !DEBUG {
		gin.SetMode(gin.ReleaseMode)
	}
	// Get port from CLI
	port := flag.Int("port", 5000, "Port #")
	flag.Parse()

	a := api.NewAPI(*port)
	log.Fatal(a.Run())
}
