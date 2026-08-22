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

func TestCatalogItemInventoryUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Items.Inventory.Update(
		context.TODO(),
		"it_pej07ckhvu62",
		openmrp.CatalogItemInventoryUpdateParams{
			UpdateItemInventoryRequest: openmrp.UpdateItemInventoryRequestParam{
				Quantity: openmrp.QuantityInputParam{
					UnitID: "un_82bd37dae5po",
					Value:  "10.5",
				},
				CustomerID: openmrp.String("ac_opnlh43ymyee"),
				LocationID: openmrp.String("lc_yonnys0hx3ju"),
				LotNumber:  openmrp.String("lot_number"),
				Operation:  openmrp.UpdateItemInventoryRequestOperationAdjust,
			},
		},
	)
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogItemInventoryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Items.Inventory.List(
		context.TODO(),
		"it_pej07ckhvu62",
		openmrp.CatalogItemInventoryListParams{
			Include: []string{"on_hand"},
		},
	)
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
