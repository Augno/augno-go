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

func TestAIAgentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Agents.New(context.TODO(), augno.AIAgentNewParams{
		CreateAgentRequest: augno.CreateAgentRequestParam{
			CategoryCode: "inventory",
			Config: augno.ConfigInputParam{
				EndpointToolReview: map[string]bool{
					"foo": true,
				},
				EndpointToolSlugs: []string{"string"},
				SystemPrompt:      augno.String("You are an order processing agent. Parse incoming emails and create draft orders."),
				Temperature:       augno.Float(0.2),
				Tier:              augno.ConfigInputTierHigh,
				TriggerConfig: augno.TriggerConfigInputParam{
					CronSchedule: augno.String("cron_schedule"),
					EventFilters: []string{"email.received"},
					Timezone:     augno.String("timezone"),
				},
			},
			Name:        "Inventory Monitor",
			Slug:        "inventory_monitor",
			TriggerType: augno.CreateAgentRequestTriggerTypeEvent,
			Description: augno.String("Monitors inventory levels and creates restock alerts."),
			RoleID:      augno.String("rl_3xknmfqflhvb"),
			Tools: []augno.ToolInputParam{{
				Tool:          augno.ToolInputToolReadDoc,
				ConfigJson:    augno.String("config_json"),
				RequireReview: augno.Bool(true),
				SortOrder:     augno.Int(1),
			}},
		},
		Include: []string{"config"},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAgentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Agents.Get(
		context.TODO(),
		"agdf_ah7tkyfxk8jl",
		augno.AIAgentGetParams{
			Include: []string{"config"},
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

func TestAIAgentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Agents.Update(
		context.TODO(),
		"agdf_ah7tkyfxk8jl",
		augno.AIAgentUpdateParams{
			Include: []string{"config"},
			UpdateAgentRequest: augno.UpdateAgentRequestParam{
				CategoryCode: augno.String("category_code"),
				Config: augno.ConfigInputParam{
					EndpointToolReview: map[string]bool{
						"foo": true,
					},
					EndpointToolSlugs: []string{"string"},
					SystemPrompt:      augno.String("You are an order processing agent. Parse incoming emails and create draft orders."),
					Temperature:       augno.Float(0.2),
					Tier:              augno.ConfigInputTierHigh,
					TriggerConfig: augno.TriggerConfigInputParam{
						CronSchedule: augno.String("cron_schedule"),
						EventFilters: []string{"email.received"},
						Timezone:     augno.String("timezone"),
					},
				},
				Description: augno.String("description"),
				Name:        augno.String("Inventory Monitor"),
				RoleID:      augno.String("role_id"),
				Slug:        augno.String("slug"),
				Tools: []augno.ToolInputParam{{
					Tool:          augno.ToolInputToolReadDoc,
					ConfigJson:    augno.String("config_json"),
					RequireReview: augno.Bool(true),
					SortOrder:     augno.Int(1),
				}},
				TriggerType: augno.UpdateAgentRequestTriggerTypeScheduled,
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

func TestAIAgentListWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Agents.List(context.TODO(), augno.AIAgentListParams{
		Cursor:          augno.String("cursor"),
		DefinitionTypes: []string{"system"},
		Include:         []string{"config"},
		Limit:           augno.Int(0),
		Q:               augno.String("q"),
		Statuses:        []string{"active"},
		TriggerTypes:    []string{"scheduled"},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAgentDelete(t *testing.T) {
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
	_, err := client.AI.Agents.Delete(context.TODO(), "agdf_ah7tkyfxk8jl")
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAgentUpdateStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Agents.UpdateStatus(
		context.TODO(),
		"agdf_ah7tkyfxk8jl",
		augno.AIAgentUpdateStatusParams{
			UpdateAgentStatusRequest: augno.UpdateAgentStatusRequestParam{
				Status: "active",
			},
			Include: []string{"config"},
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
