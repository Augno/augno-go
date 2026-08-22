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

func TestSaleAddressNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.Addresses.New(context.TODO(), openmrp.SaleAddressNewParams{
		AddressInput: openmrp.AddressInputParam{
			Country:           "US",
			Name:              "Headquarters",
			Email:             openmrp.String("warehouse@acme.com"),
			Locality:          openmrp.String("Springfield"),
			Phone:             openmrp.String("555-123-4567"),
			PostalCode:        openmrp.String("62701"),
			ReceiveCalendarID: openmrp.String("receive_calendar_id"),
			State:             openmrp.String("IL"),
			StreetLine1:       openmrp.String("123 Main St"),
			StreetLine2:       openmrp.String("Suite 400"),
			Type:              openmrp.AddressInputTypeStandard,
		},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleAddressGet(t *testing.T) {
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
	_, err := client.Sales.Addresses.Get(context.TODO(), "ad_j8cz0b79pwdb")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleAddressUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.Addresses.Update(
		context.TODO(),
		"ad_j8cz0b79pwdb",
		openmrp.SaleAddressUpdateParams{
			UpdateAddressRequest: openmrp.UpdateAddressRequestParam{
				Country:           openmrp.String("US"),
				Email:             openmrp.String("warehouse@acme.com"),
				Locality:          openmrp.String("Springfield"),
				Name:              openmrp.String("Warehouse"),
				Phone:             openmrp.String("555-123-4567"),
				PostalCode:        openmrp.String("62701"),
				ReceiveCalendarID: openmrp.String("receive_calendar_id"),
				State:             openmrp.String("IL"),
				StreetLine1:       openmrp.String("123 Main St"),
				StreetLine2:       openmrp.String("Suite 400"),
				Type:              openmrp.UpdateAddressRequestTypeStandard,
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

func TestSaleAddressListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.Addresses.List(context.TODO(), openmrp.SaleAddressListParams{
		Cursor: openmrp.String("cursor"),
		Limit:  openmrp.Int(0),
		Q:      openmrp.String("q"),
		Type:   openmrp.SaleAddressListParamsTypeStandard,
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleAddressDelete(t *testing.T) {
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
	_, err := client.Sales.Addresses.Delete(context.TODO(), "ad_j8cz0b79pwdb")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
