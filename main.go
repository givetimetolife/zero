// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"
// 	"zero/script"
// )

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World! %s", time.Now())
// }

// func zhihu(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.TODO()
// 	data, err := script.ZeroDataMap["zhihu"].Data(ctx)
// 	if err != nil {
// 		panic(err)
// 	}
// 	bd, err := json.Marshal(data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Fprint(w, string(bd))
// }
// func main() {
// 	http.HandleFunc("/", greet)
// 	http.HandleFunc("/zhihu", zhihu)
// 	http.ListenAndServe(":8080", nil)
// }

// test gin
package main

import (
	"fmt"
	"log"
	"net/http"
	gin "zero/zero-gin"
) 

func main() {
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(rw, "url.path = %q", r.URL.Path)
	// })

	// http.HandleFunc("/head", func(rw http.ResponseWriter, r *http.Request) {
	// 	for k, v := range r.Header {
	// 		fmt.Fprintf(rw, "head[%q] = %q \n", k, v)
	// 	}
	// })

	engine := gin.New()

	engine.GET("/", func(context *gin.Context) {
		fmt.Fprintf(context.Writer, "url.path = %q", context.Request.URL.Path)
	})

	engine.GET("/head", func(context *gin.Context) {
		for k, v := range context.Request.Header {
			fmt.Fprintf(context.Writer, "head[%q] = %q \n", k, v)
		}
	})

	engine.GET("/login", func(c *gin.Context) {
		var (
			name = c.Query("name")
			pwd  = c.Query("pwd")
		)
		c.String(http.StatusOK, fmt.Sprintf("%s:%s", name, pwd))
	})

	engine.POST("/print", func(c *gin.Context) {

		data := c.PostForm("data")
		c.DATA(http.StatusOK, []byte(data))
	})

	log.Fatal(engine.Run())
}
