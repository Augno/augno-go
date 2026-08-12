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

func TestOperationCarrierServiceLevelNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.Carriers.ServiceLevels.New(
		context.TODO(),
		"cr_tv5vfjtgu1n3",
		augno.OperationCarrierServiceLevelNewParams{
			CreateServiceLevelRequest: augno.CreateServiceLevelRequestParam{
				Code:                     "ground",
				IsDefault:                false,
				Name:                     "Ground Shipping",
				CustomerPortalVisibility: augno.CreateServiceLevelRequestCustomerPortalVisibilityVisible,
				DefaultTransitDays:       augno.Int(3),
			},
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

func TestOperationCarrierServiceLevelGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.Carriers.ServiceLevels.Get(
		context.TODO(),
		"crop_4ilk9p6gccrx",
		augno.OperationCarrierServiceLevelGetParams{
			CarrierID: "cr_tv5vfjtgu1n3",
			Include:   []string{"owner"},
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

func TestOperationCarrierServiceLevelUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.Carriers.ServiceLevels.Update(
		context.TODO(),
		"crop_4ilk9p6gccrx",
		augno.OperationCarrierServiceLevelUpdateParams{
			CarrierID: "cr_tv5vfjtgu1n3",
			Include:   []string{"owner"},
			UpdateServiceLevelRequest: augno.UpdateServiceLevelRequestParam{
				Code:                     augno.String("express"),
				CustomerPortalVisibility: augno.UpdateServiceLevelRequestCustomerPortalVisibilityVisible,
				DefaultTransitDays:       augno.Int(0),
				IsDefault:                augno.Bool(false),
				Name:                     augno.String("Express Shipping"),
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

func TestOperationCarrierServiceLevelListWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.Carriers.ServiceLevels.List(
		context.TODO(),
		"cr_tv5vfjtgu1n3",
		augno.OperationCarrierServiceLevelListParams{
			Cursor:  augno.String("cursor"),
			Include: []string{"owner"},
			Limit:   augno.Int(0),
			Q:       augno.String("q"),
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

func TestOperationCarrierServiceLevelDelete(t *testing.T) {
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
	_, err := client.Operations.Carriers.ServiceLevels.Delete(
		context.TODO(),
		"crop_4ilk9p6gccrx",
		augno.OperationCarrierServiceLevelDeleteParams{
			CarrierID: "cr_tv5vfjtgu1n3",
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
