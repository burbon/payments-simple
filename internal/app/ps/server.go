package ps

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v := r.Group("/")

	PaymentsRegister(v)

	v.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}

func RunServer() {
	r := SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
