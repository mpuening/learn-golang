package actuator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo(t *testing.T) {

	expectedVersion := "1.2.3"
	expectedBuild := "abcd123"
	expectedBuildTime := "Thu Feb 24 07:50:10 EST 2022"
	info := Info{expectedVersion, expectedBuild, expectedBuildTime}

	assert.Equal(t, expectedVersion, info.Version)
	assert.Equal(t, expectedBuild, info.Build)
	assert.Equal(t, expectedBuildTime, info.BuildTime)
}
