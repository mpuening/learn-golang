package web

import (
	"strconv"

	logger "learn-golang/internal/logger"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var LOG *logger.Logger = logger.NewLogger("web")

type WebBuilder struct {
	Mode            string
	Port            int
	FrontEndDistDir string
}

func (this WebBuilder) Build() *Web {
	// TODO: validate Port and Dist Dir

	if this.Mode == "" {
		this.Mode = gin.ReleaseMode
	}
	gin.SetMode(this.Mode)

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile(this.FrontEndDistDir, true)))

	api := router.Group("/api")

	return &Web{router, api, this.Port}
}

type Web struct {
	router *gin.Engine
	api    *gin.RouterGroup
	port   int
}

func (this *Web) Router() *gin.Engine {
	return this.router
}

func (this *Web) API() *gin.RouterGroup {
	return this.api
}

func (this *Web) Run() {
	port := ":" + strconv.Itoa(this.port)
	LOG.Debugf("Router starting on port %s", port)
	this.router.Run(port)
}
