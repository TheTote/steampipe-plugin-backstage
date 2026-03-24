package backstage

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type BackstageConfig struct {
	Host  *string `cty:"host"`
	Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"host": {
		Type:     schema.TypeString,
		Required: true,
	},
	"token": {
		Type:     schema.TypeString,
		Required: true,
	},
}

func ConfigInstance() interface{} {
	return &BackstageConfig{}
}

func GetConfig(connection *plugin.Connection) BackstageConfig {
	if connection == nil || connection.Config == nil {
		return BackstageConfig{}
	}
	config, _ := connection.Config.(BackstageConfig)

	return config
}
