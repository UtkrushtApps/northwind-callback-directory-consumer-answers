package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	UpstreamBaseURL string
	RequestTimeout  time.Duration
}

func Load() Config {
	baseURL := os.Getenv("NORTHWIND_UPSTREAM_BASE_URL")
	if baseURL == "" {
		baseURL = "https://jsonmock.hackerrank.com/api/countries"
	}

	timeout := 5 * time.Second
	if raw := os.Getenv("NORTHWIND_REQUEST_TIMEOUT_MS"); raw != "" {
		if ms, err := strconv.Atoi(raw); err == nil && ms > 0 {
			timeout = time.Duration(ms) * time.Millisecond
		}
	}

	return Config{
		UpstreamBaseURL: baseURL,
		RequestTimeout:  timeout,
	}
}
