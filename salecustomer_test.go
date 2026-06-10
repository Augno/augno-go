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

func TestSaleCustomerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.Customers.New(context.TODO(), augno.SaleCustomerNewParams{
		CreateCustomerRequest: augno.CreateCustomerRequestParam{
			BillToAddress: augno.AddressInputParam{
				Country:     "US",
				Name:        "Acme Inc.",
				Email:       augno.String("email"),
				Locality:    augno.String("New York"),
				Phone:       augno.String("phone"),
				PostalCode:  augno.String("10001"),
				State:       augno.String("NY"),
				StreetLine1: augno.String("123 Main St"),
				StreetLine2: augno.String("street_line_2"),
				Type:        augno.AddressInputTypeStandard,
			},
			CustomerTypeGroupID:   "acgp_018e88072d1320808dc979cfac",
			DefaultCarrierID:      "cr_01784fd54c9ba197bb4e42f0e6",
			DefaultPaymentTermID:  "pytm_018694d6601ea771cd1b52e890",
			DefaultShippingTermID: "shtm_014341ab4bb5bf94d5b6936f86",
			Name:                  "Acme Inc.",
			ShipToAddress: augno.AddressInputParam{
				Country:     "US",
				Name:        "Acme Inc.",
				Email:       augno.String("email"),
				Locality:    augno.String("New York"),
				Phone:       augno.String("phone"),
				PostalCode:  augno.String("10001"),
				State:       augno.String("NY"),
				StreetLine1: augno.String("123 Main St"),
				StreetLine2: augno.String("street_line_2"),
				Type:        augno.AddressInputTypeStandard,
			},
			CarrierBillingAccount: augno.String("carrier_billing_account"),
			CarrierBillingType:    augno.CreateCustomerRequestCarrierBillingTypeSender,
			CommissionPolicy:      augno.CreateCustomerRequestCommissionPolicyCommissionApplied,
			CreditLimit: augno.QuantityInputParam{
				UnitID: "unit_id",
				Value:  "value",
			},
			CustomerPriceGroupIDs: []string{"string"},
			DefaultPriority:       augno.CreateCustomerRequestDefaultPriorityLow,
			DefaultSalesRepID:     augno.String("default_sales_rep_id"),
			DefaultServiceLevelID: augno.String("default_service_level_id"),
			EdiStatus:             augno.CreateCustomerRequestEdiStatusEnabled,
			Email:                 augno.String("email"),
			FreightPolicy:         augno.CreateCustomerRequestFreightPolicyFreeFreight,
			Note:                  augno.String("Key enterprise account"),
			Number:                augno.String("number"),
			Phone:                 augno.String("phone"),
			Status:                augno.CreateCustomerRequestStatusNormal,
			URL:                   augno.String("url"),
		},
		Include: []string{"bill_to_address"},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleCustomerGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.Customers.Get(
		context.TODO(),
		"ac_0170df1ac58e4d24c66fc89f5f",
		augno.SaleCustomerGetParams{
			Include: []string{"bill_to_address"},
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

func TestSaleCustomerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.Customers.Update(
		context.TODO(),
		"ac_0170df1ac58e4d24c66fc89f5f",
		augno.SaleCustomerUpdateParams{
			Include: []string{"bill_to_address"},
			UpdateCustomerRequest: augno.UpdateCustomerRequestParam{
				BillToAddressID:       augno.String("bill_to_address_id"),
				CarrierBillingAccount: augno.String("carrier_billing_account"),
				CarrierBillingType:    augno.UpdateCustomerRequestCarrierBillingTypeSender,
				CommissionPolicy:      augno.UpdateCustomerRequestCommissionPolicyCommissionApplied,
				CreditLimit: augno.QuantityInputParam{
					UnitID: "unit_id",
					Value:  "value",
				},
				CustomerPriceGroupIDs: []string{"string"},
				CustomerTypeGroupID:   augno.String("customer_type_group_id"),
				DefaultCarrierID:      augno.String("cr_01784fd54c9ba197bb4e42f0e6"),
				DefaultPaymentTermID:  augno.String("default_payment_term_id"),
				DefaultPriority:       augno.UpdateCustomerRequestDefaultPriorityLow,
				DefaultSalesRepID:     augno.String("default_sales_rep_id"),
				DefaultServiceLevelID: augno.String("default_service_level_id"),
				DefaultShippingTermID: augno.String("default_shipping_term_id"),
				EdiStatus:             augno.UpdateCustomerRequestEdiStatusEnabled,
				Email:                 augno.String("email"),
				FreightPolicy:         augno.UpdateCustomerRequestFreightPolicyBilledFreight,
				Name:                  augno.String("Acme Corp Updated"),
				Note:                  augno.String("Updated account notes"),
				Number:                augno.String("number"),
				Phone:                 augno.String("phone"),
				ShipToAddressID:       augno.String("ship_to_address_id"),
				Status:                augno.UpdateCustomerRequestStatusNormal,
				URL:                   augno.String("url"),
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

func TestSaleCustomerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sales.Customers.List(context.TODO(), augno.SaleCustomerListParams{
		CarrierIDs:            []string{"string"},
		City:                  augno.String("city"),
		CommissionStatusCodes: []string{"commission_applied"},
		Cursor:                augno.String("cursor"),
		CustomerGroupIDs:      []string{"string"},
		EndDate:               augno.Time(time.Now()),
		FreightStatusCodes:    []string{"free_freight"},
		Include:               []string{"bill_to_address"},
		Limit:                 augno.Int(0),
		ParentAccountStatus:   augno.SaleCustomerListParamsParentAccountStatusParent,
		PaymentTermIDs:        []string{"string"},
		PostalCode:            augno.String("postal_code"),
		PricingGroupIDs:       []string{"string"},
		Q:                     augno.String("q"),
		SalesRepIDs:           []string{"string"},
		ServiceLevelIDs:       []string{"string"},
		ShippingTermIDs:       []string{"string"},
		StartDate:             augno.Time(time.Now()),
		State:                 augno.String("state"),
		StatusCodes:           []string{"normal"},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleCustomerDelete(t *testing.T) {
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
	_, err := client.Sales.Customers.Delete(context.TODO(), "ac_0170df1ac58e4d24c66fc89f5f")
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
