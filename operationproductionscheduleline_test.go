// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/augno/augno-go"
	"github.com/augno/augno-go/internal/testutil"
	"github.com/augno/augno-go/option"
)

func TestOperationProductionScheduleLineNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Lines.New(
		context.TODO(),
		"pnsc_0192a4c17b3e4f8a91c2d0",
		augno.OperationProductionScheduleLineNewParams{
			CreateProductionScheduleLineRequest: augno.CreateProductionScheduleLineRequestParam{
				ItemID:     "it_0131e386ac683e8c29a71f6f1f",
				MachineID:  "mc_0177d18f55a1615f783d3bf8d0",
				Quantity:   600,
				WeekIndex:  2,
				Lots:       augno.Int(0),
				Reason:     augno.CreateProductionScheduleLineRequestReasonMachineDown,
				ReasonNote: augno.String("reason_note"),
				RunHours:   augno.Float(0),
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

func TestOperationProductionScheduleLineUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Lines.Update(
		context.TODO(),
		"orln_0142f9b74268973450b3a76ce3",
		augno.OperationProductionScheduleLineUpdateParams{
			ID: "pnsc_0192a4c17b3e4f8a91c2d0",
			UpdateProductionScheduleLineRequest: augno.UpdateProductionScheduleLineRequestParam{
				Lots:          augno.Int(0),
				MachineID:     augno.String("machine_id"),
				Quantity:      augno.Float(900),
				Reason:        augno.UpdateProductionScheduleLineRequestReasonMachineDown,
				ReasonNote:    augno.String("reason_note"),
				RunHours:      augno.Float(0),
				SequenceIndex: augno.Int(0),
				Status:        augno.UpdateProductionScheduleLineRequestStatusPlanned,
				WeekIndex:     augno.Int(0),
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

func TestOperationProductionScheduleLineListWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Lines.List(
		context.TODO(),
		"pnsc_0192a4c17b3e4f8a91c2d0",
		augno.OperationProductionScheduleLineListParams{
			MachineIDs: []string{"string"},
			WeekIndex:  augno.Int(0),
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

func TestOperationProductionScheduleLineDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Lines.Delete(
		context.TODO(),
		"orln_0142f9b74268973450b3a76ce3",
		augno.OperationProductionScheduleLineDeleteParams{
			ID:         "pnsc_0192a4c17b3e4f8a91c2d0",
			Reason:     augno.OperationProductionScheduleLineDeleteParamsReasonMachineDown,
			ReasonNote: augno.String("reason_note"),
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
