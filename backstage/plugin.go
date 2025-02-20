package backstage

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type backstageConfig struct {
	BaseURL  *string `cty:"base_url"`
	ApiToken *string `cty:"api_token"`
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
				"base_url": {
					Type:     schema.TypeString,
					Required: true,
				},
				"api_token": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
		TableMap: map[string]*plugin.Table{
			"backstage_entity":    tableBackstageEntity(),
			"backstage_component": tableBackstageComponent(),
			"backstage_template":  tableBackstageTemplate(),
			"backstage_api":       tableBackstageAPI(),
			"backstage_group":     tableBackstageGroup(),
			"backstage_user":      tableBackstageUser(),
			"backstage_resource":  tableBackstageResource(),
			"backstage_system":    tableBackstageSystem(),
			"backstage_domain":    tableBackstageDomain(),
			"backstage_location":  tableBackstageLocation(),
		},
	}
	return p
}
