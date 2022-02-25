package actuator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {

	expectedStatus := true
	health := Health{expectedStatus}

	assert.Equal(t, expectedStatus, health.Up)
}
