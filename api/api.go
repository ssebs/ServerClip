package api

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"golang.design/x/clipboard"
)

type API struct {
	router *gin.Engine
	port   int
}

func NewAPI(port int) *API {

	// Use below if we want the log
	// r := gin.Default()

	// Use below to hide the log
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/", rootHandler)
	r.POST("/", uploadHandler)

	return &API{router: r, port: port}
}

func (a *API) Run() error {
	return a.router.Run(fmt.Sprintf("0.0.0.0:%d", a.port))
}

// uploadHandler will read the request body and put it on the clipboard.
func uploadHandler(c *gin.Context) {
	// read blob into b
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Error(err)
		c.String(400, "failed to read request body, err: %s", err.Error())
	}
	defer c.Request.Body.Close()

	// Copy to clipboard
	if err := clipboard.Init(); err != nil {
		c.Error(err)
		c.String(400, "failed to init clipboard, err: %s", err.Error())
	}
	// Write to clipboard
	clipboard.Write(clipboard.FmtText, b)
	fmt.Println("Copied file to clipboard!\nWaiting for file...(press CTRL+C to quit)")
	c.String(201, "Uploaded, check your clipboard.\n")
}

func rootHandler(c *gin.Context) {
	c.String(200, "Please POST with the data in the request body.")
}
