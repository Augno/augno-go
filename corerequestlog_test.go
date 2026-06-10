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

func TestCoreRequestLogGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Core.RequestLogs.Get(
		context.TODO(),
		"rq_01304bffe90e8cce9690cbefd4",
		augno.CoreRequestLogGetParams{
			Include: []string{"account"},
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

func TestCoreRequestLogListWithOptionalParams(t *testing.T) {
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
	_, err := client.Core.RequestLogs.List(context.TODO(), augno.CoreRequestLogListParams{
		AccountIDs:        []string{"string"},
		ActorIDs:          []string{"string"},
		ActorTypes:        []string{"user"},
		Cursor:            augno.String("cursor"),
		EndDate:           augno.Time(time.Now()),
		ErrorCodes:        []string{"expired_token"},
		Hosts:             []string{"string"},
		IdempotencyKey:    augno.String("idempotency_key"),
		Include:           []string{"account"},
		Limit:             augno.Int(0),
		Methods:           []string{"GET"},
		MinLatencyUs:      augno.Int(0),
		NormalizedRoutes:  []string{"string"},
		Q:                 augno.String("q"),
		StartDate:         augno.Time(time.Now()),
		StatusCodeClasses: []int64{0},
		StatusCodes:       []int64{0},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
