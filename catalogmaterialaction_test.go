// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/open-mrp/openmrp-go"
	"github.com/open-mrp/openmrp-go/internal/testutil"
	"github.com/open-mrp/openmrp-go/option"
)

func TestCatalogMaterialActionBulkUpsertWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Catalog.Materials.Actions.BulkUpsert(context.TODO(), openmrp.CatalogMaterialActionBulkUpsertParams{
		BulkUpsertMaterialsRequest: openmrp.BulkUpsertMaterialsRequestParam{
			Materials: []openmrp.UpsertMaterialInputParam{{
				Category: openmrp.ObjectIdentifierParam{
					ID:   "ic_d06g9c6yc9ck",
					Name: "name",
				},
				Properties: []openmrp.UpsertMaterialPropertyParam{{
					Name:  "name",
					Value: "value",
				}},
				SKU:         "MAT-001",
				Description: openmrp.String("description"),
				LeadTime: openmrp.QuantityInputRequestParam{
					UnitID: "unit_id",
					Value:  "value",
				},
				Notes: openmrp.String("notes"),
				OrderPoint: openmrp.QuantityInputRequestParam{
					UnitID: "unit_id",
					Value:  "value",
				},
				UnitCost: openmrp.RateInputParam{
					DenominatorUnitID: "denominator_unit_id",
					NumeratorUnitID:   "numerator_unit_id",
					Value:             "value",
				},
				UnitPrice: openmrp.RateInputParam{
					DenominatorUnitID: "denominator_unit_id",
					NumeratorUnitID:   "numerator_unit_id",
					Value:             "value",
				},
			}},
		},
		Include: []string{"created_by"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
