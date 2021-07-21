package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	type user struct {
		Name     string `json:"name"` // 这里要注意字段名的大小写，小写的话外部包访问不到
		Password string `json:"pwd"`
	}

	var smurfs = user{
		Name:     "Smurfs的格格巫",
		Password: "123456789",
	}

	c.JSON(http.StatusOK, smurfs)
}

func main() {
	r := gin.Default()

	r.GET("/index", indexHandler)

	r.Run()
}
