package backstage

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Common columns that appear in most catalog tables
var commonColumns = []*plugin.Column{

	{Name: "kind", Type: proto.ColumnType_STRING, Description: "The kind of the entity"},
	// Common JSON fields
	{Name: "spec", Type: proto.ColumnType_JSON, Description: "The specification data of the entity"},
	{Name: "metadata", Type: proto.ColumnType_JSON, Description: "The full metadata of the entity"},

	// {Name: "relations", Type: proto.ColumnType_JSON, Description: "The relations of the entity to other entities"},

	// // Required metadata fields
	// {Name: "namespace", Type: proto.ColumnType_STRING, Description: "The namespace the entity belongs to"},
	// // Optional metadata fields
	// {Name: "title", Type: proto.ColumnType_STRING, Description: "A display name of the entity"},
	// {Name: "description", Type: proto.ColumnType_STRING, Description: "A description of the entity"},
	// {Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels attached to the entity"},
	// {Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations attached to the entity"},
	// {Name: "tags", Type: proto.ColumnType_JSON, Description: "A list of tags attached to the entity"},
	// {Name: "links", Type: proto.ColumnType_JSON, Description: "A list of external hyperlinks related to the entity"},

	// // Required spec fields
	// {Name: "type", Type: proto.ColumnType_STRING, Description: "The type of the entity"},

	// // Optional spec fields
	// {Name: "owner", Type: proto.ColumnType_STRING, Description: "The owner of the entity"},
	// {Name: "lifecycle", Type: proto.ColumnType_STRING, Description: "The lifecycle of the entity"},
	// {Name: "system", Type: proto.ColumnType_STRING, Description: "The system the entity belongs to"},
}

var commonFields = []string{
	"kind",
	"metadata.name",
	"metadata.namespace",
	"metadata.title",
	"metadata.description",
	// "metadata.labels",
	// "metadata.annotations",
	"metadata.tags",
	// "metadata.links",
	"spec.type",
	"spec.lifecycle",
	"spec.owner",
	"spec.system",
	"spec.domain",
}
