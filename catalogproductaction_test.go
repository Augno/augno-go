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

func TestCatalogProductActionBulkUpsert(t *testing.T) {
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
	_, err := client.Catalog.Products.Actions.BulkUpsert(context.TODO(), augno.CatalogProductActionBulkUpsertParams{
		BulkUpsertProductsRequest: augno.BulkUpsertProductsRequestParam{
			Products: []augno.UpsertProductInputParam{{
				Category: augno.ObjectIdentifierParam{
					ID:   "ic_d06g9c6yc9ck",
					Name: "name",
				},
				Properties: []augno.UpsertProductPropertyParam{{
					Name:  "name",
					Value: "value",
				}},
				SKU:              "ALM-2024-1001",
				Description:      augno.String("description"),
				Notes:            augno.String("notes"),
				PortalVisibility: augno.UpsertProductInputPortalVisibilityVisible,
				ProductLine: augno.ObjectIdentifierParam{
					ID:   "id",
					Name: "name",
				},
				Type: augno.UpsertProductInputTypeSale,
				UnitCost: augno.RateInputParam{
					DenominatorUnitID: "denominator_unit_id",
					NumeratorUnitID:   "numerator_unit_id",
					Value:             "value",
				},
				UnitPrice: augno.RateInputParam{
					DenominatorUnitID: "denominator_unit_id",
					NumeratorUnitID:   "numerator_unit_id",
					Value:             "value",
				},
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
