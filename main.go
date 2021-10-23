package main

import (
	"go-template/pkg/events"
	"go-template/pkg/routers"
	setting "go-template/pkg/settings"

	"github.com/fasthttp/router"
	"github.com/savsgio/go-logger"
	"github.com/valyala/fasthttp"
)

func init() {
	logger.SetLevel("debug")
}

func main() {
	setting.Setup()

	events.StartApp()
	defer events.StopApp()

	router := router.New()
	router.GET("/health", routers.Health)
	server := &fasthttp.Server{
		Name:    "Go-Template",
		Handler: router.Handler,
	}
	logger.Debug("Listening in http://localhost:8000...")
	logger.Fatal(server.ListenAndServe(":8000"))

}
