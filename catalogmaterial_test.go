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

func TestCatalogMaterialNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Materials.New(context.TODO(), augno.CatalogMaterialNewParams{
		CreateMaterialRequest: augno.CreateMaterialRequestParam{
			CategoryID:   "ic_01ae7bd7bfd21ca0ab81e1357e",
			SKU:          "MAT-001",
			AttributeIDs: []string{"string"},
			Description:  augno.String("description"),
			LeadTime: augno.QuantityInputRequestParam{
				UnitID: "unit_id",
				Value:  "value",
			},
			Notes: augno.String("notes"),
			OrderPoint: augno.QuantityInputRequestParam{
				UnitID: "unit_id",
				Value:  "value",
			},
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
		Include: []string{"item"},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogMaterialGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Materials.Get(
		context.TODO(),
		"ml_014613b8f7959a091d8cc0cef4",
		augno.CatalogMaterialGetParams{
			Include: []string{"item"},
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

func TestCatalogMaterialUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Materials.Update(
		context.TODO(),
		"ml_014613b8f7959a091d8cc0cef4",
		augno.CatalogMaterialUpdateParams{
			Include: []string{"item"},
			UpdateMaterialRequest: augno.UpdateMaterialRequestParam{
				Description: augno.String("description"),
				LeadTime: augno.QuantityInputRequestParam{
					UnitID: "unit_id",
					Value:  "value",
				},
				Notes: augno.String("notes"),
				OrderPoint: augno.QuantityInputRequestParam{
					UnitID: "unit_id",
					Value:  "value",
				},
				SKU: augno.String("MAT-001-UPDATED"),
				UnitCost: augno.RateInputParam{
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

func TestCatalogMaterialListWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Materials.List(context.TODO(), augno.CatalogMaterialListParams{
		AttributeIDs: []string{"string"},
		CategoryIDs:  []string{"string"},
		Cursor:       augno.String("cursor"),
		EndDate:      augno.Time(time.Now()),
		Include:      []string{"item"},
		Limit:        augno.Int(0),
		Q:            augno.String("q"),
		StartDate:    augno.Time(time.Now()),
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogMaterialDelete(t *testing.T) {
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
	_, err := client.Catalog.Materials.Delete(context.TODO(), "ml_014613b8f7959a091d8cc0cef4")
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
