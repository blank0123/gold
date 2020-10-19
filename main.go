package main

import (
	"context"
	"fmt"
	"github.com/kainhuck/gold/models"
	"github.com/kainhuck/gold/pkg/config"
	"github.com/kainhuck/gold/pkg/log"
	"github.com/kainhuck/gold/router"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	config.Init()
	_ = log.Init()
	models.Init()
}

func main() {
	handler := router.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Collection.Server.HTTPPort),
		Handler:        handler,
		ReadTimeout:    config.Collection.Server.ReadTimeout,
		WriteTimeout:   config.Collection.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.SugarLogger.Errorf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.SugarLogger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.SugarLogger.Fatalf("Server Shutdown:", err)
	}

	log.SugarLogger.Info("Server exiting")
}
