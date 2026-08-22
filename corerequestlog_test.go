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

func TestCoreRequestLogGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Core.RequestLogs.Get(
		context.TODO(),
		"rq_0lhl3kkhme40",
		openmrp.CoreRequestLogGetParams{
			Include: []string{"account"},
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

func TestCoreRequestLogListWithOptionalParams(t *testing.T) {
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
	_, err := client.Core.RequestLogs.List(context.TODO(), openmrp.CoreRequestLogListParams{
		ActorAccountIDs:   []string{"string"},
		ActorIDs:          []string{"string"},
		ActorTypes:        []string{"user"},
		Cursor:            openmrp.String("cursor"),
		EndsAt:            openmrp.Time(time.Now()),
		ErrorCodes:        []string{"expired_token"},
		ExcludeErrorCodes: []string{"expired_token"},
		Hosts:             []string{"string"},
		IdempotencyKey:    openmrp.String("idempotency_key"),
		Include:           []string{"account"},
		Limit:             openmrp.Int(0),
		Methods:           []string{"GET"},
		MinLatencyUs:      openmrp.Int(0),
		NormalizedRoutes:  []string{"string"},
		Q:                 openmrp.String("q"),
		StartsAt:          openmrp.Time(time.Now()),
		StatusCodeClasses: []int64{0},
		StatusCodes:       []int64{0},
		TargetAccountIDs:  []string{"string"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
