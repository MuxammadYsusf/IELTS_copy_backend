package main

import (
	"log"

	"github/http/copy/task1/config"
	"github/http/copy/task1/controller"
	"github/http/copy/task1/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Настройка базы данных и конфигурации
	cfg := config.Cfg()
	conn, err := db.Postgres(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	cont := controller.Controller(conn) // Контроллер

	r := gin.Default() // Новый маршрутизатор

	r.POST("/reg", cont.Reg)
	r.DELETE("/reg", cont.DeleteLogInfo)
	r.POST("/login", cont.LoginHandler)

	// Обработчики с middleware

	authorized := r.Group("/", cont.Middleware)
	{
		//smedia
		authorized.POST("/media", cont.CreateSmedia)
		authorized.GET("/media", cont.GetAllSmedia)
		authorized.GET("/media/:id", cont.GetByIdSmedia)
		authorized.PUT("/media/:id", cont.UpdateSmedia)
		authorized.DELETE("/media/:id", cont.DeleteSmedia)
		//likes
		authorized.POST("/like", cont.LikeHandler)
		authorized.GET("/count-likes", cont.CountLikesHandler)
		//update user pass wich was hashed by bicrypt in login file
		authorized.PUT("/up", cont.UpdateUserPass)
	}
	// Запуск сервера
	log.Println("Listening...", cfg.HttpPort)
	err = r.Run(cfg.HttpPort)
	if err != nil {
		log.Fatal(err)
		return
	}
}
