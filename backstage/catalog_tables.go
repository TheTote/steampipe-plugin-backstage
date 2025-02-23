package backstage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Common columns that appear in most catalog tables
var commonColumns = []*plugin.Column{
	{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the entity"},
	{Name: "namespace", Type: proto.ColumnType_STRING, Description: "Namespace of the entity"},
	{Name: "kind", Type: proto.ColumnType_STRING, Description: "Kind of the entity"},
	{Name: "title", Type: proto.ColumnType_STRING, Description: "Display title of the entity"},
	{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the entity"},
	{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels attached to the entity"},
	{Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations on the entity"},
	{Name: "links", Type: proto.ColumnType_JSON, Description: "Links associated with the entity"},
	{Name: "metadata", Type: proto.ColumnType_JSON, Description: "Full metadata of the entity"},
	{Name: "spec", Type: proto.ColumnType_JSON, Description: "Full specification of the entity"},
}

func tableBackstageGroup() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_group",
		Description: "Groups in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listGroups,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the group"},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the group"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the entity"},
			{Name: "namespace", Type: proto.ColumnType_STRING, Description: "Namespace of the entity"},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "Display title"},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels attached to the entity"},
			{Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations on the entity"},
			{Name: "links", Type: proto.ColumnType_JSON, Description: "Links associated with the entity"},
		},
	}
}

func tableBackstageUser() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_user",
		Description: "Users in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listUsers,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the user"},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the user"},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email of the user"},
			{Name: "picture", Type: proto.ColumnType_STRING, Description: "Picture URL of the user"},
			{Name: "memberof", Type: proto.ColumnType_JSON, Description: "Groups the user belongs to"},
			{Name: "namespace", Type: proto.ColumnType_STRING, Description: "Namespace of the entity"},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels attached to the entity"},
			{Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations on the entity"},
		},
	}
}

func tableBackstageResource() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_resource",
		Description: "Backstage resource table",
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
		},
	}
}

func tableBackstageSystem() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_system",
		Description: "Backstage system table",
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
		},
	}
}

func tableBackstageDomain() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_domain",
		Description: "Backstage domain table",
		List: &plugin.ListConfig{
			Hydrate: listDomains,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the domain"},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the domain"},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the domain"},
			{Name: "namespace", Type: proto.ColumnType_STRING, Description: "Namespace of the entity"},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels attached to the entity"},
			{Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations on the entity"},
		},
	}
}

func tableBackstageLocation() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_location",
		Description: "Backstage location table",
		List: &plugin.ListConfig{
			Hydrate: listLocations,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the location"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the location"},
			{Name: "target", Type: proto.ColumnType_STRING, Description: "Target of the location"},
			{Name: "namespace", Type: proto.ColumnType_STRING, Description: "Namespace of the entity"},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels attached to the entity"},
			{Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations on the entity"},
		},
	}
}

func tableBackstageTemplate() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_template",
		Description: "Backstage template table",
		List: &plugin.ListConfig{
			Hydrate: listTemplates,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the template"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the template"},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "Title of the template"},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the template"},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the template"},
			{Name: "namespace", Type: proto.ColumnType_STRING, Description: "Namespace of the entity"},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels attached to the entity"},
			{Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations on the entity"},
		},
	}
}

func tableBackstageAPI() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_api",
		Description: "Backstage API table",
		List: &plugin.ListConfig{
			Hydrate: listAPIs,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the API"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the API"},
			{Name: "lifecycle", Type: proto.ColumnType_STRING, Description: "Lifecycle state of the API"},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the API"},
			{Name: "definition", Type: proto.ColumnType_STRING, Description: "Definition of the API"},
			{Name: "system", Type: proto.ColumnType_STRING, Description: "System the API belongs to"},
			{Name: "namespace", Type: proto.ColumnType_STRING, Description: "Namespace of the entity"},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels attached to the entity"},
			{Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations on the entity"},
		},
	}
}

