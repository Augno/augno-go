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

func TestSaleVolumeDiscountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.VolumeDiscounts.New(context.TODO(), augno.SaleVolumeDiscountNewParams{
		CreateVolumeDiscountRequest: augno.CreateVolumeDiscountRequestParam{
			Name: "Bulk Order Discount",
			Tiers: []augno.CreateVolumeDiscountTierInputParam{{
				DiscountPercentage: "5.000000000000000000000000000000",
				Name:               "100+ Units",
				Threshold:          "100.000000000000000000000000000000",
				ParentTierID:       augno.String("parent_tier_id"),
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
		var apierr *augno.Error
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
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Sales.VolumeDiscounts.Get(
		context.TODO(),
		"quds_bn7hto9s10pp",
		augno.SaleVolumeDiscountGetParams{
			Include: []string{"customer_groups"},
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

func TestSaleVolumeDiscountUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.VolumeDiscounts.Update(
		context.TODO(),
		"quds_bn7hto9s10pp",
		augno.SaleVolumeDiscountUpdateParams{
			UpdateVolumeDiscountRequest: augno.UpdateVolumeDiscountRequestParam{
				HasAttributes:     true,
				HasCategories:     true,
				HasCustomerGroups: true,
				HasProductLines:   true,
				HasTiers:          true,
				HasUnits:          true,
				AttributeIDs:      []string{"string"},
				CategoryIDs:       []string{"string"},
				CustomerGroupIDs:  []string{"string"},
				Name:              augno.String("Updated Bulk Discount"),
				ProductLineIDs:    []string{"string"},
				Tiers: []augno.UpdateVolumeDiscountTierInputParam{{
					ID:                 augno.String("id"),
					DiscountPercentage: augno.String("10.000000000000000000000000000000"),
					Name:               augno.String("50+ Units"),
					ParentTierID:       augno.String("parent_tier_id"),
					Threshold:          augno.String("50.000000000000000000000000000000"),
				}},
				UnitIDs: []string{"string"},
			},
			Include: []string{"customer_groups"},
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

func TestSaleVolumeDiscountListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.VolumeDiscounts.List(context.TODO(), augno.SaleVolumeDiscountListParams{
		Cursor:  augno.String("cursor"),
		Include: []string{"customer_groups"},
		Limit:   augno.Int(0),
		Q:       augno.String("q"),
	})
	if err != nil {
		var apierr *augno.Error
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
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Sales.VolumeDiscounts.Delete(context.TODO(), "quds_bn7hto9s10pp")
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
