package main

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/pprof"

func main() {
	r := gin.Default()
	pprof.Register(r)
	// go tool pprof -svg http://localhost:8080/debug/pprof/goroutine\?debug\=1
	r.GET("/", func(c *gin.Context) {
		// 如果用JSON，就可以返回JSON格式，并带上Content-Type application/json; charset=utf-8
		c.JSON(200, "data")
	})
	r.GET("/histroy/getParm/:name", func(c *gin.Context) {
		data := c.Param("name")
		c.JSON(200, data)
	})
	r.GET("/histroy/getQuery", func(c *gin.Context) {
		data := c.Query("name")
		c.JSON(200, data)
	})
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
