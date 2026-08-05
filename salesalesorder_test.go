// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/augno/augno-go"
	"github.com/augno/augno-go/internal/testutil"
	"github.com/augno/augno-go/option"
)

func TestSaleSalesOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.New(context.TODO(), augno.SaleSalesOrderNewParams{
		CreateSalesOrderRequest: augno.CreateSalesOrderRequestParam{
			BillToAddressID: "ad_npqa5y43q26z",
			BuyerAccountID:  "ac_opnlh43ymyee",
			Lines: []augno.CreateSalesOrderLineInputParam{{
				ProductID: "pd_07oe0r7adh2w",
				Quantity: augno.QuantityInputParam{
					UnitID: "un_82bd37dae5po",
					Value:  "10",
				},
				ProductDescription: augno.String("product_description"),
				ProductSKU:         augno.String("product_sku"),
				UnitPrice: augno.RateInputParam{
					DenominatorUnitID: "denominator_unit_id",
					NumeratorUnitID:   "numerator_unit_id",
					Value:             "value",
				},
			}},
			PriorityCode:    "normal",
			ShipToAddressID: "ad_npqa5y43q26z",
			AcknowledgementEmailContacts: []augno.SalesOrderEmailContactInputParam{{
				AccountUserID: "acus_e5zu8bde0z3h",
			}},
			CarrierBillingAccountNumber: augno.String("123456789"),
			CarrierBillingType:          augno.CreateSalesOrderRequestCarrierBillingTypeSender,
			CarrierID:                   augno.String("cr_tv5vfjtgu1n3"),
			CustomerPurchaseOrderNumber: augno.String("PO-88231"),
			InvoiceEmailContacts: []augno.SalesOrderEmailContactInputParam{{
				AccountUserID: "acus_e5zu8bde0z3h",
			}},
			Note:            augno.String("Rush order for trade show"),
			OrderDiscountID: augno.String("ords_qnbrjvq5ih2q"),
			PaymentTermID:   augno.String("pytm_skssmsy21lem"),
			PromisedAt:      augno.Time(time.Now()),
			SalesRepID:      augno.String("acus_e5zu8bde0z3h"),
			ServiceLevelID:  augno.String("crop_4ilk9p6gccrx"),
			ShippingTermID:  augno.String("shtm_c5gxy05whw6r"),
		},
		Include: []string{"customer"},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleSalesOrderGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.Get(
		context.TODO(),
		"or_9lqo07quiwyb",
		augno.SaleSalesOrderGetParams{
			Include: []string{"customer"},
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

func TestSaleSalesOrderUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.Update(
		context.TODO(),
		"or_9lqo07quiwyb",
		augno.SaleSalesOrderUpdateParams{
			Include: []string{"customer"},
			UpdateSalesOrderRequest: augno.UpdateSalesOrderRequestParam{
				AcknowledgementEmailContacts: []augno.SalesOrderEmailContactInputParam{{
					AccountUserID: "account_user_id",
				}},
				AcknowledgmentStatus:        augno.UpdateSalesOrderRequestAcknowledgmentStatusNotSent,
				BillingAddressID:            augno.String("billing_address_id"),
				CarrierBillingAccountNumber: augno.String("carrier_billing_account_number"),
				CarrierBillingType:          augno.UpdateSalesOrderRequestCarrierBillingTypeSender,
				CarrierID:                   augno.String("cr_tv5vfjtgu1n3"),
				CustomerID:                  augno.String("customer_id"),
				CustomerPurchaseOrderNumber: augno.String("customer_purchase_order_number"),
				InvoiceEmailContacts: []augno.SalesOrderEmailContactInputParam{{
					AccountUserID: "account_user_id",
				}},
				Note:              augno.String("Updated shipping instructions"),
				OrderDiscountID:   augno.String("order_discount_id"),
				PaymentTermID:     augno.String("payment_term_id"),
				PriorityCode:      augno.String("normal"),
				PromisedAt:        augno.Time(time.Now()),
				SalesRepID:        augno.String("sales_rep_id"),
				ServiceLevelID:    augno.String("service_level_id"),
				ShippingAddressID: augno.String("ad_npqa5y43q26z"),
				ShippingTermID:    augno.String("shipping_term_id"),
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

func TestSaleSalesOrderListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.List(context.TODO(), augno.SaleSalesOrderListParams{
		Cursor:           augno.String("cursor"),
		CustomerGroupIDs: []string{"string"},
		CustomerIDs:      []string{"string"},
		EndsAt:           augno.String("ends_at"),
		Include:          []string{"customer"},
		ItemIDs:          []string{"string"},
		Limit:            augno.Int(0),
		ProductLineIDs:   []string{"string"},
		Q:                augno.String("q"),
		SalesRepIDs:      []string{"string"},
		StartsAt:         augno.String("starts_at"),
		StatusCodes:      []string{"string"},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleSalesOrderDelete(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.Delete(context.TODO(), "or_9lqo07quiwyb")
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleSalesOrderCheckout(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.Checkout(
		context.TODO(),
		"or_9lqo07quiwyb",
		augno.SaleSalesOrderCheckoutParams{
			CheckoutSalesOrderRequest: augno.CheckoutSalesOrderRequestParam{
				Email: "operations@acme.example.com",
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

func TestSaleSalesOrderPriceQuote(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.PriceQuote(context.TODO(), augno.SaleSalesOrderPriceQuoteParams{
		QuoteSalesOrderPricesRequest: augno.QuoteSalesOrderPricesRequestParam{
			BuyerAccountID: "ac_opnlh43ymyee",
			Lines: []augno.QuoteSalesOrderLineInputParam{{
				ProductID: "pd_07oe0r7adh2w",
				Quantity: augno.QuantityInputParam{
					UnitID: "un_82bd37dae5po",
					Value:  "10",
				},
			}},
		},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleSalesOrderGetStatusesWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.SalesOrders.GetStatuses(context.TODO(), augno.SaleSalesOrderGetStatusesParams{
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
