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

func TestOperationShipmentActionRateShopWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.Shipments.Actions.RateShop(context.TODO(), augno.OperationShipmentActionRateShopParams{
		RateShopRequest: augno.RateShopRequestParam{
			Parcels: []augno.ParcelInputParam{{
				Height: 6,
				Length: 12,
				Weight: 5,
				Width:  8,
			}},
			ToAddress: augno.AddressInputParam{
				Country:           "US",
				Name:              "Destination",
				Email:             augno.String("warehouse@acme.com"),
				Locality:          augno.String("Los Angeles"),
				Phone:             augno.String("555-123-4567"),
				PostalCode:        augno.String("90001"),
				ReceiveCalendarID: augno.String("receive_calendar_id"),
				State:             augno.String("CA"),
				StreetLine1:       augno.String("456 Oak Avenue"),
				StreetLine2:       augno.String("Suite 400"),
				Type:              augno.AddressInputTypeStandard,
			},
			CustomerID: augno.String("customer_id"),
			FromAddress: augno.AddressInputParam{
				Country:           "US",
				Name:              "Origin Warehouse",
				Email:             augno.String("warehouse@acme.com"),
				Locality:          augno.String("San Francisco"),
				Phone:             augno.String("555-123-4567"),
				PostalCode:        augno.String("94105"),
				ReceiveCalendarID: augno.String("receive_calendar_id"),
				State:             augno.String("CA"),
				StreetLine1:       augno.String("123 Main Street"),
				StreetLine2:       augno.String("Suite 400"),
				Type:              augno.AddressInputTypeStandard,
			},
			OrderTotal:     augno.Float(0),
			ProductLineIDs: []string{"string"},
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
