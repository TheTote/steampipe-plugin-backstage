package backstage

import (
	"context"
	"fmt"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageAPI() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_api",
		Description: "APIs in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listAPIs,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the API"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the API"},
			{Name: "lifecycle", Type: proto.ColumnType_STRING, Description: "Lifecycle state of the API"},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the API"},
			{Name: "definition", Type: proto.ColumnType_JSON, Description: "API definition"},
		},
	}
}

func listAPIs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	client, err := getClient(config)
	if err != nil {
		return nil, err
	}

	opts := &backstage.ListEntityOptions{
		Filters: []string{"kind=API"},
		Fields:  []string{},
	}

	var cursor string
	for {
		apis, resp, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("backstage_catalog_api.listAPIs", "query_error", err)
			return nil, fmt.Errorf("error listing APIs: %v", err)
		}

		for _, api := range apis {
			d.StreamListItem(ctx, api)
		}

		cursor = resp.Header.Get("Link")
		if cursor == "" {
			break
		}
	}

	return nil, nil
}
