package actuator

import (
	"testing"

	entity "learn-golang/internal/actuator/entity"
	usecase "learn-golang/internal/actuator/usecase"

	web "learn-golang/internal/web"

	webtest "github.com/Valiben/gin_unit_test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestActuator(t *testing.T) {

	web := web.WebBuilder{
		Mode:            gin.TestMode,
		Port:            8080,
		FrontEndDistDir: ".",
	}.Build()

	ActuatorWebBuilder{
		API: web.API(),
		UseCase: &usecase.Actuator{
			Version:   "0.0.0",
			Build:     "abc1234",
			BuildTime: "today",
		},
	}.Build()

	webtest.SetRouter(web.Router())

	t.Run("/api/info", func(t *testing.T) {

		info := entity.Info{}
		err := webtest.TestHandlerUnMarshalResp("GET", "/api/info", "JSON", nil, &info)
		assert.NoError(t, err)

		assert.Equal(t, "0.0.0", info.Version)
		assert.Equal(t, "abc1234", info.Build)
		assert.Equal(t, "today", info.BuildTime)
	})

	t.Run("/api/health", func(t *testing.T) {

		health := entity.Health{}
		err := webtest.TestHandlerUnMarshalResp("GET", "/api/health", "JSON", nil, &health)
		assert.NoError(t, err)

		assert.Equal(t, true, health.Up)
	})
}
