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
					ContentType:  augno.String("application/pdf"),
					Filename:     augno.String("quote.pdf"),
					ResourceID:   augno.String("resource_id"),
					ResourceType: augno.String("resource_type"),
					S3Key:        augno.String("uploads/acme/quote.pdf"),
					SizeBytes:    augno.Int(20480),
					URL:          augno.String("url"),
				}},
				Audience:              augno.SendMessageRequestAudienceCustomer,
				Cc:                    []string{"ap@acme.com"},
				Channel:               augno.SendMessageRequestChannelEmail,
				LinkResourceID:        augno.String("or_01d5034136c3ccc048abecc312"),
				LinkResourceType:      augno.SendMessageRequestLinkResourceTypeSalesOrder,
				Mentions:              []string{"acus_01ea9983ddb41dacc44ecf997c"},
				Mode:                  augno.SendMessageRequestModeSend,
				ReplyToMessageID:      augno.String("mg_01h9z8q1w2e3r4t5y6u7i8mg"),
				ScheduledAt:           augno.Time(time.Now()),
				SourceThreadMessageID: augno.String("mg_01h9z8q1w2e3r4t5y6u7i8mg"),
				Subject:               augno.String("Re: Order #1042"),
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
