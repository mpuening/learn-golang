package actuator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo(t *testing.T) {

	expectedVersion := "1.2.3"
	expectedBuild := "abcd123"
	expectedBuildTime := "Thu Feb 24 07:50:10 EST 2022"

	actuator := Actuator{
		Version:   expectedVersion,
		Build:     expectedBuild,
		BuildTime: expectedBuildTime,
	}

	info, err := actuator.GetInfo()

	assert.NoError(t, err)
	assert.Equal(t, expectedVersion, info.Version)
	assert.Equal(t, expectedBuild, info.Build)
	assert.Equal(t, expectedBuildTime, info.BuildTime)
}

func TestHealth(t *testing.T) {

	actuator := Actuator{}

	health, err := actuator.GetHealth()

	assert.NoError(t, err)
	assert.Equal(t, true, health.Up)
}
