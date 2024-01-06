package server

import (
	"net/http"
	assignmentrecognizeprovider "server/libs/assignment_recognize_provider"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	echoroutes "server/frameworks-driver/webserver/routes/echo"
	"server/libs/appctx"
	uploadprovider "server/libs/upload_provider"
)

type EchoServerConf struct {
	sqlDb                       *gorm.DB
	mongoDb                     *mongo.Client
	uploadProvider              uploadprovider.Provider
	assignmentRecognizeProvider assignmentrecognizeprovider.Provider
}

func (e *EchoServerConf) CreateServer() interface{} {
	echoRouter := echo.New()

	// Add your Echo-specific configurations and routes here
	echoRouter.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	appContext := appctx.NewAppContext(e.sqlDb, e.mongoDb, e.uploadProvider, e.assignmentRecognizeProvider)

	echoroutes.SetupRoutes(appContext, echoRouter)

	return echoRouter
}
