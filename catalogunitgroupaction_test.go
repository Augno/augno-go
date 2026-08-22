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

func TestCatalogUnitGroupActionBulkUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Actions.BulkUpsert(context.TODO(), openmrp.CatalogUnitGroupActionBulkUpsertParams{
		BulkUpsertUnitGroupsRequest: openmrp.BulkUpsertUnitGroupsRequestParam{
			UnitGroups: []openmrp.UpsertUnitGroupInputParam{{
				BaseUnit: openmrp.UnitIdentifierParam{
					ID:           "un_82bd37dae5po",
					Abbreviation: "abbreviation",
					Name:         "name",
				},
				Name:  "Weight",
				Type:  openmrp.UpsertUnitGroupInputTypeMass,
				Notes: openmrp.String("notes"),
				UnitConversions: []openmrp.UpsertUnitGroupConversionInputParam{{
					Unit: openmrp.UnitIdentifierParam{
						ID:           "un_82bd37dae5po",
						Abbreviation: "abbreviation",
						Name:         "name",
					},
					DiscountPercentage: openmrp.Float(1),
				}},
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
