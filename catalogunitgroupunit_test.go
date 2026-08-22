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

func TestCatalogUnitGroupUnitNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Units.New(
		context.TODO(),
		"ug_andst6m79n41",
		openmrp.CatalogUnitGroupUnitNewParams{
			CreateUnitGroupUnitRequest: openmrp.CreateUnitGroupUnitRequestParam{
				UnitID:                   "un_82bd37dae5po",
				CustomerPortalVisibility: openmrp.CreateUnitGroupUnitRequestCustomerPortalVisibilityVisible,
				DiscountFixed:            openmrp.Float(0),
				DiscountPercentage:       openmrp.Float(1),
			},
			Include: []string{"unit"},
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

func TestCatalogUnitGroupUnitGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Units.Get(
		context.TODO(),
		"un_82bd37dae5po",
		openmrp.CatalogUnitGroupUnitGetParams{
			UnitGroupID: "ug_andst6m79n41",
			Include:     []string{"unit"},
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

func TestCatalogUnitGroupUnitUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Units.Update(
		context.TODO(),
		"un_82bd37dae5po",
		openmrp.CatalogUnitGroupUnitUpdateParams{
			UnitGroupID: "ug_andst6m79n41",
			Include:     []string{"unit"},
			UpdateUnitGroupUnitRequest: openmrp.UpdateUnitGroupUnitRequestParam{
				CustomerPortalVisibility: openmrp.UpdateUnitGroupUnitRequestCustomerPortalVisibilityVisible,
				DiscountFixed:            openmrp.Float(2.5),
				DiscountPercentage:       openmrp.Float(0.9),
				UnitID:                   openmrp.String("un_82bd37dae5po"),
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

func TestCatalogUnitGroupUnitListWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Units.List(
		context.TODO(),
		"ug_andst6m79n41",
		openmrp.CatalogUnitGroupUnitListParams{
			Include: []string{"unit"},
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

func TestCatalogUnitGroupUnitDelete(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Units.Delete(
		context.TODO(),
		"un_82bd37dae5po",
		openmrp.CatalogUnitGroupUnitDeleteParams{
			UnitGroupID: "ug_andst6m79n41",
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
