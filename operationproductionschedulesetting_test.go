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

func TestOperationProductionScheduleSettingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Operations.ProductionScheduleSettings.Update(context.TODO(), openmrp.OperationProductionScheduleSettingUpdateParams{
		UpdateProductionScheduleSettingsRequest: openmrp.UpdateProductionScheduleSettingsRequestParam{
			AutoPublishStatus:              openmrp.UpdateProductionScheduleSettingsRequestAutoPublishStatusInactive,
			CadenceStatus:                  openmrp.UpdateProductionScheduleSettingsRequestCadenceStatusInactive,
			CapacityHeadroomPct:            0.9,
			ChangeoverAvgMinutes:           0,
			ChangeoverLaborRate:            0,
			ChangeoverMaxMinutes:           0,
			ChangeoverMinMinutes:           0,
			DefaultConstraintLeadTimeWeeks: 0,
			DefaultCustomerLeadTimeDays:    30,
			DefaultFulfillmentPolicy:       openmrp.UpdateProductionScheduleSettingsRequestDefaultFulfillmentPolicyMakeToStock,
			DefaultLotUnits:                60,
			DemandBasis:                    openmrp.UpdateProductionScheduleSettingsRequestDemandBasisTrailing12,
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
			ConstraintDepartmentID:         openmrp.String("constraint_department_id"),
			GenerationCron:                 openmrp.String("generation_cron"),
			ReceiveCalendarID:              openmrp.String("receive_calendar_id"),
			ShipCalendarID:                 openmrp.String("ship_calendar_id"),
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

func TestOperationProductionScheduleSettingList(t *testing.T) {
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
	_, err := client.Operations.ProductionScheduleSettings.List(context.TODO())
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
