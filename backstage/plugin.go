package backstage

import (
	"context"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type backstageConfig struct {
	Host  *string `cty:"host"`
	Token *string `cty:"token"`
}

func ConfigInstance() interface{} {
	return &backstageConfig{}
}

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-backstage",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema: map[string]*schema.Attribute{
				"host": {
					Type:                schema.TypeString,
					Required:            true,
					MarkdownDescription: "Backstage instance URL (e.g., https://backstage.example.com)",
				},
				"token": {
					Type:                schema.TypeString,
					Required:            true,
					MarkdownDescription: "Backstage API token for authentication",
					Sensitive:           true,
				},
			},
		},
		TableMap: map[string]*plugin.Table{
			"backstage_catalog_entity":    tableBackstageEntity(),
			"backstage_catalog_component": tableBackstageComponent(),
			"backstage_catalog_template":  tableBackstageTemplate(),
			"backstage_catalog_api":       tableBackstageAPI(),
			"backstage_catalog_group":     tableBackstageGroup(),
			"backstage_catalog_user":      tableBackstageUser(),
			"backstage_catalog_resource":  tableBackstageResource(),
			"backstage_catalog_system":    tableBackstageSystem(),
			"backstage_catalog_domain":    tableBackstageDomain(),
			"backstage_catalog_location":  tableBackstageLocation(),
		},
	}
	return p
}

func GetConfig(connection *plugin.Connection) backstageConfig {
	config := backstageConfig{}

	// Load from config file first
	if connection != nil {
		if config, ok := connection.Config.(backstageConfig); ok {
			return config
		}
	}

	// Override with environment variables if present
	if host := os.Getenv("BACKSTAGE_HOST"); host != "" {
		config.Host = &host
	}
	if token := os.Getenv("BACKSTAGE_TOKEN"); token != "" {
		config.Token = &token
	}

	return config
}
