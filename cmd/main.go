package main

import (
	"flag"

	actuatorpresentation "learn-golang/internal/actuator/presentation"
	actuatorusecase "learn-golang/internal/actuator/usecase"

	logger "learn-golang/internal/logger"

	web "learn-golang/internal/web"
)

// See Makefile for values
var (
	version   string
	build     string
	buildtime string
)

var LOG *logger.Logger = logger.NewLogger("main")

func main() {
	var port int
	var frontEndDistDir string
	initFlags(&port, &frontEndDistDir)

	LOG.Info("Server starting...")

	runWebServer(port, frontEndDistDir)
}

func initFlags(port *int, frontEndDistDir *string) {
	flag.IntVar(port, "port", 8080, "The port to listen on")
	flag.StringVar(frontEndDistDir, "dist", "./frontend/dist", "The front end dist directory")
	flag.Parse()
}

func runWebServer(port int, frontEndDistDir string) {
	web := web.WebBuilder{
		Port:            port,
		FrontEndDistDir: frontEndDistDir,
	}.Build()

	actuatorpresentation.ActuatorWebBuilder{
		API: web.API(),
		UseCase: &actuatorusecase.Actuator{
			Version:   version,
			Build:     build,
			BuildTime: buildtime,
		},
	}.Build()

	web.Run()
}
