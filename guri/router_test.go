package guri

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	// fmt.Println(ok, parsePattern("/p/:name"))
	// fmt.Printf("%v\n", []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	// fmt.Println(ok, parsePattern("/p/*"))
	// fmt.Printf("%v\n", []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	// fmt.Println(ok, parsePattern("/p/*name/*"))
	// fmt.Printf("%v\n", []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/testName")
	if n == nil {
		t.Fatal("returned nil value")
	}
	if n.pattern != "/hello/:name" {
		t.Fatal("pattern does not match /hello/:name -- ", n.pattern)
	}
	if ps["name"] != "testName" {
		t.Fatal("name does not match testName --", ps["name"])
	}
	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])
}