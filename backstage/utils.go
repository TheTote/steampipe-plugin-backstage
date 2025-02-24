package backstage

import (
	"context"
	"net/http"
	"os"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// connect establishes a connection to the Backstage API using the provided token.
func connect(ctx context.Context, d *plugin.QueryData) (*backstage.Client, error) {
	// Get Backstage API Token
	token := os.Getenv("BACKSTAGE_TOKEN")
	backstageConfig := GetConfig(d.Connection)
	if backstageConfig.Token != nil {
		token = *backstageConfig.Token
	}

	httpClient := &http.Client{}
	client, err := backstage.NewClient(*backstageConfig.Host, token, httpClient)
	if err != nil {
		return nil, err
	}
	return client, nil
}
