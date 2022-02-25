package actuator

import (
	actuator "learn-golang/internal/actuator/entity"
)

//
// Actuator Presentation
//
type ActuatorPresenter interface {
	info() (actuator.Info, error)
	health() (actuator.Health, error)
}
