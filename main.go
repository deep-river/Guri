package main

import (
	"guri"
	"log"
	"net/http"
	"time"
)

func v2APILogger() guri.HandlerFunc {
	return func(c *guri.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := guri.New()
	r.Use(guri.Logger())
	r.GET("/", func(c *guri.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v2 := r.Group("/v2") 
	v2.Use(v2APILogger())
	{
		v2.GET("/hello/:name", func(c *guri.Context) {
			c.String(http.StatusOK, "hello %s, you are at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}