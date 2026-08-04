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

func TestCatalogUnitGroupUnitNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Units.New(
		context.TODO(),
		"ug_andst6m79n41",
		augno.CatalogUnitGroupUnitNewParams{
			CreateUnitGroupUnitRequest: augno.CreateUnitGroupUnitRequestParam{
				UnitID:                   "un_82bd37dae5po",
				CustomerPortalVisibility: augno.CreateUnitGroupUnitRequestCustomerPortalVisibilityVisible,
				DiscountFixed:            augno.Float(0),
				DiscountPercentage:       augno.Float(1),
			},
			Include: []string{"unit"},
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

func TestCatalogUnitGroupUnitGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Units.Get(
		context.TODO(),
		"un_82bd37dae5po",
		augno.CatalogUnitGroupUnitGetParams{
			UnitGroupID: "ug_andst6m79n41",
			Include:     []string{"unit"},
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

func TestCatalogUnitGroupUnitUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Units.Update(
		context.TODO(),
		"un_82bd37dae5po",
		augno.CatalogUnitGroupUnitUpdateParams{
			UnitGroupID: "ug_andst6m79n41",
			Include:     []string{"unit"},
			UpdateUnitGroupUnitRequest: augno.UpdateUnitGroupUnitRequestParam{
				CustomerPortalVisibility: augno.UpdateUnitGroupUnitRequestCustomerPortalVisibilityVisible,
				DiscountFixed:            augno.Float(2.5),
				DiscountPercentage:       augno.Float(0.9),
				UnitID:                   augno.String("un_82bd37dae5po"),
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

func TestCatalogUnitGroupUnitListWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Units.List(
		context.TODO(),
		"ug_andst6m79n41",
		augno.CatalogUnitGroupUnitListParams{
			Include: []string{"unit"},
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

func TestCatalogUnitGroupUnitDelete(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Units.Delete(
		context.TODO(),
		"un_82bd37dae5po",
		augno.CatalogUnitGroupUnitDeleteParams{
			UnitGroupID: "ug_andst6m79n41",
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
