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

func TestOperationProductionScheduleActionArchive(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Actions.Archive(context.TODO(), "pnsc_0192a4c17b3e4f8a91c2d0")
	if err != nil {
		var apierr *augno.Error
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
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Operations.ProductionSchedules.Actions.Preview(context.TODO(), augno.OperationProductionScheduleActionPreviewParams{
		PreviewProductionScheduleRequest: augno.PreviewProductionScheduleRequestParam{
			DemandBasis:  augno.PreviewProductionScheduleRequestDemandBasisTrailing12,
			HorizonWeeks: augno.Int(13),
			PlanningAsOf: augno.Time(time.Now()),
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

func TestOperationProductionScheduleActionPreviewRegenerateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Actions.PreviewRegenerate(
		context.TODO(),
		"pnsc_0192a4c17b3e4f8a91c2d0",
		augno.OperationProductionScheduleActionPreviewRegenerateParams{
			PreviewRegenerateProductionScheduleRequest: augno.PreviewRegenerateProductionScheduleRequestParam{
				DemandBasis:  augno.PreviewRegenerateProductionScheduleRequestDemandBasisTrailing12,
				HorizonWeeks: augno.Int(13),
				PlanningAsOf: augno.Time(time.Now()),
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

func TestOperationProductionScheduleActionPublish(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Actions.Publish(context.TODO(), "pnsc_0192a4c17b3e4f8a91c2d0")
	if err != nil {
		var apierr *augno.Error
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
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Operations.ProductionSchedules.Actions.Regenerate(
		context.TODO(),
		"pnsc_0192a4c17b3e4f8a91c2d0",
		augno.OperationProductionScheduleActionRegenerateParams{
			RegenerateProductionScheduleRequest: augno.RegenerateProductionScheduleRequestParam{
				DemandBasis:  augno.RegenerateProductionScheduleRequestDemandBasisTrailing12,
				HorizonWeeks: augno.Int(0),
				MergeMode:    augno.RegenerateProductionScheduleRequestMergeModePreserveManual,
				PlanningAsOf: augno.Time(time.Now()),
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

func TestOperationProductionScheduleActionReleaseWeekWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Actions.ReleaseWeek(
		context.TODO(),
		"pnsc_0192a4c17b3e4f8a91c2d0",
		augno.OperationProductionScheduleActionReleaseWeekParams{
			ReleaseProductionScheduleWeekRequest: augno.ReleaseProductionScheduleWeekRequestParam{
				ResponsibleUserID: "us_0151164dcaea4cbded27b50aae",
				WeekIndex:         0,
				ScanningStationID: augno.String("scanning_station_id"),
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
