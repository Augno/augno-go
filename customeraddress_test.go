// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/augno-go"
	"github.com/stainless-sdks/augno-go/internal/testutil"
	"github.com/stainless-sdks/augno-go/option"
)

func TestCustomerAddressNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Customers.Addresses.New(
		context.TODO(),
		"customer_id",
		augno.CustomerAddressNewParams{
			AddressLine1: "123 Main St",
			City:         "West Lafayette",
			Country:      "USA",
			Name:         "John Doe",
			PostalCode:   "47906",
			AddressLine2: augno.String("Apt 1"),
			Email:        augno.String("john.doe@example.com"),
			Include:      []any{"full"},
			IsDropShip:   augno.Bool(false),
			Phone:        augno.String("+15551234567"),
			State:        augno.String("IN"),
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

func TestCustomerAddressGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Customers.Addresses.Get(
		context.TODO(),
		"address_id",
		augno.CustomerAddressGetParams{
			CustomerID: "customer_id",
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

func TestCustomerAddressUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Customers.Addresses.Update(
		context.TODO(),
		"address_id",
		augno.CustomerAddressUpdateParams{
			CustomerID:   "customer_id",
			AddressLine1: augno.String("123 Main St"),
			AddressLine2: augno.String("Apt 1"),
			City:         augno.String("West Lafayette"),
			Country:      augno.String("USA"),
			Email:        augno.String("john.doe@example.com"),
			Include:      []any{"full"},
			IsDropShip:   augno.Bool(false),
			Name:         augno.String("John Doe"),
			Phone:        augno.String("+15551234567"),
			PostalCode:   augno.String("47906"),
			State:        augno.String("IN"),
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

func TestCustomerAddressList(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Customers.Addresses.List(context.TODO(), "customer_id")
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerAddressDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Customers.Addresses.Delete(
		context.TODO(),
		"address_id",
		augno.CustomerAddressDeleteParams{
			CustomerID: "customer_id",
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
