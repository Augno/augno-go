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

func TestOperationShipmentActionRateShopWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.Shipments.Actions.RateShop(context.TODO(), openmrp.OperationShipmentActionRateShopParams{
		RateShopRequest: openmrp.RateShopRequestParam{
			Parcels: []openmrp.ParcelInputParam{{
				Height: 6,
				Length: 12,
				Weight: 5,
				Width:  8,
			}},
			ToAddress: openmrp.AddressInputParam{
				Country:           "US",
				Name:              "Destination",
				Email:             openmrp.String("warehouse@acme.com"),
				Locality:          openmrp.String("Los Angeles"),
				Phone:             openmrp.String("555-123-4567"),
				PostalCode:        openmrp.String("90001"),
				ReceiveCalendarID: openmrp.String("receive_calendar_id"),
				State:             openmrp.String("CA"),
				StreetLine1:       openmrp.String("456 Oak Avenue"),
				StreetLine2:       openmrp.String("Suite 400"),
				Type:              openmrp.AddressInputTypeStandard,
			},
			CustomerID: openmrp.String("customer_id"),
			FromAddress: openmrp.AddressInputParam{
				Country:           "US",
				Name:              "Origin Warehouse",
				Email:             openmrp.String("warehouse@acme.com"),
				Locality:          openmrp.String("San Francisco"),
				Phone:             openmrp.String("555-123-4567"),
				PostalCode:        openmrp.String("94105"),
				ReceiveCalendarID: openmrp.String("receive_calendar_id"),
				State:             openmrp.String("CA"),
				StreetLine1:       openmrp.String("123 Main Street"),
				StreetLine2:       openmrp.String("Suite 400"),
				Type:              openmrp.AddressInputTypeStandard,
			},
			OrderTotal:     openmrp.Float(0),
			ProductLineIDs: []string{"string"},
		},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
