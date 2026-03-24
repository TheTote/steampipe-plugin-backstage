package backstage

import (
	"context"
	"fmt"
	"os"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// connect establishes a connection to the Backstage API using the provided token.
func connect(ctx context.Context, d *plugin.QueryData) *backstage.Client {

	cacheKey := "devlake"

	if cached, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		if client, cOk := cached.(*backstage.Client); cOk {
			return client
		}
	}

	config := GetConfig(d.Connection)

	var host string
	var token string

	if config.Host != nil {
		host = *config.Host
	}
	if config.Token != nil {
		token = *config.Token
	}

	if h := os.Getenv("BACKSTAGE_HOST"); h != "" {
		host = h
	}
	if t := os.Getenv("BACKSTAGE_TOKEN"); t != "" {
		token = t
	}

	client, err := getClient(host, token)
	if err != nil {
		panic(fmt.Sprintf("error creating backstage client: %v", err))
		return nil
	}
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client
}
