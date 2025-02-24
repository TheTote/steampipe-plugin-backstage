package backstage

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-backstage",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
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
