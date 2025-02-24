package backstage

import (
	"context"
	"fmt"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageComponent() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_component",
		Description: "Components in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listComponents,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the component"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the component"},
			{Name: "lifecycle", Type: proto.ColumnType_STRING, Description: "Lifecycle state of the component"},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the component"},
			{Name: "system", Type: proto.ColumnType_STRING, Description: "System the component belongs to"},
		},
	}
}

func listComponents(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	client, err := getClient(config)
	if err != nil {
		return nil, err
	}

	opts := &backstage.ListEntityOptions{
		Filters: []string{"kind=Component"},
		Fields:  []string{},
	}

	var cursor string
	for {
		components, resp, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("backstage_catalog_component.listComponents", "query_error", err)
			return nil, fmt.Errorf("error listing components: %v", err)
		}

		for _, component := range components {
			d.StreamListItem(ctx, component)
		}

		cursor = resp.Header.Get("Link")
		if cursor == "" {
			break
		}
	}

	return nil, nil
}
