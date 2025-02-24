package backstage

import (
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
