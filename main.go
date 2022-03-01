package main

import (
	"guri"
	"net/http"
)

func main() {
	r := guri.New()
	r.GET("/", func(c *guri.Context) {
		c.HTML(http.StatusOK, "<H1>Hello traveler!</h1>")
	})

	r.GET("/hello", func(c *guri.Context) {
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *guri.Context) {
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *guri.Context) {
		c.JSON(http.StatusOK, guri.H{"filepath": c.Param("filepath")})
	})

	r.POST("/login", func(c *guri.Context) {
		c.JSON(http.StatusOK, guri.H {
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}