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

func TestSaleVolumeDiscountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.VolumeDiscounts.New(context.TODO(), openmrp.SaleVolumeDiscountNewParams{
		CreateVolumeDiscountRequest: openmrp.CreateVolumeDiscountRequestParam{
			Name: "Bulk Order Discount",
			Tiers: []openmrp.CreateVolumeDiscountTierInputParam{{
				DiscountPercentage: "5.000000000000000000000000000000",
				Name:               "100+ Units",
				Threshold:          "100.000000000000000000000000000000",
				ParentTierID:       openmrp.String("parent_tier_id"),
			}},
			AttributeIDs:     []string{"string"},
			CategoryIDs:      []string{"string"},
			CustomerGroupIDs: []string{"string"},
			ProductLineIDs:   []string{"string"},
			UnitIDs:          []string{"string"},
		},
		Include: []string{"customer_groups"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleVolumeDiscountGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.VolumeDiscounts.Get(
		context.TODO(),
		"quds_bn7hto9s10pp",
		openmrp.SaleVolumeDiscountGetParams{
			Include: []string{"customer_groups"},
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

func TestSaleVolumeDiscountUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.VolumeDiscounts.Update(
		context.TODO(),
		"quds_bn7hto9s10pp",
		openmrp.SaleVolumeDiscountUpdateParams{
			UpdateVolumeDiscountRequest: openmrp.UpdateVolumeDiscountRequestParam{
				HasAttributes:     true,
				HasCategories:     true,
				HasCustomerGroups: true,
				HasProductLines:   true,
				HasTiers:          true,
				HasUnits:          true,
				AttributeIDs:      []string{"string"},
				CategoryIDs:       []string{"string"},
				CustomerGroupIDs:  []string{"string"},
				Name:              openmrp.String("Updated Bulk Discount"),
				ProductLineIDs:    []string{"string"},
				Tiers: []openmrp.UpdateVolumeDiscountTierInputParam{{
					ID:                 openmrp.String("id"),
					DiscountPercentage: openmrp.String("10.000000000000000000000000000000"),
					Name:               openmrp.String("50+ Units"),
					ParentTierID:       openmrp.String("parent_tier_id"),
					Threshold:          openmrp.String("50.000000000000000000000000000000"),
				}},
				UnitIDs: []string{"string"},
			},
			Include: []string{"customer_groups"},
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

func TestSaleVolumeDiscountListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.VolumeDiscounts.List(context.TODO(), openmrp.SaleVolumeDiscountListParams{
		Cursor:  openmrp.String("cursor"),
		Include: []string{"customer_groups"},
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

func TestSaleVolumeDiscountDelete(t *testing.T) {
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
	_, err := client.Sales.VolumeDiscounts.Delete(context.TODO(), "quds_bn7hto9s10pp")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
