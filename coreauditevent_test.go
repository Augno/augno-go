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

func TestCoreAuditEventGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Core.AuditEvents.Get(
		context.TODO(),
		"ae_emripvn8t1xl",
		openmrp.CoreAuditEventGetParams{
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

func TestCoreAuditEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.Core.AuditEvents.List(context.TODO(), openmrp.CoreAuditEventListParams{
		Actions:          []string{"create"},
		ActorAccountIDs:  []string{"string"},
		ActorIDs:         []string{"string"},
		ActorTypes:       []string{"user"},
		Cursor:           openmrp.String("cursor"),
		EndsAt:           openmrp.Time(time.Now()),
		Include:          []string{"account"},
		Limit:            openmrp.Int(0),
		Q:                openmrp.String("q"),
		ResourceIDs:      []string{"string"},
		ResourceTypes:    []string{"account"},
		RootResourceID:   openmrp.String("root_resource_id"),
		RootResourceType: openmrp.CoreAuditEventListParamsRootResourceTypeAccount,
		StartsAt:         openmrp.Time(time.Now()),
		TargetAccountIDs: []string{"string"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCoreAuditEventGetResourceTypes(t *testing.T) {
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
	_, err := client.Core.AuditEvents.GetResourceTypes(context.TODO())
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
