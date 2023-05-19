package main

import (
	// "fmt"
	"gee"
	"log"
	"net/http"
	"time"
)

// type Engine struct{
// }

// func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){
// 	switch req.URL.Path{
// 	case "/":
// 		fmt.Fprintf(w,"URL.Path = %q\n",req.URL.Path);
// 	case "/hello":
// 		for k,v:= range req.Header{
// 			fmt.Fprintf(w,"Header[%q] = %q\n",k,v)
// 		}
// 	default:
// 		fmt.Fprintf(w,"404 NOT FOUND: %s",req.URL)
// 	}

// }

func onlyForV2() gee.HandlerFunc{
	return func(c *gee.Context){
		t := time.Now()
		c.Fail(500, "Internal Server Error")

		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.Use(gee.Logger()) // global midlleware
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}

// func indexHandler(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "URL.Path = %q \n", req.URL.Path)
// }

// func helloHandler(w http.ResponseWriter, req *http.Request) {
// 	for k, v := range req.Header {
// 		fmt.Fprintf(w, "Header[%q] = %q \n", k, v)
// 	}
// }
