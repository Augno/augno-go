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

func TestSaleSalesOrderLineNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.Lines.New(
		context.TODO(),
		"or_01d5034136c3ccc048abecc312",
		augno.SaleSalesOrderLineNewParams{
			CreateSalesOrderLineRequest: augno.CreateSalesOrderLineRequestParam{
				ProductID:  "pd_013c29ab3f1518d0004094c316",
				ProductSKU: "WIDGET-001",
				Quantity: augno.QuantityInputParam{
					UnitID: "un_01966263f74a5a0cae356000a1",
					Value:  "10",
				},
				ProductDescription: augno.String("product_description"),
				UnitPrice: augno.RateInputParam{
					DenominatorUnitID: "denominator_unit_id",
					NumeratorUnitID:   "numerator_unit_id",
					Value:             "value",
				},
			},
			Include: []string{"product"},
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

func TestSaleSalesOrderLineUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.Lines.Update(
		context.TODO(),
		"example",
		augno.SaleSalesOrderLineUpdateParams{
			ID:      "or_01d5034136c3ccc048abecc312",
			Include: []string{"product"},
			UpdateSalesOrderLineRequest: augno.UpdateSalesOrderLineRequestParam{
				ProductDescription: augno.String("product_description"),
				ProductSKU:         augno.String("product_sku"),
				Quantity: augno.QuantityInputParam{
					UnitID: "un_01966263f74a5a0cae356000a1",
					Value:  "20",
				},
				UnitCost: augno.RateInputParam{
					DenominatorUnitID: "denominator_unit_id",
					NumeratorUnitID:   "numerator_unit_id",
					Value:             "value",
				},
				UnitPrice: augno.RateInputParam{
					DenominatorUnitID: "un_01966263f74a5a0cae356000a1",
					NumeratorUnitID:   "un_01966263f74a5a0cae356000a1",
					Value:             "30.00",
				},
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

func TestSaleSalesOrderLineDelete(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.Lines.Delete(
		context.TODO(),
		"example",
		augno.SaleSalesOrderLineDeleteParams{
			ID: "or_01d5034136c3ccc048abecc312",
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
