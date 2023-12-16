package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	ginrecover "server/frameworks-driver/webserver/middlewares/gin"
	"server/frameworks-driver/webserver/routes/gin"
	"server/libs/appctx"
	uploadprovider "server/libs/upload_provider"
	"time"
)

type GinServerConf struct {
	sqlDb          *gorm.DB
	mongoDb        *mongo.Client
	uploadProvider uploadprovider.Provider
}

func (g *GinServerConf) CreateServer() interface{} {
	ginRouter := gin.Default()

	// CORS
	ginRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	appContext := appctx.NewAppContext(g.sqlDb, g.mongoDb, g.uploadProvider)
	ginRouter.Use(ginrecover.Recover(appContext))

	v1 := ginRouter.Group("/v1")
	ginroutes.SetupRoutes(appContext, v1)

	return ginRouter
}
