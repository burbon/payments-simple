package ps

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-openapi/loads"
	//"github.com/go-openapi/swag"
	"github.com/go-openapi/runtime/middleware"

	"payments-simple/internal/pkg/rest/models"
	"payments-simple/internal/pkg/rest/restapi"
	"payments-simple/internal/pkg/rest/restapi/operations"
)

func RunServer() {
	//RunGinServer()
	RunSwaggerServer()
}

func SetupHandler() http.Handler {
	//h := SetupGinRouter()
	h := SetupSwaggerHandler()
	return h
}

func SetupGinRouter() *gin.Engine {
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

func RunGinServer() {
	r := SetupGinRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}

func SetupSwaggerAPI() *operations.PaymentsSimpleApplicationAPI {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewPaymentsSimpleApplicationAPI(swaggerSpec)

	SetupHandlers(api)

	return api
}

func SetupHandlers(api *operations.PaymentsSimpleApplicationAPI) {
	api.GetPingHandler = operations.GetPingHandlerFunc(
		func(params operations.GetPingParams) middleware.Responder {
			message := "pong"
			return operations.NewGetFetchOK().WithPayload(&models.Message{Message: &message})
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
}

func SetupSwaggerHandler() http.Handler {
	api := SetupSwaggerAPI()
	return api.Serve(nil)
}

func RunSwaggerServer() {
	api := SetupSwaggerAPI()
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// set the port this service will be run on
	server.Port = int(Config.Port)

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
