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

func TestOperationScanningStationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ScanningStations.New(context.TODO(), openmrp.OperationScanningStationNewParams{
		CreateScanningStationRequest: openmrp.CreateScanningStationRequestParam{
			DepartmentID:        "dp_m0jayebxnkos",
			Name:                "Packaging Line 1",
			OperatorRequirement: openmrp.CreateScanningStationRequestOperatorRequirementNone,
			Type:                openmrp.CreateScanningStationRequestTypeInitBatch,
			LabelSize:           openmrp.CreateScanningStationRequestLabelSize1x1,
			LabelType:           openmrp.CreateScanningStationRequestLabelTypeTag,
			Notes:               openmrp.String("Primary intake station on the receiving dock."),
		},
		Include: []string{"department"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperationScanningStationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ScanningStations.Get(
		context.TODO(),
		"scst_t71bn7lq5yov",
		openmrp.OperationScanningStationGetParams{
			Include: []string{"department"},
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

func TestOperationScanningStationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ScanningStations.Update(
		context.TODO(),
		"scst_t71bn7lq5yov",
		openmrp.OperationScanningStationUpdateParams{
			Include: []string{"department"},
			UpdateScanningStationRequest: openmrp.UpdateScanningStationRequestParam{
				LabelSize:           openmrp.UpdateScanningStationRequestLabelSize1x1,
				LabelType:           openmrp.UpdateScanningStationRequestLabelTypeTag,
				Name:                openmrp.String("Station B"),
				Notes:               openmrp.String("Relocated to the finishing area."),
				OperatorRequirement: openmrp.UpdateScanningStationRequestOperatorRequirementMaterialCheck,
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

func TestOperationScanningStationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ScanningStations.List(context.TODO(), openmrp.OperationScanningStationListParams{
		Cursor:  openmrp.String("cursor"),
		Include: []string{"department"},
		Limit:   openmrp.Int(0),
		Q:       openmrp.String("q"),
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperationScanningStationDelete(t *testing.T) {
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
	_, err := client.Operations.ScanningStations.Delete(context.TODO(), "scst_t71bn7lq5yov")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
