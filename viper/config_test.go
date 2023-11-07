package main

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestLoadConfig tests the loading of the configuration file into the Config struct.
func TestLoadConfig(t *testing.T) {
	// Set up Viper to read from the test file
	viper.SetConfigType("yaml")
	viper.SetConfigName("config_test") // Use a dedicated test config file
	viper.AddConfigPath(".")           // The path to look for the configuration file in

	// Attempt to read the configuration
	err := viper.ReadInConfig()
	assert.NoError(t, err, "Failed to read the configuration file")

	var config Config
	err = viper.Unmarshal(&config)
	assert.NoError(t, err, "Failed to unmarshal configuration")

	// Assert that the values are what we expect
	assert.Equal(t, 9233, config.HttpServer.Port, "HttpServer.Port does not match")
	assert.Equal(t, 2*time.Minute, config.HttpServer.ShutdownTimeout, "HttpServer.ShutdownTimeout does not match")

	// Add more assertions as needed for other configuration parameters
}

// You would typically have a separate test configuration file named `config_test.yaml`
// with the values you want to test against.
