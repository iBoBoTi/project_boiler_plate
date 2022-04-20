package main

import (
	logger "github.com/iBoBoTi/project_boiler_plate/pkg/log"
	"github.com/iBoBoTi/project_boiler_plate/pkg/router"
	"github.com/iBoBoTi/project_boiler_plate/pkg/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	l, err := logger.NewLoggerFactory(logger.InstanceLogrusLogger)
	if err != nil {
		log.Fatalln(err)
	}

	env := os.Getenv("GIN_MODE")
	if env != "release" {
		if err := godotenv.Load(); err != nil {
			l.Fatalln("couldn't load env vars: %v", err)
		}
	}

	r := router.NewRouter()
	s, err := server.NewServerFactory(server.InstanceGin, l, r)
	if err != nil {
		l.Fatalln("couldn't create server: %v", err)
	}
	s.Run()
}
