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

func TestOperationMachineDowntimeEventNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.MachineDowntimeEvents.New(context.TODO(), openmrp.OperationMachineDowntimeEventNewParams{
		CreateMachineDowntimeEventRequest: openmrp.CreateMachineDowntimeEventRequestParam{
			MachineID: "mc_ffcfk9dxixis",
			Reason:    openmrp.CreateMachineDowntimeEventRequestReasonBreakdown,
			StartedAt: time.Now(),
			BatchID:   openmrp.String("batch_id"),
			Duration: openmrp.QuantityInputParam{
				UnitID: "unit_id",
				Value:  "value",
			},
			EndedAt:         openmrp.Time(time.Now()),
			ItemID:          openmrp.String("item_id"),
			Note:            openmrp.String("note"),
			ProductionRunID: openmrp.String("production_run_id"),
			Source:          openmrp.CreateMachineDowntimeEventRequestSourceManual,
		},
		Include: []string{"machine"},
	})
	if err != nil {
		var apierr *openmrp.Error
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
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Operations.MachineDowntimeEvents.Get(
		context.TODO(),
		"mcdt_ff5te1hqttco",
		openmrp.OperationMachineDowntimeEventGetParams{
			Include: []string{"machine"},
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

func TestOperationMachineDowntimeEventUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.MachineDowntimeEvents.Update(
		context.TODO(),
		"mcdt_ff5te1hqttco",
		openmrp.OperationMachineDowntimeEventUpdateParams{
			Include: []string{"machine"},
			UpdateMachineDowntimeEventRequest: openmrp.UpdateMachineDowntimeEventRequestParam{
				BatchID: openmrp.String("batch_id"),
				Duration: openmrp.QuantityInputParam{
					UnitID: "unit_id",
					Value:  "value",
				},
				EndedAt:         openmrp.Time(time.Now()),
				ItemID:          openmrp.String("item_id"),
				MachineID:       openmrp.String("machine_id"),
				Note:            openmrp.String("note"),
				ProductionRunID: openmrp.String("production_run_id"),
				Reason:          openmrp.UpdateMachineDowntimeEventRequestReasonBreakdown,
				StartedAt:       openmrp.Time(time.Now()),
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

func TestOperationMachineDowntimeEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.MachineDowntimeEvents.List(context.TODO(), openmrp.OperationMachineDowntimeEventListParams{
		Cursor:        openmrp.String("cursor"),
		DepartmentIDs: []string{"string"},
		EndsAt:        openmrp.String("ends_at"),
		Include:       []string{"machine"},
		Limit:         openmrp.Int(0),
		MachineIDs:    []string{"string"},
		Open:          openmrp.Bool(true),
		Q:             openmrp.String("q"),
		Reasons:       []string{"breakdown"},
		StartsAt:      openmrp.String("starts_at"),
	})
	if err != nil {
		var apierr *openmrp.Error
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
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Operations.MachineDowntimeEvents.Delete(context.TODO(), "mcdt_ff5te1hqttco")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
