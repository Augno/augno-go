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

func TestMessagingConversationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.Conversations.New(context.TODO(), augno.MessagingConversationNewParams{
		CreateConversationRequest: augno.CreateConversationRequestParam{
			Type:                      augno.CreateConversationRequestTypeDirectMessage,
			GroupID:                   augno.String("group_id"),
			ParticipantAccountUserIDs: []string{"acus_01ea9983ddb41dacc44ecf997c"},
			Title:                     augno.String("title"),
			TopicResourceID:           augno.String("topic_resource_id"),
			TopicResourceType:         augno.CreateConversationRequestTopicResourceTypeAccount,
		},
		Include: []string{"assignee"},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessagingConversationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.Conversations.Get(
		context.TODO(),
		"cv_01h9z8q1w2e3r4t5y6u7i8cv",
		augno.MessagingConversationGetParams{
			Include: []string{"assignee"},
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

func TestMessagingConversationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.Conversations.Update(
		context.TODO(),
		"cv_01h9z8q1w2e3r4t5y6u7i8cv",
		augno.MessagingConversationUpdateParams{
			Include: []string{"assignee"},
			UpdateConversationRequest: augno.UpdateConversationRequestParam{
				Title: augno.String("Fulfillment war room"),
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

func TestMessagingConversationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.Conversations.List(context.TODO(), augno.MessagingConversationListParams{
		AssigneeResourceID: augno.String("assignee_resource_id"),
		Audience:           augno.MessagingConversationListParamsAudienceInternal,
		Cursor:             augno.String("cursor"),
		Include:            []string{"assignee"},
		IncludeArchived:    augno.Bool(true),
		Limit:              augno.Int(0),
		Q:                  augno.String("q"),
		Status:             augno.MessagingConversationListParamsStatusActive,
		TopicResourceID:    augno.String("topic_resource_id"),
		TopicResourceType:  augno.MessagingConversationListParamsTopicResourceTypeAccount,
		Type:               augno.MessagingConversationListParamsTypeDirectMessage,
		Unassigned:         augno.Bool(true),
		WorkflowStatus:     augno.MessagingConversationListParamsWorkflowStatusNew,
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
