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

func TestCatalogProductNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Products.New(context.TODO(), augno.CatalogProductNewParams{
		CreateProductRequest: augno.CreateProductRequestParam{
			CategoryID:       "ic_d06g9c6yc9ck",
			SKU:              "ALM-2024-1001",
			Type:             augno.CreateProductRequestTypeSale,
			AttributeIDs:     []string{"at_rf1w295jt5ia"},
			Description:      augno.String("Wireless barcode scanner with charging cradle"),
			Notes:            augno.String("Ships with a 2-year warranty; register for extended coverage."),
			PortalVisibility: augno.CreateProductRequestPortalVisibilityVisible,
			ProductLineID:    augno.String("pdln_k9bnlgvxhxjh"),
			UnitCost: augno.RateInputParam{
				DenominatorUnitID: "un_82bd37dae5po",
				NumeratorUnitID:   "un_82bd37dae5po",
				Value:             "112.00",
			},
			UnitPrice: augno.RateInputParam{
				DenominatorUnitID: "un_82bd37dae5po",
				NumeratorUnitID:   "un_82bd37dae5po",
				Value:             "199.00",
			},
		},
		Include: []string{"product_line"},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogProductGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Products.Get(
		context.TODO(),
		"pd_07oe0r7adh2w",
		augno.CatalogProductGetParams{
			Include: []string{"product_line"},
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

func TestCatalogProductUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Products.Update(
		context.TODO(),
		"pd_07oe0r7adh2w",
		augno.CatalogProductUpdateParams{
			Include: []string{"product_line"},
			UpdateProductRequest: augno.UpdateProductRequestParam{
				Description:      augno.String("Wireless barcode scanner with charging cradle (v2)"),
				Notes:            augno.String("Firmware 2.1 improves Bluetooth pairing reliability."),
				PortalVisibility: augno.UpdateProductRequestPortalVisibilityVisible,
				SKU:              augno.String("SKU-002"),
				UnitPrice: augno.RateInputParam{
					DenominatorUnitID: "un_82bd37dae5po",
					NumeratorUnitID:   "un_82bd37dae5po",
					Value:             "219.00",
				},
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

func TestCatalogProductListWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Products.List(context.TODO(), augno.CatalogProductListParams{
		AttributeIDs:     []string{"string"},
		CategoryIDs:      []string{"string"},
		Cursor:           augno.String("cursor"),
		CustomerIDs:      []string{"string"},
		EndDate:          augno.Time(time.Now()),
		Include:          []string{"product_line"},
		Limit:            augno.Int(0),
		PortalVisibility: augno.CatalogProductListParamsPortalVisibilityVisible,
		ProductLineIDs:   []string{"string"},
		Q:                augno.String("q"),
		StartDate:        augno.Time(time.Now()),
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogProductDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Products.Delete(
		context.TODO(),
		"pd_07oe0r7adh2w",
		augno.CatalogProductDeleteParams{
			Include: []string{"product_line"},
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

func TestCatalogProductChangeProductLineWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Products.ChangeProductLine(
		context.TODO(),
		"pdln_k9bnlgvxhxjh",
		augno.CatalogProductChangeProductLineParams{
			ID:      "pd_07oe0r7adh2w",
			Include: []string{"product_line"},
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
