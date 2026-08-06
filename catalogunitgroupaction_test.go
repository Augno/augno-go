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

func TestCatalogUnitGroupActionBulkUpsert(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Actions.BulkUpsert(context.TODO(), augno.CatalogUnitGroupActionBulkUpsertParams{
		BulkUpsertUnitGroupsRequest: augno.BulkUpsertUnitGroupsRequestParam{
			UnitGroups: []augno.UpsertUnitGroupInputParam{{
				BaseUnit: augno.UnitIdentifierParam{
					ID:           "un_82bd37dae5po",
					Abbreviation: "abbreviation",
					Name:         "name",
				},
				Name:  "Weight",
				Type:  augno.UpsertUnitGroupInputTypeMass,
				Notes: augno.String("notes"),
				UnitConversions: []augno.UpsertUnitGroupConversionInputParam{{
					Unit: augno.UnitIdentifierParam{
						ID:           "un_82bd37dae5po",
						Abbreviation: "abbreviation",
						Name:         "name",
					},
					DiscountPercentage: augno.Float(1),
				}},
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
