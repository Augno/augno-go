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

func TestSaleAccountPriceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.AccountPrices.New(context.TODO(), openmrp.SaleAccountPriceNewParams{
		CreateAccountPriceRequest: openmrp.CreateAccountPriceRequestParam{
			ProductLineID: "pdln_k9bnlgvxhxjh",
			Rate: openmrp.RateInputParam{
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
		var apierr *openmrp.Error
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
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Sales.AccountPrices.Get(
		context.TODO(),
		"acpr_7l4j483kf32p",
		openmrp.SaleAccountPriceGetParams{
			Include: []string{"recipient_account"},
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

func TestSaleAccountPriceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.AccountPrices.Update(
		context.TODO(),
		"acpr_7l4j483kf32p",
		openmrp.SaleAccountPriceUpdateParams{
			Include: []string{"recipient_account"},
			UpdateAccountPriceRequest: openmrp.UpdateAccountPriceRequestParam{
				AttributeIDs:  []string{"string"},
				CategoryIDs:   []string{"string"},
				ProductLineID: openmrp.String("product_line_id"),
				Rate: openmrp.RateInputParam{
					DenominatorUnitID: "un_82bd37dae5po",
					NumeratorUnitID:   "un_82bd37dae5po",
					Value:             "30.000000000000000000000000000000",
				},
				RecipientAccountID: openmrp.String("recipient_account_id"),
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

func TestSaleAccountPriceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.AccountPrices.List(context.TODO(), openmrp.SaleAccountPriceListParams{
		Cursor:             openmrp.String("cursor"),
		Include:            []string{"recipient_account"},
		Limit:              openmrp.Int(0),
		Q:                  openmrp.String("q"),
		RecipientAccountID: openmrp.String("recipient_account_id"),
	})
	if err != nil {
		var apierr *openmrp.Error
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
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Sales.AccountPrices.Delete(context.TODO(), "acpr_7l4j483kf32p")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
