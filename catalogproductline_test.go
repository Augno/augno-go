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

func TestCatalogProductLineNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.ProductLines.New(context.TODO(), openmrp.CatalogProductLineNewParams{
		CreateProductLineRequest: openmrp.CreateProductLineRequestParam{
			CommissionPolicy: openmrp.CreateProductLineRequestCommissionPolicyCommissionExempt,
			FreightPolicy:    openmrp.CreateProductLineRequestFreightPolicyBilledFreight,
			Name:             "Industrial Fasteners",
			UnitGroupID:      "ug_andst6m79n41",
			DefaultLot: openmrp.QuantityInputParam{
				UnitID: "unit_id",
				Value:  "value",
			},
			FulfillmentPolicy: openmrp.CreateProductLineRequestFulfillmentPolicyMakeToStock,
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

func TestCatalogProductLineGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.ProductLines.Get(
		context.TODO(),
		"pdln_k9bnlgvxhxjh",
		openmrp.CatalogProductLineGetParams{
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

func TestCatalogProductLineUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.ProductLines.Update(
		context.TODO(),
		"pdln_k9bnlgvxhxjh",
		openmrp.CatalogProductLineUpdateParams{
			Include: []string{"owner"},
			UpdateProductLineRequest: openmrp.UpdateProductLineRequestParam{
				CommissionPolicy: openmrp.UpdateProductLineRequestCommissionPolicyCommissionApplied,
				DefaultLot: openmrp.QuantityInputParam{
					UnitID: "unit_id",
					Value:  "value",
				},
				FreightPolicy:     openmrp.UpdateProductLineRequestFreightPolicyBilledFreight,
				FulfillmentPolicy: openmrp.UpdateProductLineRequestFulfillmentPolicyMakeToStock,
				Name:              openmrp.String("Updated Product Line"),
				UnitGroupID:       openmrp.String("ug_andst6m79n41"),
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

func TestCatalogProductLineListWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.ProductLines.List(context.TODO(), openmrp.CatalogProductLineListParams{
		Cursor:  openmrp.String("cursor"),
		Include: []string{"owner"},
		Limit:   openmrp.Int(0),
		Q:       openmrp.String("q"),
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogProductLineDelete(t *testing.T) {
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
	_, err := client.Catalog.ProductLines.Delete(context.TODO(), "pdln_k9bnlgvxhxjh")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
