package ps

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/go-openapi/loads"
	//"github.com/go-openapi/swag"

	"payments-simple/internal/pkg/rest/restapi"
	"payments-simple/internal/pkg/rest/restapi/operations"
)

func SetupAPI() *operations.PaymentsSimpleApplicationAPI {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewPaymentsSimpleApplicationAPI(swaggerSpec)

	SetupAPIHandlers(api)

	return api
}

func SetupHandler() http.Handler {
	api := SetupAPI()
	return api.Serve(nil)
}

func RunServer() {
	api := SetupAPI()
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// set the port this service will be run on
	server.Port = int(Config.Port)

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
