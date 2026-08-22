// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/open-mrp/openmrp-go"
	"github.com/open-mrp/openmrp-go/internal/testutil"
	"github.com/open-mrp/openmrp-go/option"
)

func TestSaleAccountUserSalesTargetNew(t *testing.T) {
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
	_, err := client.Sales.AccountUsers.SalesTargets.New(
		context.TODO(),
		"acus_e5zu8bde0z3h",
		openmrp.SaleAccountUserSalesTargetNewParams{
			CreateSalesTargetRequest: openmrp.CreateSalesTargetRequestParam{
				AmountUnitID: "un_82bd37dae5po",
				AmountValue:  "50000.00",
				EndsAt:       time.Now(),
				StartsAt:     time.Now(),
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

func TestSaleAccountUserSalesTargetUpdate(t *testing.T) {
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
	_, err := client.Sales.AccountUsers.SalesTargets.Update(
		context.TODO(),
		"example",
		openmrp.SaleAccountUserSalesTargetUpdateParams{
			ID: "acus_e5zu8bde0z3h",
			UpsertSalesTargetRequest: openmrp.UpsertSalesTargetRequestParam{
				AmountUnitID: "un_82bd37dae5po",
				AmountValue:  "75000.00",
				EndsAt:       time.Now(),
				StartsAt:     time.Now(),
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

func TestSaleAccountUserSalesTargetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.AccountUsers.SalesTargets.List(
		context.TODO(),
		"acus_e5zu8bde0z3h",
		openmrp.SaleAccountUserSalesTargetListParams{
			Cursor: openmrp.String("cursor"),
			Limit:  openmrp.Int(0),
			Q:      openmrp.String("q"),
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
