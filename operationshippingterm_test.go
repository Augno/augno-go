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

func TestOperationShippingTermNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ShippingTerms.New(context.TODO(), augno.OperationShippingTermNewParams{
		CreateShippingTermRequest: augno.CreateShippingTermRequestParam{
			Name: "Prepaid",
			Type: augno.CreateShippingTermRequestTypeFlatRateFreight,
			FlatRate: augno.QuantityInputParam{
				UnitID: "un_01966263f74a5a0cae356000a1",
				Value:  "15.00",
			},
			FreeShippingServiceLevelIDs: []string{"crop_01cfaf03f104e90ef9680e2a30"},
			MinimumOrderValue: augno.QuantityInputParam{
				UnitID: "un_01966263f74a5a0cae356000a1",
				Value:  "500.00",
			},
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

func TestOperationShippingTermGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ShippingTerms.Get(
		context.TODO(),
		"shtm_014341ab4bb5bf94d5b6936f86",
		augno.OperationShippingTermGetParams{
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

func TestOperationShippingTermUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ShippingTerms.Update(
		context.TODO(),
		"shtm_014341ab4bb5bf94d5b6936f86",
		augno.OperationShippingTermUpdateParams{
			Include: []string{"owner"},
			UpdateShippingTermRequest: augno.UpdateShippingTermRequestParam{
				FlatRate: augno.QuantityInputParam{
					UnitID: "un_01966263f74a5a0cae356000a1",
					Value:  "15.00",
				},
				FreeShippingServiceLevelIDs: []string{"crop_01cfaf03f104e90ef9680e2a30"},
				MinimumOrderValue: augno.QuantityInputParam{
					UnitID: "un_01966263f74a5a0cae356000a1",
					Value:  "500.00",
				},
				Name: augno.String("Collect"),
				Type: augno.UpdateShippingTermRequestTypeFlatRateFreight,
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

func TestOperationShippingTermListWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ShippingTerms.List(context.TODO(), augno.OperationShippingTermListParams{
		Cursor:  augno.String("cursor"),
		Include: []string{"owner"},
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

func TestOperationShippingTermDelete(t *testing.T) {
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
	_, err := client.Operations.ShippingTerms.Delete(context.TODO(), "shtm_014341ab4bb5bf94d5b6936f86")
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
