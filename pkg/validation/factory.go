package validation

import (
	"errors"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
)

var (
	errInvalidValidatorInstance = errors.New("invalid validator instance")
)

const (
	InstanceGoPlayground int = iota
)

// NewValidatorFactory takes in an instance of any validator and returns validator interface
func NewValidatorFactory(instance int) (ports.Validator, error) {
	switch instance {
	case InstanceGoPlayground:
		return NewGoPlayground()
	default:
		return nil, errInvalidValidatorInstance
	}
}
