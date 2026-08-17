// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/augno/augno-go"
	"github.com/augno/augno-go/internal/testutil"
	"github.com/augno/augno-go/option"
)

func TestOperationLocationActionBulkUpsertWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Operations.Locations.Actions.BulkUpsert(context.TODO(), augno.OperationLocationActionBulkUpsertParams{
		BulkUpsertLocationsRequest: augno.BulkUpsertLocationsRequestParam{
			Locations: []augno.UpsertLocationInputParam{{
				Name: "Warehouse A",
				Type: augno.LocationTypeCodeBuilding,
				Children: []augno.ObjectIdentifierParam{{
					ID:   "id",
					Name: "name",
				}},
				Parent: augno.ObjectIdentifierParam{
					ID:   "id",
					Name: "name",
				},
			}},
		},
		Include: []string{"created_by"},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
