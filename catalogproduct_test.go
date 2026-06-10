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
			CategoryID:       "ic_01ae7bd7bfd21ca0ab81e1357e",
			SKU:              "ALM-2024-1001",
			Type:             augno.CreateProductRequestTypeSale,
			AttributeIDs:     []string{"string"},
			Description:      augno.String("description"),
			Notes:            augno.String("notes"),
			PortalVisibility: augno.CreateProductRequestPortalVisibilityVisible,
			ProductLineID:    augno.String("product_line_id"),
			UnitCost: augno.RateInputParam{
				DenominatorUnitID: "denominator_unit_id",
				NumeratorUnitID:   "numerator_unit_id",
				Value:             "value",
			},
			UnitPrice: augno.RateInputParam{
				DenominatorUnitID: "denominator_unit_id",
				NumeratorUnitID:   "numerator_unit_id",
				Value:             "value",
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
		"pd_013c29ab3f1518d0004094c316",
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
		"pd_013c29ab3f1518d0004094c316",
		augno.CatalogProductUpdateParams{
			Include: []string{"product_line"},
			UpdateProductRequest: augno.UpdateProductRequestParam{
				Description:      augno.String("description"),
				Notes:            augno.String("notes"),
				PortalVisibility: augno.UpdateProductRequestPortalVisibilityVisible,
				SKU:              augno.String("SKU-002"),
				UnitPrice: augno.RateInputParam{
					DenominatorUnitID: "denominator_unit_id",
					NumeratorUnitID:   "numerator_unit_id",
					Value:             "value",
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
		"pd_013c29ab3f1518d0004094c316",
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
		"pl_01996357326a0d3f7b129542ea",
		augno.CatalogProductChangeProductLineParams{
			ID:      "pd_013c29ab3f1518d0004094c316",
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
