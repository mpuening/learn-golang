package actuator

import (
	entity "learn-golang/internal/actuator/entity"
)

//
// Actuator Use Cases
//
type Actuator struct {
	Version   string
	Build     string
	BuildTime string
}

func (this Actuator) GetInfo() (entity.Info, error) {
	return entity.Info{Version: this.Version, Build: this.Build, BuildTime: this.BuildTime}, nil
}

func (this Actuator) GetHealth() (entity.Health, error) {
	return entity.Health{Up: true}, nil
}
