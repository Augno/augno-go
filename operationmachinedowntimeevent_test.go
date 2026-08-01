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

func TestOperationMachineDowntimeEventNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.MachineDowntimeEvents.New(context.TODO(), augno.OperationMachineDowntimeEventNewParams{
		CreateMachineDowntimeEventRequest: augno.CreateMachineDowntimeEventRequestParam{
			MachineID:       "mc_0177d18f55a1615f783d3bf8d0",
			Reason:          augno.CreateMachineDowntimeEventRequestReasonBreakdown,
			StartedAt:       time.Now(),
			BatchID:         augno.String("batch_id"),
			EndedAt:         augno.Time(time.Now()),
			ItemID:          augno.String("item_id"),
			Note:            augno.String("note"),
			ProductionRunID: augno.String("production_run_id"),
			Source:          augno.CreateMachineDowntimeEventRequestSourceManual,
		},
		Include: []string{"machine"},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperationMachineDowntimeEventGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.MachineDowntimeEvents.Get(
		context.TODO(),
		"mcdt_0192a4c17b3e4f8a91c2d05e77",
		augno.OperationMachineDowntimeEventGetParams{
			Include: []string{"machine"},
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

func TestOperationMachineDowntimeEventUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.MachineDowntimeEvents.Update(
		context.TODO(),
		"mcdt_0192a4c17b3e4f8a91c2d05e77",
		augno.OperationMachineDowntimeEventUpdateParams{
			Include: []string{"machine"},
			UpdateMachineDowntimeEventRequest: augno.UpdateMachineDowntimeEventRequestParam{
				BatchID:         augno.String("batch_id"),
				EndedAt:         augno.Time(time.Now()),
				ItemID:          augno.String("item_id"),
				Note:            augno.String("note"),
				ProductionRunID: augno.String("production_run_id"),
				Reason:          augno.UpdateMachineDowntimeEventRequestReasonBreakdown,
				StartedAt:       augno.Time(time.Now()),
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

func TestOperationMachineDowntimeEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.MachineDowntimeEvents.List(context.TODO(), augno.OperationMachineDowntimeEventListParams{
		Cursor:        augno.String("cursor"),
		DepartmentIDs: []string{"string"},
		EndDate:       augno.String("end_date"),
		Include:       []string{"machine"},
		Limit:         augno.Int(0),
		MachineIDs:    []string{"string"},
		Open:          augno.Bool(true),
		Q:             augno.String("q"),
		Reasons:       []string{"breakdown"},
		StartDate:     augno.String("start_date"),
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperationMachineDowntimeEventDelete(t *testing.T) {
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
	_, err := client.Operations.MachineDowntimeEvents.Delete(context.TODO(), "mcdt_0192a4c17b3e4f8a91c2d05e77")
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
