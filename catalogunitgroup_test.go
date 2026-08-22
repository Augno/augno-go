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

func TestCatalogUnitGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.New(context.TODO(), openmrp.CatalogUnitGroupNewParams{
		CreateUnitGroupRequest: openmrp.CreateUnitGroupRequestParam{
			BaseUnitID: "un_82bd37dae5po",
			Name:       "Weight Units",
			Type:       openmrp.CreateUnitGroupRequestTypeMass,
			AssociatedUnits: []openmrp.CreateUnitGroupUnitParam{{
				UnitID:                   "un_82bd37dae5po",
				CustomerPortalVisibility: openmrp.CreateUnitGroupUnitParamCustomerPortalVisibilityVisible,
				DiscountFixed:            openmrp.Float(0),
				DiscountPercentage:       openmrp.Float(1),
			}},
			Notes: openmrp.String("Used for raw-material weight tracking across the warehouse."),
		},
		Include: []string{"owner"},
	})
	if err != nil {
		var apierr *openmrp.Error
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
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Catalog.UnitGroups.Get(
		context.TODO(),
		"ug_andst6m79n41",
		openmrp.CatalogUnitGroupGetParams{
			Include: []string{"owner"},
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

func TestCatalogUnitGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.Update(
		context.TODO(),
		"ug_andst6m79n41",
		openmrp.CatalogUnitGroupUpdateParams{
			Include: []string{"owner"},
			UpdateUnitGroupRequest: openmrp.UpdateUnitGroupRequestParam{
				AssociatedUnits: []openmrp.CreateUnitGroupUnitParam{{
					UnitID:                   "un_82bd37dae5po",
					CustomerPortalVisibility: openmrp.CreateUnitGroupUnitParamCustomerPortalVisibilityVisible,
					DiscountFixed:            openmrp.Float(0),
					DiscountPercentage:       openmrp.Float(1),
				}},
				BaseUnitID: openmrp.String("un_82bd37dae5po"),
				Name:       openmrp.String("Weight Units (Updated)"),
				Notes:      openmrp.String("Added kilogram association for metric orders."),
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

func TestCatalogUnitGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.UnitGroups.List(context.TODO(), openmrp.CatalogUnitGroupListParams{
		Cursor:  openmrp.String("cursor"),
		Include: []string{"owner"},
		Limit:   openmrp.Int(0),
		Q:       openmrp.String("q"),
		Type:    openmrp.CatalogUnitGroupListParamsTypeCurrency,
	})
	if err != nil {
		var apierr *openmrp.Error
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
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Catalog.UnitGroups.Delete(context.TODO(), "ug_andst6m79n41")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
