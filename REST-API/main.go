package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"restful/config"
	userController "restful/controllers/user"
	"restful/models"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	models.AutoMigrate()

	gin.SetMode(config.App.RunMode)
	router := gin.Default()

	userRouter := router.Group("/user")
	{
		userRouter.GET("/", userController.GetAll)
		userRouter.POST("/", userController.Create)

		userRouter.GET("/:id", func(c *gin.Context) {
			c.Redirect(http.StatusPermanentRedirect, c.Request.URL.Path+"/profile")
		})

		userRouter.POST("/:id/token", userController.CreateTokenByID)
		userRouter.GET("/:id/profile", userController.GetProfileByID)
		userRouter.DELETE("/me", userController.Auth, userController.Delete)
		userRouter.GET("/me/profile", userController.Auth, userController.GetSelfProfile)
		userRouter.PUT("/me/profile", userController.Auth, userController.UpdateSelfProfile)
		userRouter.PATCH("/me/profile", userController.Auth, userController.PatchSelfProfile)
	}

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", config.App.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(config.App.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.App.WriteTimeout) * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("listen:", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
