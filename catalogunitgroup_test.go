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

func TestCatalogUnitGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.New(context.TODO(), augno.CatalogUnitGroupNewParams{
		CreateUnitGroupRequest: augno.CreateUnitGroupRequestParam{
			BaseUnitID: "un_01966263f74a5a0cae356000a1",
			Name:       "Weight Units",
			Type:       augno.CreateUnitGroupRequestTypeMass,
			AssociatedUnits: []augno.CreateUnitGroupUnitParam{{
				UnitID:                   "un_01966263f74a5a0cae356000a1",
				CustomerPortalVisibility: augno.CreateUnitGroupUnitParamCustomerPortalVisibilityVisible,
				DiscountFixed:            augno.Float(0),
				DiscountPercentage:       augno.Float(1),
			}},
			Notes: augno.String("notes"),
		},
		Include: []string{"owner"},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogUnitGroupGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Get(
		context.TODO(),
		"ug_01aad07abb8e41fd392d2d7013",
		augno.CatalogUnitGroupGetParams{
			Include: []string{"owner"},
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

func TestCatalogUnitGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Update(
		context.TODO(),
		"ug_01aad07abb8e41fd392d2d7013",
		augno.CatalogUnitGroupUpdateParams{
			Include: []string{"owner"},
			UpdateUnitGroupRequest: augno.UpdateUnitGroupRequestParam{
				AssociatedUnits: []augno.CreateUnitGroupUnitParam{{
					UnitID:                   "unit_id",
					CustomerPortalVisibility: augno.CreateUnitGroupUnitParamCustomerPortalVisibilityVisible,
					DiscountFixed:            augno.Float(0),
					DiscountPercentage:       augno.Float(0),
				}},
				BaseUnitID: augno.String("un_01966263f74a5a0cae356000a1"),
				Name:       augno.String("Weight Units (Updated)"),
				Notes:      augno.String("notes"),
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

func TestCatalogUnitGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.List(context.TODO(), augno.CatalogUnitGroupListParams{
		Cursor:  augno.String("cursor"),
		Include: []string{"owner"},
		Limit:   augno.Int(0),
		Q:       augno.String("q"),
		Type:    augno.CatalogUnitGroupListParamsTypeCurrency,
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogUnitGroupDelete(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Delete(context.TODO(), "ug_01aad07abb8e41fd392d2d7013")
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
