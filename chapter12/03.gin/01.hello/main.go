package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("您好"))
	})
	r.Run(":8080")
}
