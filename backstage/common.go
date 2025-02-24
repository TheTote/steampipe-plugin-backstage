package backstage

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Common columns that appear in most catalog tables
var commonColumns = []*plugin.Column{
	// Required metadata fields
	{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the entity"},
	{Name: "namespace", Type: proto.ColumnType_STRING, Description: "The namespace the entity belongs to"},
	{Name: "kind", Type: proto.ColumnType_STRING, Description: "The kind of the entity"},

	// Common fields
	{Name: "metadata", Type: proto.ColumnType_JSON, Description: "The full metadata of the entity"},
	{Name: "spec", Type: proto.ColumnType_JSON, Description: "The specification data of the entity"},
	{Name: "relations", Type: proto.ColumnType_JSON, Description: "The relations of the entity to other entities"},

	// Optional metadata fields
	{Name: "title", Type: proto.ColumnType_STRING, Description: "A display name of the entity"},
	{Name: "description", Type: proto.ColumnType_STRING, Description: "A description of the entity"},
	{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels attached to the entity"},
	{Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations attached to the entity"},
	{Name: "tags", Type: proto.ColumnType_JSON, Description: "A list of tags attached to the entity"},
	{Name: "links", Type: proto.ColumnType_JSON, Description: "A list of external hyperlinks related to the entity"},
}
