package teamcity10

import (
	"context"
	"errors"
	"fmt"
)

type config struct {
	BaseURL string
}

func getConfig(ctx context.Context) (*config, error) {
	if ctx == nil {
		return nil, errors.New("no context provided  - cannot generate config")
	}

	baseURL := fmt.Sprint(ctx.Value(TeamCityBaseURL))
	if len(baseURL) < 1 {
		return nil, errors.New("no baseURL in context - cannot generate config")
	}

	return &config{BaseURL: baseURL}, nil
}
