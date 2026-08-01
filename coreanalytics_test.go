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
			EndDate:       time.Now(),
			StartDate:     time.Now(),
			DepartmentIDs: []string{"dp_01791c25ab59da4704cba61874"},
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
			EndDate:       time.Now(),
			StartDate:     time.Now(),
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
