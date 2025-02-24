package backstage

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageUser() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_user",
		Description: "Users in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listUsers,
		},
		Columns: append(commonColumns, []*plugin.Column{
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email of the user"},
			{Name: "picture", Type: proto.ColumnType_STRING, Description: "Picture URL of the user"},
			{Name: "memberof", Type: proto.ColumnType_JSON, Description: "Groups the user belongs to"},
		}...), // Union of commonColumns and specific columns for users
	}
}

func listUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Existing implementation...
	return nil, nil
}
