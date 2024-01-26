package main

import (
	"fmt"
	"github.com/hudayberdipolat/go-url-shortener/internal/app"
	"github.com/hudayberdipolat/go-url-shortener/internal/setup/constructor"
	"log"
)

func main() {
	appDependencies, err := app.GetDependencies()
	if err != nil {
		log.Fatal(err.Error())
	}
	constructor.Build(appDependencies)
	appRouter := app.NewApp(appDependencies)
	runServer := fmt.Sprintf("%s:%s",
		appDependencies.Config.HttpServer.ServerHost, appDependencies.Config.HttpServer.ServerPort)
	if errRunServer := appRouter.Listen(runServer); errRunServer != nil {
		log.Fatal("error run server : --->", errRunServer.Error())
	}
}
