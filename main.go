package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tianmarillio/instant-note-backend/config"
	"github.com/tianmarillio/instant-note-backend/src/controllers"
	"github.com/tianmarillio/instant-note-backend/src/middlewares"
)

func main() {
	config.LoadEnv()
	config.ConnectToDB()

	router := gin.Default()
	// router.Use(cors.Default())

	corsConfig := cors.New(cors.Config{
		// AllowOrigins:     []string{"*"},
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	})

	router.Use(corsConfig)

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/auth-test", middlewares.Authenticate, controllers.AuthTest)

	iNoteRouter := router.Group("/i-note", middlewares.Authenticate)
	// iNoteRouter.Use(corsConfig)
	iNoteRouter.POST("", controllers.INoteCreate)
	iNoteRouter.GET("", controllers.INoteFindAll)
	iNoteRouter.GET("/:id", controllers.INoteFindById)
	iNoteRouter.PATCH("/:id", controllers.INoteUpdate)
	iNoteRouter.DELETE("/:id", controllers.INoteDelete)

	router.Run()
}
