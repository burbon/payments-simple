package ps

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchPayments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "fetched"})
}

func PaymentsRegister(r *gin.RouterGroup) {
	r.GET("/fetch", FetchPayments)
}
