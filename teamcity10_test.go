package teamcity10

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/bmizerany/assert"
)

func TestSetup(t *testing.T) {
	t.Run("GetToken", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), TeamCityCreds, os.Getenv("TC_CREDS"))
		ctx = context.WithValue(ctx, TeamCityBaseURL, os.Getenv("TC_BASE"))

		t.Run("200Response", func(t *testing.T) {
			tok, err := GetToken(ctx)
			assert.Equal(t, nil, err)
			assert.NotEqual(t, nil, tok)
		})
	})

	t.Run("GetToken", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), TeamCityCreds, os.Getenv("TC_CREDS"))
		ctx = context.WithValue(ctx, TeamCityBaseURL, os.Getenv("TC_BASE"))

		t.Run("200Response", func(t *testing.T) {
			tok, err := GetToken(ctx)
			assert.Equal(t, nil, err)
			assert.NotEqual(t, nil, tok)
		})
	})

	t.Run("GetBuilds", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), TeamCityCreds, os.Getenv("TC_CREDS"))
		ctx = context.WithValue(ctx, TeamCityBaseURL, os.Getenv("TC_BASE"))

		testID := os.Getenv("TEST_ID")
		t.Run("200Response", func(t *testing.T) {
			builds, err := GetBuilds(ctx, fmt.Sprint(testID))
			assert.Equal(t, nil, err)
			assert.Equal(t, nil, builds)
		})
	})
}

func TestTriggerBuild(t *testing.T) {
	ctx := context.WithValue(context.Background(), TeamCityCreds, os.Getenv("TC_CREDS"))
	ctx = context.WithValue(ctx, TeamCityBaseURL, os.Getenv("TC_BASE"))

	t.Run("Success", func(t *testing.T) {
		resp, err := TriggerBuild(ctx, os.Getenv("TEST_ID"), "master", os.Getenv("TEST_PROJ"))
		assert.Equal(t, nil, err)
		assert.Equal(t, resp, true)
	})
}
