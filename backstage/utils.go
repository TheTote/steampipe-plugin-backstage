package backstage

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// connect establishes a connection to the Backstage API using the provided token.
func connect(ctx context.Context, d *plugin.QueryData) (*backstage.Client, error) {
	logger := plugin.Logger(ctx)

	// Get Backstage API Token
	token := os.Getenv("BACKSTAGE_TOKEN")
	backstageConfig := GetConfig(d.Connection)

	logger.Debug("backstage.connect", "host", backstageConfig.Host)

	if backstageConfig.Host == nil {
		logger.Error("backstage.connect", "connection_error", "host configuration is missing")
		return nil, fmt.Errorf("host configuration is missing")
	}

	if backstageConfig.Token != nil {
		token = *backstageConfig.Token
		logger.Debug("backstage.connect", "token_source", "config")
	} else {
		logger.Debug("backstage.connect", "token_source", "environment")
	}

	if token == "" {
		logger.Error("backstage.connect", "connection_error", "token is required but not provided")
		return nil, fmt.Errorf("token is required but not provided")
	}

	httpClient := &http.Client{}
	client, err := backstage.NewClient(*backstageConfig.Host, token, httpClient)
	if err != nil {
		logger.Error("backstage.connect", "client_error", err)
		return nil, err
	}

	logger.Debug("backstage.connect", "status", "connection successful")
	return client, nil
}
