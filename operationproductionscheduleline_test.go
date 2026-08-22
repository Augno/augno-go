// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/open-mrp/openmrp-go"
	"github.com/open-mrp/openmrp-go/internal/testutil"
	"github.com/open-mrp/openmrp-go/option"
)

func TestOperationProductionScheduleLineNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Lines.New(
		context.TODO(),
		"pnsc_m4zt3z8g8src",
		openmrp.OperationProductionScheduleLineNewParams{
			CreateProductionScheduleLineRequest: openmrp.CreateProductionScheduleLineRequestParam{
				ItemID:     "it_pej07ckhvu62",
				MachineID:  "mc_ffcfk9dxixis",
				Quantity:   600,
				WeekIndex:  2,
				Lots:       openmrp.Int(0),
				Reason:     openmrp.CreateProductionScheduleLineRequestReasonMachineDown,
				ReasonNote: openmrp.String("reason_note"),
				RunHours:   openmrp.Float(0),
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

func TestOperationProductionScheduleLineUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Lines.Update(
		context.TODO(),
		"orln_la01fxgrwcnr",
		openmrp.OperationProductionScheduleLineUpdateParams{
			ID: "pnsc_m4zt3z8g8src",
			UpdateProductionScheduleLineRequest: openmrp.UpdateProductionScheduleLineRequestParam{
				Lots:          openmrp.Int(0),
				MachineID:     openmrp.String("machine_id"),
				Quantity:      openmrp.Float(900),
				Reason:        openmrp.UpdateProductionScheduleLineRequestReasonMachineDown,
				ReasonNote:    openmrp.String("reason_note"),
				RunHours:      openmrp.Float(0),
				SequenceIndex: openmrp.Int(0),
				Status:        openmrp.UpdateProductionScheduleLineRequestStatusPlanned,
				WeekIndex:     openmrp.Int(0),
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

func TestOperationProductionScheduleLineListWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Lines.List(
		context.TODO(),
		"pnsc_m4zt3z8g8src",
		openmrp.OperationProductionScheduleLineListParams{
			MachineIDs: []string{"string"},
			WeekIndex:  openmrp.Int(0),
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

func TestOperationProductionScheduleLineDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionSchedules.Lines.Delete(
		context.TODO(),
		"orln_la01fxgrwcnr",
		openmrp.OperationProductionScheduleLineDeleteParams{
			ID:         "pnsc_m4zt3z8g8src",
			Reason:     openmrp.OperationProductionScheduleLineDeleteParamsReasonMachineDown,
			ReasonNote: openmrp.String("reason_note"),
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
