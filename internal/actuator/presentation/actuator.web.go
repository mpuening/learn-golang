package actuator

import (
	entity "learn-golang/internal/actuator/entity"
	usecase "learn-golang/internal/actuator/usecase"

	"github.com/gin-gonic/gin"
)

//
// Actuator Web Presentation
//
type ActuatorWebBuilder struct {
	API     *gin.RouterGroup
	UseCase *usecase.Actuator
}

func (this ActuatorWebBuilder) Build() *ActuatorWeb {
	// TODO: Validate builder fields
	presentation := &ActuatorWeb{
		usecase: this.UseCase,
	}
	this.API.GET("/info", func(c *gin.Context) {
		info, err := presentation.info()
		if err == nil {
			c.JSON(200, gin.H{
				"version":   info.Version,
				"build":     info.Build,
				"buildTime": info.BuildTime,
			})
		} else {
			renderError(c, err)
		}
	})
	this.API.GET("/health", func(c *gin.Context) {
		health, err := presentation.health()
		if err == nil {
			c.JSON(200, gin.H{
				"UP": health.Up,
			})
		} else {
			renderError(c, err)
		}
	})
	return presentation
}

func renderError(c *gin.Context, err error) {
	c.JSON(500, gin.H{
		"error": err.Error(),
	})
}

type ActuatorWeb struct {
	usecase *usecase.Actuator
}

func (this *ActuatorWeb) info() (entity.Info, error) {
	return this.usecase.GetInfo()
}

func (this *ActuatorWeb) health() (entity.Health, error) {
	return this.usecase.GetHealth()

}
