package backstage

import (
	"context"
	"fmt"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageResource() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_resource",
		Description: "Resources in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listResources,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the resource"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the resource"},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the resource"},
			{Name: "system", Type: proto.ColumnType_STRING, Description: "System the resource belongs to"},
			{Name: "namespace", Type: proto.ColumnType_STRING, Description: "Namespace of the entity"},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels attached to the entity"},
			{Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations on the entity"},
			{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Full metadata of the resource"},
			{Name: "spec", Type: proto.ColumnType_JSON, Description: "Full specification of the resource"},
		},
	}
}

func listResources(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	client, err := getClient(config)
	if err != nil {
		return nil, err
	}

	opts := &backstage.ListEntityOptions{
		Filters: []string{"kind=Resource"},
		Fields:  []string{},
	}

	var cursor string
	for {
		resources, resp, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("backstage_catalog_resource.listResources", "query_error", err)
			return nil, fmt.Errorf("error listing resources: %v", err)
		}

		for _, resource := range resources {
			d.StreamListItem(ctx, resource)
		}

		cursor = resp.Header.Get("Link")
		if cursor == "" {
			break
		}
	}

	return nil, nil
}
