// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno_test

import (
	"context"
	"os"
	"testing"

	"github.com/Augno/go-sdk"
	"github.com/Augno/go-sdk/internal/testutil"
	"github.com/Augno/go-sdk/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	t.Skip("Some required params are not supported yet.")
	response, err := client.Auth.RefreshToken(context.TODO())
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response)
}
