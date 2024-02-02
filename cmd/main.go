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
	if !DEBUG {
		gin.SetMode(gin.ReleaseMode)
	}
	// Get port from CLI
	port := flag.Int("p", 5000, "Port #")
	flag.Parse()

	// Create new API / HTTP router
	a := api.NewAPI(*port)

	// Get hostname / IP of laptop
	ip, err := getHostIP()
	if err != nil {
		log.Fatal(err)
	}

	// Print out helper msg for CURLing to here
	printCURLMessageWithHostInfo(ip, *port)

	// Run
	log.Fatal(a.Run())
}

// printCURLMessageWithHostInfo... does what it says
func printCURLMessageWithHostInfo(ip string, port int) {
	// Use ssh -R port:localhost:port for each jump
	// Then curl -XPOST localhost:5000/ --data-binary @/path/to/your_file to
	curlCmd := fmt.Sprintf("$ curl -XPOST 127.0.0.1:%d --data-binary @/path/to/your_file\n", port)

	msg := fmt.Sprintf("ServerClip - Listening on http://%s:%d\t(press CTRL+C to quit)\n", ip, port)
	msg += fmt.Sprintf("SSH to your host using:\n$ ssh -R %d:127.0.0.1:%d <user>@<hostname>\n", port, port)
	msg += "To send a file to your clipboard, run this command:\n" + curlCmd
	msg += "Waiting for file...\n"

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
