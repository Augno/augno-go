// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/augno-go"
	"github.com/stainless-sdks/augno-go/internal/testutil"
	"github.com/stainless-sdks/augno-go/option"
)

func TestAuthPasswordActionRequestPasswordResetWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Auth.Passwords.Actions.RequestPasswordReset(context.TODO(), augno.AuthPasswordActionRequestPasswordResetParams{
		AccountSlug: augno.String("account_slug"),
		Identifier:  augno.String("jdoe@augno.com"),
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthPasswordActionResetPasswordWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := augno.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Auth.Passwords.Actions.ResetPassword(context.TODO(), augno.AuthPasswordActionResetPasswordParams{
		Token:    augno.String("eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1Z25vLmNvbSIsInN1YiI6InVzXzAxZ2Y3YTgyMDBlMXNyMjBwZzl3eDZkMmswIiwiZXhwIjoxNzU2ODIzMzI5LCJpYXQiOjE3NTY4MTk3Mjl9.2ZodhtiHDqIQnDjzrJZvqIdEbQbmkgbTaz4OXdbXCWNjzEsy2-5e78XQRu-aZ8MoZ2dusIVKQcN1Tm-arKR0_Q"),
		Password: augno.String("new-super-secret-password"),
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
