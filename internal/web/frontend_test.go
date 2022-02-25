package web

import (
	"testing"

	webtest "github.com/Valiben/gin_unit_test"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFrontend(t *testing.T) {

	web := WebBuilder{
		Mode:            gin.TestMode,
		Port:            8080,
		FrontEndDistDir: ".",
	}.Build()

	webtest.SetRouter(web.router)

	t.Run("index.html", func(t *testing.T) {
		// We are fetching frontend.go
		body, err := webtest.TestOrdinaryHandler("GET", "/frontend.go", "FORM", nil)
		assert.NoError(t, err)

		// frontend.go file contains gin import
		file := string(body)
		assert.Contains(t, file, "github.com/gin-gonic/gin")
	})
}
