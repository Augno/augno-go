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

func TestOperationDemandOverrideNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.DemandOverrides.New(context.TODO(), augno.OperationDemandOverrideNewParams{
		CreateDemandOverrideRequest: augno.CreateDemandOverrideRequestParam{
			Adjustment:     augno.CreateDemandOverrideRequestAdjustmentDeltaUnits,
			PeriodEndsAt:   time.Now(),
			PeriodStartsAt: time.Now(),
			ScopeRefID:     "it_pej07ckhvu62",
			ScopeType:      augno.CreateDemandOverrideRequestScopeTypeItem,
			Value:          5000,
			Active:         augno.Bool(false),
			EffectiveAt:    augno.Time(time.Now()),
			ExpiresAt:      augno.Time(time.Now()),
			Note:           augno.String("note"),
			Reason:         augno.CreateDemandOverrideRequestReasonNewCustomer,
			UnitID:         augno.String("unit_id"),
		},
		Include: []string{"scope"},
	})
	if err != nil {
		var apierr *augno.Error
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
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Operations.DemandOverrides.Get(
		context.TODO(),
		"deov_p8roudstrung",
		augno.OperationDemandOverrideGetParams{
			Include: []string{"scope"},
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

func TestOperationDemandOverrideUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.DemandOverrides.Update(
		context.TODO(),
		"deov_p8roudstrung",
		augno.OperationDemandOverrideUpdateParams{
			Include: []string{"scope"},
			UpdateDemandOverrideRequest: augno.UpdateDemandOverrideRequestParam{
				Active:         augno.Bool(false),
				Adjustment:     augno.UpdateDemandOverrideRequestAdjustmentAbsolute,
				ExpiresAt:      augno.Time(time.Now()),
				Note:           augno.String("note"),
				PeriodEndsAt:   augno.Time(time.Now()),
				PeriodStartsAt: augno.Time(time.Now()),
				Reason:         augno.UpdateDemandOverrideRequestReasonNewCustomer,
				UnitID:         augno.String("unit_id"),
				Value:          augno.Float(7500),
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

func TestOperationDemandOverrideListWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.DemandOverrides.List(context.TODO(), augno.OperationDemandOverrideListParams{
		Adjustments: []string{"absolute"},
		Cursor:      augno.String("cursor"),
		EndsAt:      augno.String("ends_at"),
		Include:     []string{"scope"},
		Limit:       augno.Int(0),
		Q:           augno.String("q"),
		ScopeRefIDs: []string{"string"},
		ScopeTypes:  []string{"item"},
		StartsAt:    augno.String("starts_at"),
		Statuses:    []string{"active"},
	})
	if err != nil {
		var apierr *augno.Error
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
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Operations.DemandOverrides.Delete(context.TODO(), "deov_p8roudstrung")
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
