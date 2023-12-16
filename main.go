package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"server/frameworks-driver/database"
	"server/frameworks-driver/webserver/server"
	uploadprovider "server/libs/upload_provider"
)

func main() {
	// config mongodb
	//mongoConf := database.DBConfig{
	//	IdentificationName: "mongo",
	//	DB:                 "clean_architecture",
	//	User:               "",
	//	Password:           "",
	//	Host:               "localhost",
	//	Port:               "27017",
	//	Type:               "mongo",
	//	SSLMode:            "true",
	//	TimeZone:           "",
	//}
	//db := database.NewDatabase(mongoConf.Type)
	//
	//mongoConnect, err := db.Database.Connect(mongoConf)
	//if err != nil {
	//	// Xử lý lỗi nếu có
	//	fmt.Println("has err", err)
	//	panic(err)
	//}
	//
	//mongoDb, ok := mongoConnect.(*mongo.Client)
	//if !ok {
	//	fmt.Println("err mongo type")
	//}

	// Config postgresql
	configPostgres := database.DBConfig{
		IdentificationName: "postgres",
		DB:                 "project_3",
		User:               "postgres",
		Password:           "admin",
		Host:               "localhost",
		Port:               "5432",
		Type:               "postgres",
		SSLMode:            "disable",
		TimeZone:           "Asia/Ho_Chi_Minh",
	}

	postgresConf := database.NewDatabase(configPostgres.Type)
	postgresDBImpl, err := postgresConf.Database.Connect(configPostgres)
	if err != nil {
		fmt.Println("err postgres type")
		panic(err)
	}

	postgresDb, ok := postgresDBImpl.(*gorm.DB)
	if !ok {
		fmt.Println("err postgres type")
	}

	postgresDb.Debug()

	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3APIKey := os.Getenv("S3_API_KEY")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")
	s3Domain := os.Getenv("S3_DOMAIN")

	uploadProvider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	ginServer := server.NewServerFactory("gin", postgresDb, nil, uploadProvider)
	router := ginServer.CreateServer()
	appRouter := router.(*gin.Engine)

	appRouter.Run()
	//
	//echoServer := server.NewServerFactory("echo", postgresDb, mongoDb)
	//
	//router := echoServer.CreateServer()
	//
	//appRouter := router.(*echo.Echo)
	//appRouter.Start(":8080")
}
