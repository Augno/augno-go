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
			BillToAddressID: "ad_012c2e4aeeb20f56c1a3d06cc7",
			BuyerAccountID:  "ac_0170df1ac58e4d24c66fc89f5f",
			Lines: []augno.CreateSalesOrderLineInputParam{{
				ProductID: "pd_013c29ab3f1518d0004094c316",
				Quantity: augno.QuantityInputParam{
					UnitID: "un_01966263f74a5a0cae356000a1",
					Value:  "10",
				},
				EdiLineItemID:      augno.String("edi_line_item_id"),
				ProductDescription: augno.String("product_description"),
				ProductSKU:         augno.String("product_sku"),
				UnitPrice: augno.RateInputParam{
					DenominatorUnitID: "denominator_unit_id",
					NumeratorUnitID:   "numerator_unit_id",
					Value:             "value",
				},
			}},
			PriorityCode:    "normal",
			ShipToAddressID: "ad_012c2e4aeeb20f56c1a3d06cc7",
			AcknowledgementEmailContacts: []augno.SalesOrderEmailContactInputParam{{
				AccountUserID: "account_user_id",
			}},
			CarrierBillingAccountNumber: augno.String("carrier_billing_account_number"),
			CarrierBillingType:          augno.CreateSalesOrderRequestCarrierBillingTypeSender,
			CarrierID:                   augno.String("cr_01784fd54c9ba197bb4e42f0e6"),
			CustomerPurchaseOrderNumber: augno.String("customer_purchase_order_number"),
			InvoiceEmailContacts: []augno.SalesOrderEmailContactInputParam{{
				AccountUserID: "account_user_id",
			}},
			Note:            augno.String("Rush order for trade show"),
			OrderDiscountID: augno.String("order_discount_id"),
			PaymentTermID:   augno.String("payment_term_id"),
			PromisedAt:      augno.Time(time.Now()),
			SalesRepID:      augno.String("sales_rep_id"),
			ServiceLevelID:  augno.String("crop_01cfaf03f104e90ef9680e2a30"),
			ShippingTermID:  augno.String("shipping_term_id"),
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
		Cursor:                augno.String("cursor"),
		CustomerGroupIDs:      []string{"string"},
		CustomerIDs:           []string{"string"},
		EndDate:               augno.String("end_date"),
		ExcludeInternalOrders: augno.Bool(true),
		Include:               []string{"customer"},
		ItemIDs:               []string{"string"},
		Limit:                 augno.Int(0),
		ProductLineIDs:        []string{"string"},
		Q:                     augno.String("q"),
		SalesRepIDs:           []string{"string"},
		StartDate:             augno.String("start_date"),
		StatusCodes:           []string{"string"},
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
