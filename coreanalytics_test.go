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

func TestCoreAnalyticsUpdateDeliveryPerformanceWithOptionalParams(t *testing.T) {
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
	_, err := client.Core.Analytics.UpdateDeliveryPerformance(context.TODO(), openmrp.CoreAnalyticsUpdateDeliveryPerformanceParams{
		AnalyzeDeliveryPerformanceRequest: openmrp.AnalyzeDeliveryPerformanceRequestParam{
			EndsAt:           time.Now(),
			StartsAt:         time.Now(),
			CustomerGroupIDs: []string{"string"},
			CustomerIDs:      []string{"string"},
			Granularity:      openmrp.AnalyzeDeliveryPerformanceRequestGranularityWeek,
			ProductLineIDs:   []string{"string"},
			SalesRepIDs:      []string{"string"},
		},
	})
	if err != nil {
		var apierr *openmrp.Error
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
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Core.Analytics.UpdateOee(context.TODO(), openmrp.CoreAnalyticsUpdateOeeParams{
		AnalyzeOeeRequest: openmrp.AnalyzeOeeRequestParam{
			EndsAt:        time.Now(),
			StartsAt:      time.Now(),
			DepartmentIDs: []string{"dp_m0jayebxnkos"},
			PlannedTime: []openmrp.OeeDepartmentPlannedTimeParam{{
				DepartmentID: "department_id",
				PlannedHours: 0,
			}},
		},
	})
	if err != nil {
		var apierr *openmrp.Error
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
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Core.Analytics.UpdateOeeTrend(context.TODO(), openmrp.CoreAnalyticsUpdateOeeTrendParams{
		AnalyzeOeeTrendRequest: openmrp.AnalyzeOeeTrendRequestParam{
			EndsAt:        time.Now(),
			StartsAt:      time.Now(),
			DepartmentIDs: []string{"dp_m0jayebxnkos"},
		},
	})
	if err != nil {
		var apierr *openmrp.Error
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
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Core.Analytics.UpdateScheduleAttainment(context.TODO(), openmrp.CoreAnalyticsUpdateScheduleAttainmentParams{
		AnalyzeScheduleAttainmentRequest: openmrp.AnalyzeScheduleAttainmentRequestParam{
			EndsAt:        time.Now(),
			StartsAt:      time.Now(),
			DepartmentIDs: []string{"string"},
			GroupBy:       openmrp.AnalyzeScheduleAttainmentRequestGroupByWeek,
			MachineIDs:    []string{"string"},
		},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
