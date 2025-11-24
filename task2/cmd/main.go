package main

import (
	"fmt"
	"log"
	"net"

	"github/http/copy/task2/config"
	"github/http/copy/task2/controller"
	client "github/http/copy/task2/grpc/client"
	"github/http/copy/task2/pkg/db"
	"github/http/copy/task2/server"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Cfg()
	conn, err := db.Postgres(&cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
		return
	}

	clients, err := client.NewGRPCClient()
	if err != nil {
		log.Fatalf("Ошибка создания gRPC-клиента: %v", err)
		return
	}

	cont := controller.Controller(clients)

	grpcServer := server.SetUpServer(conn, clients)

	go func() {
		lis, err := net.Listen("tcp", ":9090")
		if err != nil {
			log.Fatalf("Failed to listen on port: %v!", err)
		}

		fmt.Println("running...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve server on port 9090: %v!", err)
		}
	}()

	r := gin.Default()

	r.POST("/questions/reg", cont.Register)
	r.POST("/questions/login", cont.LogIn)

	authorized := r.Group("/", cont.AddDescHandler)
	{
		authorized.POST("/questions", cont.CreateQuestion)
		authorized.GET("/questions/get/:subject_id/:grade_id", cont.GetQuestion)
		authorized.PUT("/questions/up", cont.UpdateQuestion)
		authorized.DELETE("/questions/del/:id", cont.DeleteQuestion)
		authorized.POST("/questions/test/", cont.AnswerToQuestions)
		authorized.PUT("/questions/uppass", cont.UpdatePassword)
		authorized.GET("/questions/gethistory/:attempt_id", cont.GetResultByAttempt)
		authorized.GET("/questions/getallattempts", cont.GetAttemptsList)
	}

	r.Run(cfg.HttpPort)

}
