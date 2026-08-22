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

func TestOperationDepartmentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.Departments.New(context.TODO(), openmrp.OperationDepartmentNewParams{
		CreateDepartmentRequest: openmrp.CreateDepartmentRequestParam{
			Name: "Fabrication",
			LaborRate: openmrp.DepartmentRateInputParam{
				DenominatorUnitID: "denominator_unit_id",
				NumeratorUnitID:   "numerator_unit_id",
				Value:             "value",
			},
			LocationID:         openmrp.String("location_id"),
			MachineIDs:         []string{"mc_ffcfk9dxixis"},
			Notes:              openmrp.String("notes"),
			ScanningStationIDs: []string{"scst_t71bn7lq5yov"},
		},
		Include: []string{"location"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOperationDepartmentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.Departments.Get(
		context.TODO(),
		"dp_m0jayebxnkos",
		openmrp.OperationDepartmentGetParams{
			Include: []string{"location"},
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

func TestOperationDepartmentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.Departments.Update(
		context.TODO(),
		"dp_m0jayebxnkos",
		openmrp.OperationDepartmentUpdateParams{
			Include: []string{"location"},
			UpdateDepartmentRequest: openmrp.UpdateDepartmentRequestParam{
				LaborRate: openmrp.DepartmentRateInputParam{
					DenominatorUnitID: "denominator_unit_id",
					NumeratorUnitID:   "numerator_unit_id",
					Value:             "value",
				},
				LocationID:         openmrp.String("location_id"),
				MachineIDs:         []string{"string"},
				Name:               openmrp.String("Production"),
				Notes:              openmrp.String("notes"),
				ScanningStationIDs: []string{"string"},
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

func TestOperationDepartmentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.Departments.List(context.TODO(), openmrp.OperationDepartmentListParams{
		Cursor:  openmrp.String("cursor"),
		Include: []string{"location"},
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

func TestOperationDepartmentDelete(t *testing.T) {
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
	_, err := client.Operations.Departments.Delete(context.TODO(), "dp_m0jayebxnkos")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
