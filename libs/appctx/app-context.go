package appctx

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"server/libs/upload_provider"
)

type AppContext interface {
	GetMainSQLDbConnection() *gorm.DB

	GetSecretKey() string

	GetMongoDbConnection() *mongo.Client

	GetUploadProvider() uploadprovider.Provider
}

type appCtx struct {
	sqlDb          *gorm.DB
	mongoDb        *mongo.Client
	secretKey      string
	uploadProvider uploadprovider.Provider
}

func NewAppContext(sqlDb *gorm.DB, mongoDb *mongo.Client, uploadProvider uploadprovider.Provider) *appCtx {
	return &appCtx{sqlDb: sqlDb, mongoDb: mongoDb, uploadProvider: uploadProvider}
}

func (ctx *appCtx) GetSecretKey() string {
	return ctx.secretKey
}
func (ctx *appCtx) GetMainSQLDbConnection() *gorm.DB {
	return ctx.sqlDb
}

func (ctx *appCtx) GetUploadProvider() uploadprovider.Provider {
	return ctx.uploadProvider
}

func (ctx *appCtx) GetMongoDbConnection() *mongo.Client {
	return ctx.mongoDb
}
