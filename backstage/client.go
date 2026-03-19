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

	httpClient := &http.Client{
		Transport: &tokenRoundTripper{
			token:  *config.Token,
			client: http.DefaultTransport,
		},
	}
	client, err := backstage.NewClient(*config.Host, "", httpClient)
	if err != nil {
		return nil, fmt.Errorf("error creating backstage client: %v", err)
	}

	return client, nil
}

type tokenRoundTripper struct {
	token  string
	client http.RoundTripper
}

func (t *tokenRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.token))
	return t.client.RoundTrip(req)
}
