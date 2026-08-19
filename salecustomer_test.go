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
				Country:           "US",
				Name:              "Acme Inc.",
				Email:             augno.String("warehouse@acme.com"),
				Locality:          augno.String("New York"),
				Phone:             augno.String("555-123-4567"),
				PostalCode:        augno.String("10001"),
				ReceiveCalendarID: augno.String("receive_calendar_id"),
				State:             augno.String("NY"),
				StreetLine1:       augno.String("123 Main St"),
				StreetLine2:       augno.String("Suite 400"),
				Type:              augno.AddressInputTypeStandard,
			},
			CustomerTypeGroupID:   "acgp_6p4z57e9alaf",
			DefaultCarrierID:      "cr_tv5vfjtgu1n3",
			DefaultPaymentTermID:  "pytm_skssmsy21lem",
			DefaultShippingTermID: "shtm_c5gxy05whw6r",
			Name:                  "Acme Inc.",
			ShipToAddress: augno.AddressInputParam{
				Country:           "US",
				Name:              "Acme Inc.",
				Email:             augno.String("warehouse@acme.com"),
				Locality:          augno.String("New York"),
				Phone:             augno.String("555-123-4567"),
				PostalCode:        augno.String("10001"),
				ReceiveCalendarID: augno.String("receive_calendar_id"),
				State:             augno.String("NY"),
				StreetLine1:       augno.String("123 Main St"),
				StreetLine2:       augno.String("Suite 400"),
				Type:              augno.AddressInputTypeStandard,
			},
			CarrierBillingAccount: augno.String("123456789"),
			CarrierBillingType:    augno.CreateCustomerRequestCarrierBillingTypeSender,
			CommissionPolicy:      augno.CreateCustomerRequestCommissionPolicyCommissionApplied,
			CreditLimit: augno.QuantityInputParam{
				UnitID: "un_82bd37dae5po",
				Value:  "10000.00",
			},
			CustomerPriceGroupIDs: []string{"acgp_6p4z57e9alaf"},
			DefaultPriority:       augno.CreateCustomerRequestDefaultPriorityNormal,
			DefaultSalesRepID:     augno.String("acus_e5zu8bde0z3h"),
			DefaultServiceLevelID: augno.String("crop_4ilk9p6gccrx"),
			EdiStatus:             augno.CreateCustomerRequestEdiStatusDisabled,
			Email:                 augno.String("orders@acme.com"),
			FreightPolicy:         augno.CreateCustomerRequestFreightPolicyBilledFreight,
			LeadTimeDays:          augno.Int(0),
			Note:                  augno.String("Key enterprise account"),
			Number:                augno.String("100042"),
			Phone:                 augno.String("555-123-4567"),
			ReceiveCalendarID:     augno.String("receive_calendar_id"),
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
		"ac_opnlh43ymyee",
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
		"ac_opnlh43ymyee",
		augno.SaleCustomerUpdateParams{
			Include: []string{"bill_to_address"},
			UpdateCustomerRequest: augno.UpdateCustomerRequestParam{
				BillToAddressID:       augno.String("ad_npqa5y43q26z"),
				CarrierBillingAccount: augno.String("123456789"),
				CarrierBillingType:    augno.UpdateCustomerRequestCarrierBillingTypeSender,
				CommissionPolicy:      augno.UpdateCustomerRequestCommissionPolicyCommissionApplied,
				CreditLimit: augno.QuantityInputParam{
					UnitID: "un_82bd37dae5po",
					Value:  "10000.00",
				},
				CustomerPriceGroupIDs: []string{"acgp_6p4z57e9alaf"},
				CustomerTypeGroupID:   augno.String("acgp_6p4z57e9alaf"),
				DefaultCarrierID:      augno.String("cr_tv5vfjtgu1n3"),
				DefaultPaymentTermID:  augno.String("pytm_skssmsy21lem"),
				DefaultPriority:       augno.UpdateCustomerRequestDefaultPriorityNormal,
				DefaultSalesRepID:     augno.String("acus_e5zu8bde0z3h"),
				DefaultServiceLevelID: augno.String("crop_4ilk9p6gccrx"),
				DefaultShippingTermID: augno.String("shtm_c5gxy05whw6r"),
				EdiStatus:             augno.UpdateCustomerRequestEdiStatusDisabled,
				Email:                 augno.String("orders@acme.com"),
				FreightPolicy:         augno.UpdateCustomerRequestFreightPolicyBilledFreight,
				LeadTimeDays:          augno.Int(0),
				Name:                  augno.String("Acme Corp Updated"),
				Note:                  augno.String("Updated account notes"),
				Number:                augno.String("100042"),
				Phone:                 augno.String("555-123-4567"),
				ReceiveCalendarID:     augno.String("receive_calendar_id"),
				ShipToAddressID:       augno.String("ad_npqa5y43q26z"),
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
		EndsAt:                augno.Time(time.Now()),
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
		StartsAt:              augno.Time(time.Now()),
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
	_, err := client.Sales.Customers.Delete(context.TODO(), "ac_opnlh43ymyee")
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSaleCustomerGetLeadTime(t *testing.T) {
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
	_, err := client.Sales.Customers.GetLeadTime(context.TODO(), "ac_opnlh43ymyee")
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
