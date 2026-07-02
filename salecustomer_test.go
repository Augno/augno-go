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
				Email:       augno.String("warehouse@acme.com"),
				Locality:    augno.String("New York"),
				Phone:       augno.String("555-123-4567"),
				PostalCode:  augno.String("10001"),
				State:       augno.String("NY"),
				StreetLine1: augno.String("123 Main St"),
				StreetLine2: augno.String("Suite 400"),
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
				Email:       augno.String("warehouse@acme.com"),
				Locality:    augno.String("New York"),
				Phone:       augno.String("555-123-4567"),
				PostalCode:  augno.String("10001"),
				State:       augno.String("NY"),
				StreetLine1: augno.String("123 Main St"),
				StreetLine2: augno.String("Suite 400"),
				Type:        augno.AddressInputTypeStandard,
			},
			CarrierBillingAccount: augno.String("123456789"),
			CarrierBillingType:    augno.CreateCustomerRequestCarrierBillingTypeSender,
			CommissionPolicy:      augno.CreateCustomerRequestCommissionPolicyCommissionApplied,
			CreditLimit: augno.QuantityInputParam{
				UnitID: "un_01966263f74a5a0cae356000a1",
				Value:  "10000.00",
			},
			CustomerPriceGroupIDs: []string{"acgp_018e88072d1320808dc979cfac"},
			DefaultPriority:       augno.CreateCustomerRequestDefaultPriorityNormal,
			DefaultSalesRepID:     augno.String("acus_01ea9983ddb41dacc44ecf997c"),
			DefaultServiceLevelID: augno.String("crop_01cfaf03f104e90ef9680e2a30"),
			EdiStatus:             augno.CreateCustomerRequestEdiStatusDisabled,
			Email:                 augno.String("orders@acme.com"),
			FreightPolicy:         augno.CreateCustomerRequestFreightPolicyBilledFreight,
			Note:                  augno.String("Key enterprise account"),
			Number:                augno.String("100042"),
			Phone:                 augno.String("555-123-4567"),
			Status:                augno.CreateCustomerRequestStatusNormal,
			URL:                   augno.String("https://acme.com"),
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
				BillToAddressID:       augno.String("ad_012c2e4aeeb20f56c1a3d06cc7"),
				CarrierBillingAccount: augno.String("123456789"),
				CarrierBillingType:    augno.UpdateCustomerRequestCarrierBillingTypeSender,
				CommissionPolicy:      augno.UpdateCustomerRequestCommissionPolicyCommissionApplied,
				CreditLimit: augno.QuantityInputParam{
					UnitID: "un_01966263f74a5a0cae356000a1",
					Value:  "10000.00",
				},
				CustomerPriceGroupIDs: []string{"acgp_018e88072d1320808dc979cfac"},
				CustomerTypeGroupID:   augno.String("acgp_018e88072d1320808dc979cfac"),
				DefaultCarrierID:      augno.String("cr_01784fd54c9ba197bb4e42f0e6"),
				DefaultPaymentTermID:  augno.String("pytm_018694d6601ea771cd1b52e890"),
				DefaultPriority:       augno.UpdateCustomerRequestDefaultPriorityNormal,
				DefaultSalesRepID:     augno.String("acus_01ea9983ddb41dacc44ecf997c"),
				DefaultServiceLevelID: augno.String("crop_01cfaf03f104e90ef9680e2a30"),
				DefaultShippingTermID: augno.String("shtm_014341ab4bb5bf94d5b6936f86"),
				EdiStatus:             augno.UpdateCustomerRequestEdiStatusDisabled,
				Email:                 augno.String("orders@acme.com"),
				FreightPolicy:         augno.UpdateCustomerRequestFreightPolicyBilledFreight,
				Name:                  augno.String("Acme Corp Updated"),
				Note:                  augno.String("Updated account notes"),
				Number:                augno.String("100042"),
				Phone:                 augno.String("555-123-4567"),
				ShipToAddressID:       augno.String("ad_012c2e4aeeb20f56c1a3d06cc7"),
				Status:                augno.UpdateCustomerRequestStatusNormal,
				URL:                   augno.String("https://acme.com"),
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
