package ps

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchPayments(c *gin.Context) {
	body, err := FetchPaymentsFromSource()
	var message string
	if err != nil {
		message = "failed"
	}
	message = string(body)

	c.JSON(http.StatusOK, gin.H{"message": message})
}

func PaymentsRegister(r *gin.RouterGroup) {
	r.GET("/fetch", FetchPayments)
}
