package conf

import (
	"strings"

	"github.com/xcode75/xcore/app/commander"
	loggerservice "github.com/xcode75/xcore/app/log/command"
	observatoryservice "github.com/xcode75/xcore/app/observatory/command"
	handlerservice "github.com/xcode75/xcore/app/proxyman/command"
	statsservice "github.com/xcode75/xcore/app/stats/command"
	ruleservice "github.com/xcode75/xcore/app/router/command"
	"github.com/xcode75/xcore/common/serial"
)

type APIConfig struct {
	Tag      string   `json:"tag"`
	Services []string `json:"services"`
}

func (c *APIConfig) Build() (*commander.Config, error) {
	if c.Tag == "" {
		return nil, newError("API tag can't be empty.")
	}

	services := make([]*serial.TypedMessage, 0, 16)
	for _, s := range c.Services {
		switch strings.ToLower(s) {
		case "reflectionservice":
			services = append(services, serial.ToTypedMessage(&commander.ReflectionConfig{}))
		case "handlerservice":
			services = append(services, serial.ToTypedMessage(&handlerservice.Config{}))
		case "loggerservice":
			services = append(services, serial.ToTypedMessage(&loggerservice.Config{}))
		case "statsservice":
			services = append(services, serial.ToTypedMessage(&statsservice.Config{}))
		case "observatoryservice":
			services = append(services, serial.ToTypedMessage(&observatoryservice.Config{}))
		case "ruleservice":
			services = append(services, serial.ToTypedMessage(&ruleservice.Config{}))	
		}
	}

	return &commander.Config{
		Tag:     c.Tag,
		Service: services,
	}, nil
}
