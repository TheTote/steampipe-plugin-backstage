package backstage

import (
	"context"
	"fmt"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageEntity() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_entity",
		Description: "Generic entities in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listEntities,
		},
		Columns: commonColumns,
	}
}

func listEntities(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	if config.Host == nil || config.Token == nil {
		return nil, fmt.Errorf("host and token must be configured")
	}

	client, err := getClient(config)
	if err != nil {
		return nil, err
	}

	opts := &backstage.ListEntityOptions{
		Fields: []string{},
	}

	var cursor string
	for {
		entities, resp, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("backstage_catalog_entity.listEntities", "query_error", err)
			return nil, fmt.Errorf("error listing entities: %v", err)
		}

		for _, entity := range entities {
			d.StreamListItem(ctx, entity)
		}

		cursor = resp.Header.Get("Link")
		if cursor == "" {
			break
		}
	}

	return nil, nil
}
