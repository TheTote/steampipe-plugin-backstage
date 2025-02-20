package backstage

import (
	"context"

	"net/http"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageGroup() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_group",
		Description: "Backstage group table",
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
		Name:        "backstage_user",
		Description: "Backstage user table",
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
		Name:        "backstage_resource",
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
		Name:        "backstage_system",
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
		Name:        "backstage_domain",
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
		Name:        "backstage_location",
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
		Name:        "backstage_template",
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
		Name:        "backstage_api",
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
		Name:        "backstage_component",
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
		Name:        "backstage_entity",
		Description: "Backstage entity table",
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
	// Implement your group listing logic here
	return nil, nil
}

func listUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Implement your user listing logic here
	return nil, nil
}

func listResources(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Implement your resource listing logic here
	return nil, nil
}

func listSystems(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Implement your system listing logic here
	return nil, nil
}

func listDomains(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Implement your domain listing logic here
	return nil, nil
}

func listLocations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Implement your location listing logic here
	return nil, nil
}

func listTemplates(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Implement your template listing logic here
	return nil, nil
}

func listAPIs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Implement your API listing logic here
	return nil, nil
}

func listComponents(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Implement your component listing logic here
	return nil, nil
}

func listEntities(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	baseURL := "http://your-backstage-url"
	httpClient := &http.Client{}
	client, err := backstage.NewClient(baseURL, "default", httpClient)
	if err != nil {
		return nil, err
	}

	entities, _, err := client.Catalog.Entities.List(ctx, &backstage.ListEntityOptions{
		Filters: []string{},
		Fields:  []string{},
		Order:   []backstage.ListEntityOrder{{Direction: backstage.OrderDescending, Field: "metadata.name"}},
	})
	if err != nil {
		return nil, err
	}

	for _, entity := range entities {
		d.StreamListItem(ctx, entity)
	}

	return nil, nil
}
