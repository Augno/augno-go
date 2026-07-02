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

func TestMessagingConversationMessageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.Conversations.Messages.New(
		context.TODO(),
		"cv_01h9z8q1w2e3r4t5y6u7i8cv",
		augno.MessagingConversationMessageNewParams{
			SendMessageRequest: augno.SendMessageRequestParam{
				Body:            "Sounds good — shipping it today.",
				ClientMessageID: "client_msg_8c7d2f",
				Attachments: []augno.MessageAttachmentInputParam{{
					Kind:         augno.MessageAttachmentInputKindFile,
					ContentType:  augno.String("content_type"),
					Filename:     augno.String("filename"),
					ResourceID:   augno.String("resource_id"),
					ResourceType: augno.String("resource_type"),
					S3Key:        augno.String("s3_key"),
					SizeBytes:    augno.Int(0),
					URL:          augno.String("url"),
				}},
				Audience:              augno.SendMessageRequestAudienceInternal,
				Cc:                    []string{"string"},
				Channel:               augno.SendMessageRequestChannelMessage,
				LinkResourceID:        augno.String("link_resource_id"),
				LinkResourceType:      augno.SendMessageRequestLinkResourceTypeAccount,
				Mentions:              []string{"string"},
				Mode:                  augno.SendMessageRequestModeSend,
				ReplyToMessageID:      augno.String("reply_to_message_id"),
				ScheduledAt:           augno.Time(time.Now()),
				SourceThreadMessageID: augno.String("source_thread_message_id"),
				Subject:               augno.String("subject"),
			},
			Include: []string{"sender"},
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

func TestMessagingConversationMessageListWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.Conversations.Messages.List(
		context.TODO(),
		"cv_01h9z8q1w2e3r4t5y6u7i8cv",
		augno.MessagingConversationMessageListParams{
			AfterSequence: augno.Int(0),
			Cursor:        augno.String("cursor"),
			Include:       []string{"sender"},
			Limit:         augno.Int(0),
			Q:             augno.String("q"),
			Status:        augno.MessagingConversationMessageListParamsStatusDraft,
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
