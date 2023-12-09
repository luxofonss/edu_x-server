package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	ginrecover "server/frameworks-driver/webserver/middlewares/gin"
	"server/frameworks-driver/webserver/routes/gin"
	"server/libs/appctx"
	uploadprovider "server/libs/upload_provider"
)

type GinServerConf struct {
	sqlDb          *gorm.DB
	mongoDb        *mongo.Client
	uploadProvider uploadprovider.Provider
}

func (g *GinServerConf) CreateServer() interface{} {
	ginRouter := gin.Default()

	ginRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	appContext := appctx.NewAppContext(g.sqlDb, g.mongoDb, g.uploadProvider)
	ginRouter.Use(ginrecover.Recover(appContext))

	v1 := ginRouter.Group("/v1")
	ginroutes.SetupRoutes(appContext, v1)

	return ginRouter
}
