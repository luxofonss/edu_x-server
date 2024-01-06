package server

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	assignmentrecognizeprovider "server/libs/assignment_recognize_provider"
	uploadprovider "server/libs/upload_provider"
)

type Server struct{}

type ServerFactory interface {
	CreateServer() interface{}
}

func NewServerFactory(framework string, sqlDb *gorm.DB, mongoDb *mongo.Client, uploadProvider uploadprovider.Provider, assignmentRecognizeProvider assignmentrecognizeprovider.Provider) ServerFactory {
	if framework == "gin" {
		return &GinServerConf{sqlDb: sqlDb, mongoDb: mongoDb, uploadProvider: uploadProvider, assignmentRecognizeProvider: assignmentRecognizeProvider}
	} else if framework == "echo" {
		return &EchoServerConf{sqlDb: sqlDb, mongoDb: mongoDb, uploadProvider: uploadProvider, assignmentRecognizeProvider: assignmentRecognizeProvider}
	}
	return nil
}
