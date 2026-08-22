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

func TestOperationCarrierServiceLevelNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.Carriers.ServiceLevels.New(
		context.TODO(),
		"cr_tv5vfjtgu1n3",
		openmrp.OperationCarrierServiceLevelNewParams{
			CreateServiceLevelRequest: openmrp.CreateServiceLevelRequestParam{
				Code:                     "ground",
				IsDefault:                false,
				Name:                     "Ground Shipping",
				CustomerPortalVisibility: openmrp.CreateServiceLevelRequestCustomerPortalVisibilityVisible,
				DefaultTransitDays:       openmrp.Int(3),
			},
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

func TestOperationCarrierServiceLevelGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.Carriers.ServiceLevels.Get(
		context.TODO(),
		"crop_4ilk9p6gccrx",
		openmrp.OperationCarrierServiceLevelGetParams{
			CarrierID: "cr_tv5vfjtgu1n3",
			Include:   []string{"owner"},
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

func TestOperationCarrierServiceLevelUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.Carriers.ServiceLevels.Update(
		context.TODO(),
		"crop_4ilk9p6gccrx",
		openmrp.OperationCarrierServiceLevelUpdateParams{
			CarrierID: "cr_tv5vfjtgu1n3",
			Include:   []string{"owner"},
			UpdateServiceLevelRequest: openmrp.UpdateServiceLevelRequestParam{
				Code:                     openmrp.String("express"),
				CustomerPortalVisibility: openmrp.UpdateServiceLevelRequestCustomerPortalVisibilityVisible,
				DefaultTransitDays:       openmrp.Int(0),
				IsDefault:                openmrp.Bool(false),
				Name:                     openmrp.String("Express Shipping"),
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

func TestOperationCarrierServiceLevelListWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.Carriers.ServiceLevels.List(
		context.TODO(),
		"cr_tv5vfjtgu1n3",
		openmrp.OperationCarrierServiceLevelListParams{
			Cursor:  openmrp.String("cursor"),
			Include: []string{"owner"},
			Limit:   openmrp.Int(0),
			Q:       openmrp.String("q"),
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

func TestOperationCarrierServiceLevelDelete(t *testing.T) {
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
	_, err := client.Operations.Carriers.ServiceLevels.Delete(
		context.TODO(),
		"crop_4ilk9p6gccrx",
		openmrp.OperationCarrierServiceLevelDeleteParams{
			CarrierID: "cr_tv5vfjtgu1n3",
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
