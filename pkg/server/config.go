package server

import (
	"errors"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
	"github.com/iBoBoTi/project_boiler_plate/pkg/router"
)

type server struct{}

var (
	errInvalidWebServerInstance = errors.New("invalid router server instance")
)

const (
	InstanceGin int = iota
)

// NewServerFactory sets up the router based of the instance router provided
func NewServerFactory(instance int, log ports.Logger, router *router.Router) (ports.Router, error) {
	switch instance {
	case InstanceGin:
		return newGinServer(log, router), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}
