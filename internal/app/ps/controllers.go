package ps

import (
	"fmt"
	//	"net/http"
	// log "github.com/sirupsen/logrus"
	"github.com/go-openapi/runtime/middleware"

	"payments-simple/internal/pkg/rest/models"
	"payments-simple/internal/pkg/rest/restapi/operations"
)

func SetupAPIHandlers(api *operations.PaymentsSimpleApplicationAPI) {
	api.GetPingHandler = operations.GetPingHandlerFunc(
		func(params operations.GetPingParams) middleware.Responder {
			message := "pong"
			return operations.NewGetPingOK().WithPayload(&models.Message{Message: &message})
		})

	api.GetFetchHandler = operations.GetFetchHandlerFunc(
		func(params operations.GetFetchParams) middleware.Responder {
			pcnt, err := FetchPaymentsFromSource()
			var message string
			if err != nil {
				message = "failed"
			}
			message = fmt.Sprintf("fetched %d", pcnt)
			return operations.NewGetFetchOK().WithPayload(&models.Message{Message: &message})
		})

	api.GetHandler = operations.GetHandlerFunc(
		func(params operations.GetParams) middleware.Responder {
			payments := QGetPayments()
			return operations.NewGetOK().WithPayload(payments)
		})
}
