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

func TestOperationScanningStationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ScanningStations.New(context.TODO(), augno.OperationScanningStationNewParams{
		CreateScanningStationRequest: augno.CreateScanningStationRequestParam{
			DepartmentID:        "dp_01791c25ab59da4704cba61874",
			Name:                "Packaging Line 1",
			OperatorRequirement: augno.CreateScanningStationRequestOperatorRequirementNone,
			Type:                augno.CreateScanningStationRequestTypeInitBatch,
			LabelSize:           augno.CreateScanningStationRequestLabelSize1x1,
			LabelType:           augno.CreateScanningStationRequestLabelTypeTag,
			Notes:               augno.String("Primary intake station on the receiving dock."),
		},
		Include: []string{"department"},
	})
	if err != nil {
		var apierr *augno.Error
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
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Operations.ScanningStations.Get(
		context.TODO(),
		"scst_0129335dd6286056a97024fcc1",
		augno.OperationScanningStationGetParams{
			Include: []string{"department"},
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

func TestOperationScanningStationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ScanningStations.Update(
		context.TODO(),
		"scst_0129335dd6286056a97024fcc1",
		augno.OperationScanningStationUpdateParams{
			Include: []string{"department"},
			UpdateScanningStationRequest: augno.UpdateScanningStationRequestParam{
				LabelSize:           augno.UpdateScanningStationRequestLabelSize1x1,
				LabelType:           augno.UpdateScanningStationRequestLabelTypeTag,
				Name:                augno.String("Station B"),
				Notes:               augno.String("Relocated to the finishing area."),
				OperatorRequirement: augno.UpdateScanningStationRequestOperatorRequirementMaterialCheck,
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

func TestOperationScanningStationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ScanningStations.List(context.TODO(), augno.OperationScanningStationListParams{
		Cursor:  augno.String("cursor"),
		Include: []string{"department"},
		Limit:   augno.Int(0),
		Q:       augno.String("q"),
	})
	if err != nil {
		var apierr *augno.Error
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
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Operations.ScanningStations.Delete(context.TODO(), "scst_0129335dd6286056a97024fcc1")
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
