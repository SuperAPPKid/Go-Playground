package main

import (
	"fmt"
	"log"
	"net/http"
	"restful/config"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.App.RunMode)
	router := gin.Default()

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", config.App.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(config.App.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.App.WriteTimeout) * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalln("listen:", err)
	}
}
