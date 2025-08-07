// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/augno-go"
	"github.com/stainless-sdks/augno-go/internal/testutil"
	"github.com/stainless-sdks/augno-go/option"
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
	getCustomerAddress, err := client.Customers.Addresses.List(context.TODO(), "REPLACE_ME")
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", getCustomerAddress.ID)
}
