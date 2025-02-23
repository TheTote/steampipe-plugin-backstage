package backstage

import (
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type backstageConfig struct {
	Host  *string `cty:"host"`
	Token *string `cty:"token"`
}

func ConfigInstance() interface{} {
	return &backstageConfig{}
}

func GetConfig(connection *plugin.Connection) backstageConfig {
	config := backstageConfig{}

	if connection != nil {
		if config, ok := connection.Config.(backstageConfig); ok {
			return config
		}
	}

	if host := os.Getenv("BACKSTAGE_HOST"); host != "" {
		config.Host = &host
	}
	if token := os.Getenv("BACKSTAGE_TOKEN"); token != "" {
		config.Token = &token
	}

	return config
}
