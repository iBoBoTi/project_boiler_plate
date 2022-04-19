package server

import (
	"github.com/iBoBoTi/project_boiler_plate/pkg/database"
	"strconv"
	"time"
)

type serverConfig struct {
	appName       string
	db            *database.Db
	ctxTimeout    time.Duration
	webServerPort ro.Port
	webServer     router.Server
}

func NewServerConfig() *serverConfig {
	return &config{}
}

func (c *serverConfig) ContextTimeout(t time.Duration) *serverConfig {
	c.ctxTimeout = t
	return c
}

func (c *serverConfig) Name(name string) *serverConfig {
	c.appName = name
	return c
}

func (c *serverConfig) Logger(instance int) *serverConfig {
	log, err := log.NewLoggerFactory(instance)
	if err != nil {
		log.Fatalln(err)
	}

	c.logger = log
	c.logger.Infof("Successfully configured log")
	return c
}

func (c *serverConfig) DbSQL(instance int) *serverConfig {
	db, err := database.NewDatabaseFactory(instance)
	if err != nil {
		c.logger.Fatalln(err, "Could not make a connection to the database")
	}

	c.logger.Infof("Successfully connected to the SQL database")

	c.db = db
	return c
}

func (c *serverConfig) Validator(instance int) *serverConfig {
	v, err := validation.NewValidatorFactory(instance)
	if err != nil {
		c.logger.Fatalln(err)
	}

	c.logger.Infof("Successfully configured validator")

	c.validator = v
	return c
}

func (c *serverConfig) WebServer(instance int) *serverConfig {
	s, err := router.NewWebServerFactory(
		instance,
		c.logger,
		c.dbSQL,
		c.dbNoSQL,
		c.validator,
		c.webServerPort,
		c.ctxTimeout,
	)

	if err != nil {
		c.logger.Fatalln(err)
	}

	c.logger.Infof("Successfully configured router server")

	c.webServer = s
	return c
}

func (c *serverConfig) WebServerPort(port string) *serverConfig {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		c.logger.Fatalln(err)
	}

	c.webServerPort = router.Port(p)
	return c
}

func (c *serverConfig) Start() {
	c.webServer.Listen()
}
