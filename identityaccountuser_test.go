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

func TestIdentityAccountUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Identity.AccountUsers.New(context.TODO(), openmrp.IdentityAccountUserNewParams{
		Include: []string{"user"},
		CreateAccountUserRequest: openmrp.CreateAccountUserRequestParam{
			DepartmentID:         openmrp.String("dp_m0jayebxnkos"),
			Email:                openmrp.String("jdoe@openmrp.ai"),
			IsCommissionEligible: openmrp.Bool(false),
			Name:                 openmrp.String("John Doe"),
			Password:             openmrp.String("QgS7Z8Hhj3&1"),
			Preferences: []openmrp.NotificationPreferenceItemParam{{
				Enabled:          true,
				NotificationType: openmrp.NotificationPreferenceItemNotificationTypeOrderAcknowledgement,
			}},
			RoleID:   openmrp.String("rl_3xknmfqflhvb"),
			Username: openmrp.String("jdoe"),
		},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIdentityAccountUserGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Identity.AccountUsers.Get(
		context.TODO(),
		"acus_e5zu8bde0z3h",
		openmrp.IdentityAccountUserGetParams{
			Include: []string{"user"},
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

func TestIdentityAccountUserUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Identity.AccountUsers.Update(
		context.TODO(),
		"acus_e5zu8bde0z3h",
		openmrp.IdentityAccountUserUpdateParams{
			Include: []string{"user"},
			UpdateAccountUserRequest: openmrp.UpdateAccountUserRequestParam{
				DepartmentID:         openmrp.String("dp_m0jayebxnkos"),
				Email:                openmrp.String("jdoe@openmrp.ai"),
				IsCommissionEligible: openmrp.Bool(false),
				Name:                 openmrp.String("John Doe"),
				Preferences: []openmrp.NotificationPreferenceItemParam{{
					Enabled:          true,
					NotificationType: openmrp.NotificationPreferenceItemNotificationTypeOrderAcknowledgement,
				}},
				RoleID:   openmrp.String("rl_3xknmfqflhvb"),
				Username: openmrp.String("jdoe"),
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

func TestIdentityAccountUserListWithOptionalParams(t *testing.T) {
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
	_, err := client.Identity.AccountUsers.List(context.TODO(), openmrp.IdentityAccountUserListParams{
		Cursor:               openmrp.String("cursor"),
		Include:              []string{"user"},
		IsCommissionEligible: openmrp.Bool(true),
		Limit:                openmrp.Int(0),
		Q:                    openmrp.String("q"),
		RemovedScope:         openmrp.IdentityAccountUserListParamsRemovedScopeExcluded,
		RoleType:             openmrp.IdentityAccountUserListParamsRoleTypeAdmin,
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
