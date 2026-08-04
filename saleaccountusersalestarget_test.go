// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/augno/augno-go"
	"github.com/augno/augno-go/internal/testutil"
	"github.com/augno/augno-go/option"
)

func TestSaleAccountUserSalesTargetNew(t *testing.T) {
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
	_, err := client.Sales.AccountUsers.SalesTargets.New(
		context.TODO(),
		"acus_e5zu8bde0z3h",
		augno.SaleAccountUserSalesTargetNewParams{
			CreateSalesTargetRequest: augno.CreateSalesTargetRequestParam{
				AmountUnitID: "un_82bd37dae5po",
				AmountValue:  "50000.00",
				EndDate:      time.Now(),
				StartDate:    time.Now(),
			},
		},
	)
	if err != nil {
		var apierr *augno.Error
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
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Sales.AccountUsers.SalesTargets.Update(
		context.TODO(),
		"example",
		augno.SaleAccountUserSalesTargetUpdateParams{
			ID: "acus_e5zu8bde0z3h",
			UpsertSalesTargetRequest: augno.UpsertSalesTargetRequestParam{
				AmountUnitID: "un_82bd37dae5po",
				AmountValue:  "75000.00",
				EndDate:      time.Now(),
				StartDate:    time.Now(),
			},
		},
	)
	if err != nil {
		var apierr *augno.Error
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
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Sales.AccountUsers.SalesTargets.List(
		context.TODO(),
		"acus_e5zu8bde0z3h",
		augno.SaleAccountUserSalesTargetListParams{
			Cursor: augno.String("cursor"),
			Limit:  augno.Int(0),
			Q:      augno.String("q"),
		},
	)
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
