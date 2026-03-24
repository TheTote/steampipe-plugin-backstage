package main

import (
	"github.com/thetote/steampipe-plugin-backstage/backstage"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Plugin entry point
func main() {
	// Start the plugin server
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: backstage.Plugin,
	})
}
