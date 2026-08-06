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

func TestCatalogUnitActionBulkUpsert(t *testing.T) {
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
	_, err := client.Catalog.Units.Actions.BulkUpsert(context.TODO(), augno.CatalogUnitActionBulkUpsertParams{
		BulkUpsertUnitsRequest: augno.BulkUpsertUnitsRequestParam{
			Units: []augno.UpsertUnitInputParam{{
				Abbreviation:      "kg",
				IsBaseUnit:        false,
				Name:              "Kilogram",
				OffsetDenominator: "1",
				OffsetNumerator:   "0",
				RatioDenominator:  "1",
				RatioNumerator:    "1000",
				Type:              augno.UpsertUnitInputTypeMass,
			}},
		},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
