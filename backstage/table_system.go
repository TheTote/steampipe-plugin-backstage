package backstage

import (
	"context"
	"fmt"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageSystem() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_system",
		Description: "Systems in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listSystems,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the system"},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the system"},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the system"},
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "Domain the system belongs to"},
			{Name: "namespace", Type: proto.ColumnType_STRING, Description: "Namespace of the entity"},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels attached to the entity"},
			{Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations on the entity"},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Full metadata of the system"},
			{Name: "spec", Type: proto.ColumnType_JSON, Description: "Full specification of the system"},
		},
	}
}

func listSystems(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	client, err := getClient(config)
	if err != nil {
		return nil, err
	}

	opts := &backstage.ListEntityOptions{
		Filters: []string{"kind=System"},
		Fields:  []string{},
	}

	var cursor string
	for {
		systems, resp, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("backstage_catalog_system.listSystems", "query_error", err)
			return nil, fmt.Errorf("error listing systems: %v", err)
		}

		for _, system := range systems {
			d.StreamListItem(ctx, system)
		}

		cursor = resp.Header.Get("Link")
		if cursor == "" {
			break
		}
	}

	return nil, nil
}
