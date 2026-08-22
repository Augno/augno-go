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

func TestOperationProductionScheduleActionArchive(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Actions.Archive(context.TODO(), "pnsc_m4zt3z8g8src")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperationProductionScheduleActionPreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Actions.Preview(context.TODO(), openmrp.OperationProductionScheduleActionPreviewParams{
		PreviewProductionScheduleRequest: openmrp.PreviewProductionScheduleRequestParam{
			DemandBasis:  openmrp.PreviewProductionScheduleRequestDemandBasisTrailing12,
			HorizonWeeks: openmrp.Int(13),
			PlanningAsOf: openmrp.Time(time.Now()),
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

func TestOperationProductionScheduleActionPreviewRegenerateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Actions.PreviewRegenerate(
		context.TODO(),
		"pnsc_m4zt3z8g8src",
		openmrp.OperationProductionScheduleActionPreviewRegenerateParams{
			PreviewRegenerateProductionScheduleRequest: openmrp.PreviewRegenerateProductionScheduleRequestParam{
				DemandBasis:  openmrp.PreviewRegenerateProductionScheduleRequestDemandBasisTrailing12,
				HorizonWeeks: openmrp.Int(13),
				PlanningAsOf: openmrp.Time(time.Now()),
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

func TestOperationProductionScheduleActionPublish(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Actions.Publish(context.TODO(), "pnsc_m4zt3z8g8src")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperationProductionScheduleActionRegenerateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Actions.Regenerate(
		context.TODO(),
		"pnsc_m4zt3z8g8src",
		openmrp.OperationProductionScheduleActionRegenerateParams{
			RegenerateProductionScheduleRequest: openmrp.RegenerateProductionScheduleRequestParam{
				DemandBasis:  openmrp.RegenerateProductionScheduleRequestDemandBasisTrailing12,
				HorizonWeeks: openmrp.Int(0),
				MergeMode:    openmrp.RegenerateProductionScheduleRequestMergeModePreserveManual,
				PlanningAsOf: openmrp.Time(time.Now()),
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

func TestOperationProductionScheduleActionReleaseWeekWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Actions.ReleaseWeek(
		context.TODO(),
		"pnsc_m4zt3z8g8src",
		openmrp.OperationProductionScheduleActionReleaseWeekParams{
			ReleaseProductionScheduleWeekRequest: openmrp.ReleaseProductionScheduleWeekRequestParam{
				ResponsibleUserID: "us_43irtlt2ajz6",
				WeekIndex:         0,
				ScanningStationID: openmrp.String("scanning_station_id"),
				SkipCarryForward:  openmrp.Bool(false),
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
