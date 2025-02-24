package backstage

import (
	"fmt"
	"net/http"

	"github.com/datolabs-io/go-backstage/v3"
)

func getClient(config BackstageConfig) (*backstage.Client, error) {
	if config.Host == nil || config.Token == nil {
		return nil, fmt.Errorf("host and token must be configured")
	}

	httpClient := &http.Client{}
	client, err := backstage.NewClient(*config.Host, *config.Token, httpClient)
	if err != nil {
		return nil, fmt.Errorf("error creating backstage client: %v", err)
	}

	return client, nil
}
