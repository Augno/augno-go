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

func TestMessagingConversationMessageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.Conversations.Messages.New(
		context.TODO(),
		"cv_w35z4ck68yq7",
		openmrp.MessagingConversationMessageNewParams{
			SendMessageRequest: openmrp.SendMessageRequestParam{
				Body:            "Sounds good — shipping it today.",
				ClientMessageID: "client_msg_8c7d2f",
				Attachments: []openmrp.MessageAttachmentInputParam{{
					Kind:         openmrp.MessageAttachmentInputKindFile,
					ContentType:  openmrp.String("application/pdf"),
					Filename:     openmrp.String("quote.pdf"),
					ResourceID:   openmrp.String("resource_id"),
					ResourceType: openmrp.String("resource_type"),
					S3Key:        openmrp.String("uploads/acme/quote.pdf"),
					SizeBytes:    openmrp.Int(20480),
					URL:          openmrp.String("url"),
				}},
				Audience:              openmrp.SendMessageRequestAudienceCustomer,
				Cc:                    []string{"ap@acme.com"},
				Channel:               openmrp.SendMessageRequestChannelEmail,
				LinkResourceID:        openmrp.String("or_9lqo07quiwyb"),
				LinkResourceType:      openmrp.SendMessageRequestLinkResourceTypeSalesOrder,
				Mentions:              []string{"acus_e5zu8bde0z3h"},
				Mode:                  openmrp.SendMessageRequestModeSend,
				ReplyToMessageID:      openmrp.String("mg_fdny8633ebgw"),
				ScheduledAt:           openmrp.Time(time.Now()),
				SourceThreadMessageID: openmrp.String("mg_fdny8633ebgw"),
				Subject:               openmrp.String("Re: Order #1042"),
			},
			Include: []string{"sender"},
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

func TestMessagingConversationMessageListWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.Conversations.Messages.List(
		context.TODO(),
		"cv_w35z4ck68yq7",
		openmrp.MessagingConversationMessageListParams{
			AfterSequence: openmrp.Int(0),
			Cursor:        openmrp.String("cursor"),
			Include:       []string{"sender"},
			Limit:         openmrp.Int(0),
			Q:             openmrp.String("q"),
			Status:        openmrp.MessagingConversationMessageListParamsStatusDraft,
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
