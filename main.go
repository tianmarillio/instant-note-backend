package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tianmarillio/instant-note-backend/config"
	"github.com/tianmarillio/instant-note-backend/src/controllers"
	"github.com/tianmarillio/instant-note-backend/src/middlewares"
)

func main() {
	config.LoadEnv()
	config.ConnectToDB()

	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/auth-test", middlewares.Authenticate, controllers.AuthTest)

	iNoteRouter := r.Group("/i-note", middlewares.Authenticate)
	iNoteRouter.POST("/", controllers.INoteCreate)
	iNoteRouter.GET("/", controllers.INoteFindAll)
	iNoteRouter.GET("/:id", controllers.INoteFindById)
	iNoteRouter.PATCH("/:id", controllers.INoteUpdate)
	iNoteRouter.DELETE("/:id", controllers.INoteDelete)

	r.Run()
}
