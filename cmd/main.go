package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/ssebs/ServerClip/api"
)

// Change this when debugging...
const DEBUG = false

func main() {
	fmt.Print("ServerClip")

	if !DEBUG {
		gin.SetMode(gin.ReleaseMode)
	}
	// Get port from CLI
	port := flag.Int("p", 5000, "Port #")
	flag.Parse()

	// Create new API / router
	a := api.NewAPI(*port)

	// Get hostname / IP of laptop
	ip, err := getHostIP()
	if err != nil {
		log.Fatal(err)
	}
	// ... and print it
	fmt.Printf(" - %s:%d\n", ip, *port)

	// Print out helper msg for CURLing to here
	printCURLMessageWithHostInfo(ip, *port)

	// Run
	log.Fatal(a.Run())
}

// printCURLMessageWithHostInfo... does what it says
func printCURLMessageWithHostInfo(ip string, port int) {

	// Build CURL cmd
	curlCmd := fmt.Sprintf("curl -XPOST %s:%d/upload -d @/path/to/your_file", ip, port)
	msg := "Upload a file via this command:\n$ " + curlCmd + "\n\nCTRL+C to exit."

	// Print
	fmt.Println(msg)
}

// getHostIP will get the IP of the laptop
func getHostIP() (string, error) {
	// Get IP address
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipnet.IP.IsLoopback() {
			continue
		}
		if ipnet.IP.To4() != nil {
			return ipnet.IP.String(), nil
		}
	}
	return "", nil
}
