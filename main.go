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
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

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
