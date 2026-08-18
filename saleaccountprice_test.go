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

func TestSaleAccountPriceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.AccountPrices.New(context.TODO(), augno.SaleAccountPriceNewParams{
		CreateAccountPriceRequest: augno.CreateAccountPriceRequestParam{
			ProductLineID: "pdln_k9bnlgvxhxjh",
			Rate: augno.RateInputParam{
				DenominatorUnitID: "un_82bd37dae5po",
				NumeratorUnitID:   "un_82bd37dae5po",
				Value:             "25.50",
			},
			RecipientAccountID: "ac_ykxoradjoeb3",
			AttributeIDs:       []string{"at_rf1w295jt5ia"},
			CategoryIDs:        []string{"ic_d06g9c6yc9ck"},
		},
		Include: []string{"recipient_account"},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleAccountPriceGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.AccountPrices.Get(
		context.TODO(),
		"acpr_7l4j483kf32p",
		augno.SaleAccountPriceGetParams{
			Include: []string{"recipient_account"},
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

func TestSaleAccountPriceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.AccountPrices.Update(
		context.TODO(),
		"acpr_7l4j483kf32p",
		augno.SaleAccountPriceUpdateParams{
			Include: []string{"recipient_account"},
			UpdateAccountPriceRequest: augno.UpdateAccountPriceRequestParam{
				AttributeIDs:  []string{"string"},
				CategoryIDs:   []string{"string"},
				ProductLineID: augno.String("product_line_id"),
				Rate: augno.RateInputParam{
					DenominatorUnitID: "un_82bd37dae5po",
					NumeratorUnitID:   "un_82bd37dae5po",
					Value:             "30.000000000000000000000000000000",
				},
				RecipientAccountID: augno.String("recipient_account_id"),
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

func TestSaleAccountPriceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.AccountPrices.List(context.TODO(), augno.SaleAccountPriceListParams{
		Cursor:             augno.String("cursor"),
		Include:            []string{"recipient_account"},
		Limit:              augno.Int(0),
		Q:                  augno.String("q"),
		RecipientAccountID: augno.String("recipient_account_id"),
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleAccountPriceDelete(t *testing.T) {
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
	_, err := client.Sales.AccountPrices.Delete(context.TODO(), "acpr_7l4j483kf32p")
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
