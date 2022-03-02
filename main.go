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
	r := guri.Default()

	r.GET("/", func(c *guri.Context) {
		c.String(http.StatusOK, "Hello traveler\n")
	})
	r.GET("/panic_test", func(c *guri.Context) {
		names := []string{"guri"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}