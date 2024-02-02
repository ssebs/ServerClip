package main

import (
	"fmt"
	"log"

	"github.com/ssebs/ServerClip/api"
)

func main() {
	fmt.Println("ServerClip")

	a := api.NewAPI(5000)
	log.Fatal(a.Run())
}
