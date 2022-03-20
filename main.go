package main

import (
	"context"
	"todo-backend/config"
	"todo-backend/route"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"time"
)

func main() {
	var conf *viper.Viper
	var err error

	if conf, err = config.LoadConfig("./"); err != nil {
		panic(err)
	}

	e := echo.New()

	route.Route(e)


	serverEnv := conf.Get("ENV_NAME")
	serverPort := conf.GetInt("SERVER_PORT")
	serverHost := conf.GetString("SERVER_HOST")
	serverAddress := fmt.Sprintf("%s:%d", serverHost, serverPort)

	go func() {
		fmt.Printf("Running on %s environment", serverEnv)
		if err := e.Start(serverAddress); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// shutdown gracefully
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// wait server to shutdown with timeout 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
