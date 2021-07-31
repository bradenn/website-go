package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	/* Heartbeat */
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong!")
	})

	/* Homepage */
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{"title": "bradenNicholson", "desc": "Software Engineer In California"})
	})

	return r
}
