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

func TestCoreAnalyticsUpdateDeliveryPerformanceWithOptionalParams(t *testing.T) {
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
	_, err := client.Core.Analytics.UpdateDeliveryPerformance(context.TODO(), augno.CoreAnalyticsUpdateDeliveryPerformanceParams{
		AnalyzeDeliveryPerformanceRequest: augno.AnalyzeDeliveryPerformanceRequestParam{
			EndsAt:           time.Now(),
			StartsAt:         time.Now(),
			CustomerGroupIDs: []string{"string"},
			CustomerIDs:      []string{"string"},
			Granularity:      augno.AnalyzeDeliveryPerformanceRequestGranularityWeek,
			ProductLineIDs:   []string{"string"},
			SalesRepIDs:      []string{"string"},
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

func TestCoreAnalyticsUpdateOeeWithOptionalParams(t *testing.T) {
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
	_, err := client.Core.Analytics.UpdateOee(context.TODO(), augno.CoreAnalyticsUpdateOeeParams{
		AnalyzeOeeRequest: augno.AnalyzeOeeRequestParam{
			EndsAt:        time.Now(),
			StartsAt:      time.Now(),
			DepartmentIDs: []string{"dp_m0jayebxnkos"},
			PlannedTime: []augno.OeeDepartmentPlannedTimeParam{{
				DepartmentID: "department_id",
				PlannedHours: 0,
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

func TestCoreAnalyticsUpdateOeeTrendWithOptionalParams(t *testing.T) {
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
	_, err := client.Core.Analytics.UpdateOeeTrend(context.TODO(), augno.CoreAnalyticsUpdateOeeTrendParams{
		AnalyzeOeeTrendRequest: augno.AnalyzeOeeTrendRequestParam{
			EndsAt:        time.Now(),
			StartsAt:      time.Now(),
			DepartmentIDs: []string{"dp_m0jayebxnkos"},
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

func TestCoreAnalyticsUpdateScheduleAttainmentWithOptionalParams(t *testing.T) {
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
	_, err := client.Core.Analytics.UpdateScheduleAttainment(context.TODO(), augno.CoreAnalyticsUpdateScheduleAttainmentParams{
		AnalyzeScheduleAttainmentRequest: augno.AnalyzeScheduleAttainmentRequestParam{
			EndsAt:        time.Now(),
			StartsAt:      time.Now(),
			DepartmentIDs: []string{"string"},
			GroupBy:       augno.AnalyzeScheduleAttainmentRequestGroupByWeek,
			MachineIDs:    []string{"string"},
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
