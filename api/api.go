package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type API struct {
	router *gin.Engine
	port   int
}

func NewAPI(port int) *API {
	r := gin.Default()

	r.GET("/", rootHandler)

	return &API{router: r, port: port}
}

func (a *API) Run() error {
	return a.router.Run(fmt.Sprintf("0.0.0.0:%d", a.port))
}

func rootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}
