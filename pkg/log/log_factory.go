package log

import (
	"errors"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
)

const (
	InstanceLogrusLogger int = iota
)

var (
	errInvalidLoggerInstance = errors.New("invalid log instance")
)

// NewLoggerFactory returns a logger interface type based on the provided logger instance
func NewLoggerFactory(instance int) (ports.Logger, error) {
	switch instance {
	case InstanceLogrusLogger:
		return NewLogrusLogger(), nil
	default:
		return nil, errInvalidLoggerInstance
	}
}
