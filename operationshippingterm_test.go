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

func TestOperationShippingTermNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ShippingTerms.New(context.TODO(), openmrp.OperationShippingTermNewParams{
		CreateShippingTermRequest: openmrp.CreateShippingTermRequestParam{
			Name: "Prepaid",
			Type: openmrp.CreateShippingTermRequestTypeFlatRateFreight,
			FlatRate: openmrp.QuantityInputParam{
				UnitID: "un_82bd37dae5po",
				Value:  "15.00",
			},
			FreeShippingServiceLevelIDs: []string{"crop_4ilk9p6gccrx"},
			MinimumOrderValue: openmrp.QuantityInputParam{
				UnitID: "un_82bd37dae5po",
				Value:  "500.00",
			},
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

func TestOperationShippingTermGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ShippingTerms.Get(
		context.TODO(),
		"shtm_c5gxy05whw6r",
		openmrp.OperationShippingTermGetParams{
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

func TestOperationShippingTermUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ShippingTerms.Update(
		context.TODO(),
		"shtm_c5gxy05whw6r",
		openmrp.OperationShippingTermUpdateParams{
			Include: []string{"owner"},
			UpdateShippingTermRequest: openmrp.UpdateShippingTermRequestParam{
				FlatRate: openmrp.QuantityInputParam{
					UnitID: "un_82bd37dae5po",
					Value:  "15.00",
				},
				FreeShippingServiceLevelIDs: []string{"crop_4ilk9p6gccrx"},
				MinimumOrderValue: openmrp.QuantityInputParam{
					UnitID: "un_82bd37dae5po",
					Value:  "500.00",
				},
				Name: openmrp.String("Collect"),
				Type: openmrp.UpdateShippingTermRequestTypeFlatRateFreight,
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

func TestOperationShippingTermListWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ShippingTerms.List(context.TODO(), openmrp.OperationShippingTermListParams{
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

func TestOperationShippingTermDelete(t *testing.T) {
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
	_, err := client.Operations.ShippingTerms.Delete(context.TODO(), "shtm_c5gxy05whw6r")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
