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
	r.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	/* Homepage */
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{"title": "bradenNicholson", "desc": "Software Engineer In California"})
	})

	/* Projects */
	r.GET("/*", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{"title": "bradenNicholson", "desc": "Software Engineer In California"})
	})

	return r
}
