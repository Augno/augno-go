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

func TestOperationProductionScheduleSettingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionScheduleSettings.Update(context.TODO(), augno.OperationProductionScheduleSettingUpdateParams{
		UpdateProductionScheduleSettingsRequest: augno.UpdateProductionScheduleSettingsRequestParam{
			AutoPublishStatus:              augno.UpdateProductionScheduleSettingsRequestAutoPublishStatusInactive,
			CadenceStatus:                  augno.UpdateProductionScheduleSettingsRequestCadenceStatusInactive,
			CapacityHeadroomPct:            0.9,
			ChangeoverAvgMinutes:           0,
			ChangeoverLaborRate:            0,
			ChangeoverMaxMinutes:           0,
			ChangeoverMinMinutes:           0,
			DefaultConstraintLeadTimeWeeks: 0,
			DefaultCustomerLeadTimeDays:    30,
			DefaultFulfillmentPolicy:       augno.UpdateProductionScheduleSettingsRequestDefaultFulfillmentPolicyMakeToStock,
			DefaultLotUnits:                60,
			DemandBasis:                    augno.UpdateProductionScheduleSettingsRequestDemandBasisTrailing12,
			DemandWindowMonths:             12,
			FinishLeadTimeWeeks:            0,
			ForecastHistoryMonths:          24,
			ForecastMonths:                 12,
			ForecastZ:                      0,
			FrozenWeeks:                    1,
			GenerationTimezone:             "UTC",
			HoldingRatePct:                 0,
			HoursPerShift:                  7,
			MaxFlowDepth:                   10,
			MaxWeeksSupply:                 12,
			PlanningHorizonWeeks:           13,
			ServiceLevelZ:                  0,
			ShiftsPerDay:                   2,
			WeekStartDay:                   1,
			WeeksPerYear:                   52,
			WorkDaysPerWeek:                5,
			ConstraintDepartmentID:         augno.String("constraint_department_id"),
			GenerationCron:                 augno.String("generation_cron"),
			ReceiveCalendarID:              augno.String("receive_calendar_id"),
			ShipCalendarID:                 augno.String("ship_calendar_id"),
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

func TestOperationProductionScheduleSettingList(t *testing.T) {
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
	_, err := client.Operations.ProductionScheduleSettings.List(context.TODO())
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