func tableBackstageComponent() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_component",
		Description: "Backstage component table",
		List: &plugin.ListConfig{
			Hydrate: listComponents,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the component"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the component"},
			{Name: "lifecycle", Type: proto.ColumnType_STRING, Description: "Lifecycle state of the component"},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the component"},
			{Name: "system", Type: proto.ColumnType_STRING, Description: "System the component belongs to"},
			{Name: "namespace", Type: proto.ColumnType_STRING, Description: "Namespace of the entity"},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels attached to the entity"},
			{Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations on the entity"},
		},
	}
}

func tableBackstageEntity() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_entity",
		Description: "All entities in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listEntities,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the entity"},
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "Kind of the entity"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the entity"},
			{Name: "namespace", Type: proto.ColumnType_STRING, Description: "Namespace of the entity"},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "Display title"},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the entity"},
			{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels attached to the entity"},
			{Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations on the entity"},
			{Name: "relations", Type: proto.ColumnType_JSON, Description: "Relations to other entities"},
			{Name: "spec", Type: proto.ColumnType_JSON, Description: "The specification data of the entity"},
			{Name: "status", Type: proto.ColumnType_JSON, Description: "The status of the entity"},
		},
	}
}

func listGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Get plugin config
	config := GetConfig(d.Connection)

	if config.Host == nil || config.Token == nil {
		return nil, fmt.Errorf("host and token must be configured")
	}

	client, err := getClient(config)
	if err != nil {
		return nil, err
	}

	// Get the limit
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit < 1 {
			return nil, nil
		}
	}

	opts := &backstage.ListEntityOptions{
		Filters: []string{"kind=Group"},
		Fields:  []string{},
		Limit:   int(limit),
	}

	// Handle pagination
	var cursor string
	for {
		groups, resp, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("backstage_catalog_group.listGroups", "query_error", err)
			return nil, fmt.Errorf("error listing groups: %v", err)
		}

		for _, group := range groups {
			d.StreamListItem(ctx, group)
		}

		cursor = resp.Header.Get("Link")
		if cursor == "" {
			break
		}
		opts.After = cursor
	}

	return nil, nil
}

func listUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Get plugin config
	config := GetConfig(d.Connection)

	if config.Host == nil || config.Token == nil {
		return nil, fmt.Errorf("host and token must be configured")
	}

	plugin.Logger(ctx).Debug("backstage_catalog_user.listUsers", "config", config)

	httpClient := &http.Client{}
	client, err := backstage.NewClient(*config.Host, *config.Token, httpClient)
	if err != nil {
		plugin.Logger(ctx).Error("backstage_catalog_user.listUsers", "connection_error", err)
		return nil, fmt.Errorf("error creating backstage client: %v", err)
	}

	opts := &backstage.ListEntityOptions{
		Filters: []string{"kind=User"},
		Fields:  []string{},
	}

	users, _, err := client.Catalog.Entities.List(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("backstage_catalog_user.listUsers", "query_error", err)
		return nil, fmt.Errorf("error listing users: %v", err)
	}

	for _, user := range users {
		d.StreamListItem(ctx, user)
	}

	return nil, nil
}

func listResources(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement resource listing
	return nil, nil
}

func listSystems(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement system listing
	return nil, nil
}

func listDomains(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement domain listing
	return nil, nil
}

func listLocations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement location listing
	return nil, nil
}

func listTemplates(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement template listing
	return nil, nil
}

func listAPIs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement API listing
	return nil, nil
}

func listComponents(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement component listing
	return nil, nil
}

func listEntities(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Get plugin config
	config := GetConfig(d.Connection)

	if config.Host == nil || config.Token == nil {
		return nil, fmt.Errorf("host and token must be configured")
	}

	plugin.Logger(ctx).Debug("backstage_catalog_entity.listEntities", "config", config)

	client, err := getClient(config)
	if err != nil {
		return nil, err
	}

	// Get the limit
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit < 1 {
			return nil, nil
		}
	}

	opts := &backstage.ListEntityOptions{
		Fields: []string{},
		Limit:  int(limit),
	}

	// Handle pagination using cursor-based pagination
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

		// Check if we've reached the limit
		if limit != nil && d.RowsProcessed() >= *limit {
			return nil, nil
		}

		// Get next page cursor
		cursor = resp.Header.Get("Link")
		if cursor == "" {
			break
		}
		opts.After = cursor
	}

	return nil, nil
}
