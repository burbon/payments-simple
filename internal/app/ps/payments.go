package ps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	// log "github.com/sirupsen/logrus"
)

func FetchPayments(c *gin.Context) {
	pcnt, err := FetchPaymentsFromSource()
	var message string
	if err != nil {
		message = "failed"
	}
	message = fmt.Sprintf("fetched %d", pcnt)

	c.JSON(http.StatusOK, gin.H{"message": message})
}

func PaymentsRegister(r *gin.RouterGroup) {
	r.GET("/fetch", FetchPayments)
}
