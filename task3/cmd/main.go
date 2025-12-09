package main

import (
	"log"
	"net"

	"github/http/copy/task3/config"
	controller "github/http/copy/task3/controller"
	"github/http/copy/task3/grpc/client"
	"github/http/copy/task3/pkg/db"
	"github/http/copy/task3/server"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Cfg()

	conn, err := db.Postgres(&cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	clients, err := client.NewGRPCClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	cont := controller.NewController(clients)

	grpcServer := server.NewServer(conn, clients)

	go func() {
		lis, err := net.Listen("tcp", ":9090")
		if err != nil {
			log.Fatal("ERROR! ", err)
			return
		}

		log.Println("running...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve server on port 9090: %v!", err)
		}
	}()

	r := gin.Default()

	r.POST("/test/reg", cont.Reg)
	r.POST("/test/login", cont.Login)

	authorized := r.Group("/", cont.Midleware)

	{
		authorized.POST("/test/create", cont.CreateNewTest)

		authorized.POST("/test/writing/create", cont.CreateWritingQuestions)
		authorized.GET("/test/writing/get-questions/:taskId/:testId", cont.GetWritingQuestions)

		authorized.POST("/test/speaking/create", cont.CreateSpeakingQuestions)
		authorized.GET("/test/speaking/get-questions/:partId/:testId", cont.GetSpeakingQuestions)

		authorized.POST("/test/reading-questions/create", cont.CreateReadingQuestions)
		authorized.POST("/test/reading-content/create", cont.CreateReadinContent)
		authorized.POST("test/reading/get-content", cont.GetReadinContent)
		authorized.POST("test/reading/answer", cont.SaveReadingAnswers)

		authorized.POST("/test/listening-content/create", cont.CreateListeningContent)
		authorized.POST("/test/listening-questions/create", cont.CreateListeningQuestions)
		authorized.POST("/test/listening/get-content", cont.GetListeningContent)
		authorized.POST("/test/listening/answer", cont.SaveListeningAnswers)
	}
	r.Run(cfg.HttpPort)
}
