package server

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	uploadprovider "server/libs/upload_provider"
)

type Server struct{}

type ServerFactory interface {
	CreateServer() interface{}
}

func NewServerFactory(framework string, sqlDb *gorm.DB, mongoDb *mongo.Client, uploadProvider uploadprovider.Provider) ServerFactory {
	if framework == "gin" {
		return &GinServerConf{sqlDb: sqlDb, mongoDb: mongoDb, uploadProvider: uploadProvider}
	} else if framework == "echo" {
		return &EchoServerConf{sqlDb: sqlDb, mongoDb: mongoDb, uploadProvider: uploadProvider}
	}
	return nil
}
