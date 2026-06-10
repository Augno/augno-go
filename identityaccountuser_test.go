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

func TestIdentityAccountUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Identity.AccountUsers.New(context.TODO(), augno.IdentityAccountUserNewParams{
		Include: []string{"user"},
		CreateAccountUserRequest: augno.CreateAccountUserRequestParam{
			DepartmentID: augno.String("department_id"),
			Email:        augno.String("jdoe@augno.com"),
			Name:         augno.String("John Doe"),
			Password:     augno.String("QgS7Z8Hhj3&1"),
			Preferences: []augno.NotificationPreferenceItemParam{{
				Enabled:          true,
				NotificationType: augno.NotificationPreferenceItemNotificationTypeOrderAcknowledgement,
			}},
			RoleID:   augno.String("rl_01c16d2eb637c0d1f3a372937c"),
			Username: augno.String("jdoe"),
		},
	})
	if err != nil {
		var apierr *augno.Error
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
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Identity.AccountUsers.Get(
		context.TODO(),
		"acus_01ea9983ddb41dacc44ecf997c",
		augno.IdentityAccountUserGetParams{
			Include: []string{"user"},
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

func TestIdentityAccountUserUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Identity.AccountUsers.Update(
		context.TODO(),
		"acus_01ea9983ddb41dacc44ecf997c",
		augno.IdentityAccountUserUpdateParams{
			Include: []string{"user"},
			UpdateAccountUserRequest: augno.UpdateAccountUserRequestParam{
				DepartmentID: augno.String("dp_01791c25ab59da4704cba61874"),
				Email:        augno.String("email"),
				Name:         augno.String("John Doe"),
				Preferences: []augno.NotificationPreferenceItemParam{{
					Enabled:          true,
					NotificationType: augno.NotificationPreferenceItemNotificationTypeInvoice,
				}},
				RoleID:   augno.String("rl_01c16d2eb637c0d1f3a372937c"),
				Username: augno.String("username"),
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

func TestIdentityAccountUserListWithOptionalParams(t *testing.T) {
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
	_, err := client.Identity.AccountUsers.List(context.TODO(), augno.IdentityAccountUserListParams{
		Cursor:       augno.String("cursor"),
		Include:      []string{"user"},
		Limit:        augno.Int(0),
		Q:            augno.String("q"),
		RemovedScope: augno.IdentityAccountUserListParamsRemovedScopeExcluded,
		RoleType:     augno.IdentityAccountUserListParamsRoleTypeAdmin,
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
