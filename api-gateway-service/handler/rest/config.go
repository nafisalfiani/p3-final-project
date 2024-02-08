package rest

import (
	"time"
)

type Config struct {
	Port            string
	Mode            string
	LogRequest      bool
	LogResponse     bool
	Timeout         time.Duration
	ShutdownTimeout time.Duration
	Cors            CorsConfig
	Meta            MetaConfig
	Swagger         SwaggerConfig
	Platform        PlatformConfig
}

type CorsConfig struct {
	Mode string
}

type MetaConfig struct {
	Title       string
	Description string
	Host        string
	BasePath    string
	Version     string
	Environment string
}

type SwaggerConfig struct {
	Enabled   bool
	Path      string
	BasicAuth BasicAuthConf
}

type PlatformConfig struct {
	Enabled   bool
	Path      string
	BasicAuth BasicAuthConf
}

type BasicAuthConf struct {
	Username string
	Password string
}
