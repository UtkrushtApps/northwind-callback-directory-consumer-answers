package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"northwind/callbackdirectory/config"
)

type CountryRecord struct {
	Name         string   `json:"name"`
	CallingCodes []string `json:"callingCodes"`
}

type CountryPayload struct {
	Data []CountryRecord `json:"data"`
}

type CountriesClient interface {
	FetchCountry(country string) (CountryPayload, error)
}

type HTTPCountriesClient struct {
	cfg        config.Config
	httpClient *http.Client
}

func NewHTTPCountriesClient(cfg config.Config) *HTTPCountriesClient {
	return &HTTPCountriesClient{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: cfg.RequestTimeout,
		},
	}
}

func (c *HTTPCountriesClient) FetchCountry(country string) (CountryPayload, error) {
	parsed, err := url.Parse(c.cfg.UpstreamBaseURL)
	if err != nil {
		return CountryPayload{}, err
	}
	query := parsed.Query()
	query.Set("name", country)
	parsed.RawQuery = query.Encode()

	resp, err := c.httpClient.Get(parsed.String())
	if err != nil {
		return CountryPayload{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return CountryPayload{}, fmt.Errorf("upstream returned status %d", resp.StatusCode)
	}

	var payload CountryPayload
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return CountryPayload{}, err
	}
	return payload, nil
}
