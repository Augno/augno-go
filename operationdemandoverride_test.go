// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/open-mrp/openmrp-go"
	"github.com/open-mrp/openmrp-go/internal/testutil"
	"github.com/open-mrp/openmrp-go/option"
)

func TestOperationDemandOverrideNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.DemandOverrides.New(context.TODO(), openmrp.OperationDemandOverrideNewParams{
		CreateDemandOverrideRequest: openmrp.CreateDemandOverrideRequestParam{
			Adjustment:     openmrp.CreateDemandOverrideRequestAdjustmentDeltaUnits,
			PeriodEndsAt:   time.Now(),
			PeriodStartsAt: time.Now(),
			ScopeRefID:     "it_pej07ckhvu62",
			ScopeType:      openmrp.CreateDemandOverrideRequestScopeTypeItem,
			Value:          5000,
			Active:         openmrp.Bool(false),
			EffectiveAt:    openmrp.Time(time.Now()),
			ExpiresAt:      openmrp.Time(time.Now()),
			Note:           openmrp.String("note"),
			Reason:         openmrp.CreateDemandOverrideRequestReasonNewCustomer,
			UnitID:         openmrp.String("unit_id"),
		},
		Include: []string{"scope"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperationDemandOverrideGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.DemandOverrides.Get(
		context.TODO(),
		"deov_p8roudstrung",
		openmrp.OperationDemandOverrideGetParams{
			Include: []string{"scope"},
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

func TestOperationDemandOverrideUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.DemandOverrides.Update(
		context.TODO(),
		"deov_p8roudstrung",
		openmrp.OperationDemandOverrideUpdateParams{
			Include: []string{"scope"},
			UpdateDemandOverrideRequest: openmrp.UpdateDemandOverrideRequestParam{
				Active:         openmrp.Bool(false),
				Adjustment:     openmrp.UpdateDemandOverrideRequestAdjustmentAbsolute,
				ExpiresAt:      openmrp.Time(time.Now()),
				Note:           openmrp.String("note"),
				PeriodEndsAt:   openmrp.Time(time.Now()),
				PeriodStartsAt: openmrp.Time(time.Now()),
				Reason:         openmrp.UpdateDemandOverrideRequestReasonNewCustomer,
				UnitID:         openmrp.String("unit_id"),
				Value:          openmrp.Float(7500),
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

func TestOperationDemandOverrideListWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.DemandOverrides.List(context.TODO(), openmrp.OperationDemandOverrideListParams{
		Adjustments: []string{"absolute"},
		Cursor:      openmrp.String("cursor"),
		EndsAt:      openmrp.String("ends_at"),
		Include:     []string{"scope"},
		Limit:       openmrp.Int(0),
		Q:           openmrp.String("q"),
		ScopeRefIDs: []string{"string"},
		ScopeTypes:  []string{"item"},
		StartsAt:    openmrp.String("starts_at"),
		Statuses:    []string{"active"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperationDemandOverrideDelete(t *testing.T) {
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
	_, err := client.Operations.DemandOverrides.Delete(context.TODO(), "deov_p8roudstrung")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
