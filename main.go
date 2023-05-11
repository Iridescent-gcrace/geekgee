package main

import (
	// "fmt"
	"gee"
	"net/http"
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

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
	// log.Fatal(http.ListenAndServe(":9999", egine))
}

// func indexHandler(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "URL.Path = %q \n", req.URL.Path)
// }

// func helloHandler(w http.ResponseWriter, req *http.Request) {
// 	for k, v := range req.Header {
// 		fmt.Fprintf(w, "Header[%q] = %q \n", k, v)
// 	}
// }
