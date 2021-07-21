package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	username := c.DefaultQuery("username", "Smurfs的格格巫")
	address := c.Query("address")
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"address":  address,
	})
}

func main() {
	r := gin.Default()

	r.GET("/query_string", indexHandler)

	r.Run()
}
