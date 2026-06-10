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
		"ug_01aad07abb8e41fd392d2d7013",
		augno.CatalogUnitGroupUnitNewParams{
			CreateUnitGroupUnitRequest: augno.CreateUnitGroupUnitRequestParam{
				UnitID:                   "un_01966263f74a5a0cae356000a1",
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
		"un_01966263f74a5a0cae356000a1",
		augno.CatalogUnitGroupUnitGetParams{
			UnitGroupID: "ug_01aad07abb8e41fd392d2d7013",
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
		"un_01966263f74a5a0cae356000a1",
		augno.CatalogUnitGroupUnitUpdateParams{
			UnitGroupID: "ug_01aad07abb8e41fd392d2d7013",
			Include:     []string{"unit"},
			UpdateUnitGroupUnitRequest: augno.UpdateUnitGroupUnitRequestParam{
				CustomerPortalVisibility: augno.UpdateUnitGroupUnitRequestCustomerPortalVisibilityVisible,
				DiscountFixed:            augno.Float(0),
				DiscountPercentage:       augno.Float(0.9),
				UnitID:                   augno.String("un_01966263f74a5a0cae356000a1"),
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
		"ug_01aad07abb8e41fd392d2d7013",
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
		"un_01966263f74a5a0cae356000a1",
		augno.CatalogUnitGroupUnitDeleteParams{
			UnitGroupID: "ug_01aad07abb8e41fd392d2d7013",
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
