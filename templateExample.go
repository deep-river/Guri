package main

import (
	"fmt"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

// func main() {
// 	r := guri.New()
// 	r.Use(guri.Logger())
// 	r.SetFuncMap(template.FuncMap{
// 		"FormatAsDate": FormatAsDate,
// 	})
// 	r.LoadHTMLGlob("templates/*")
// 	r.Static("/assets", "./static")

// 	stu1 := &student{Name: "stu1", Age: 20}
// 	stu2 := &student{Name: "stu2", Age: 22}
// 	r.GET("/", func(c *guri.Context) {
// 		c.HTML(http.StatusOK, "css.tmpl", nil)
// 	})
// 	r.GET("/students", func(c *guri.Context) {
// 		c.HTML(http.StatusOK, "arr.tmpl", guri.H{
// 			"title": "guri",
// 			"stuArr": [2]*student{stu1, stu2},
// 		})
// 	})
// 	r.GET("/date", func(c *guri.Context) {
// 		c.HTML(http.StatusOK, "custom_func.tmpl", guri.H {
// 			"title": "guri",
// 			"now": time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC),
// 		})
// 	})
	
// 	r.Run(":9999")
// }