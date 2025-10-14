package container

import (
	"Weather-API-Application/internal/client"
	"Weather-API-Application/internal/config"
	"Weather-API-Application/internal/services/weather_service"
)

// Container holds all application dependencies
type Container struct {
	Config           *config.Config
	WeatherClient    client.WeatherClient
	WeatherService   weather_service.WeatherService
}

// NewContainer creates a new dependency injection container
func NewContainer(cfg *config.Config) *Container {
	container := &Container{
		Config: cfg,
	}

	// Initialize clients
	container.initClients()
	
	// Initialize services
	container.initServices()

	return container
}

// initClients initializes all clients
func (c *Container) initClients() {
	clientFactory := client.NewWeatherClientFactory()
	c.WeatherClient = clientFactory.CreateWeatherClient(c.Config)
}

// initServices initializes all services with their dependencies
func (c *Container) initServices() {
	serviceFactory := weather_service.NewWeatherServiceFactory()
	c.WeatherService = serviceFactory.CreateWeatherService(c.WeatherClient)
}
