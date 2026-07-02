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

func TestCatalogPartNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Parts.New(context.TODO(), augno.CatalogPartNewParams{
		CreatePartRequest: augno.CreatePartRequestParam{
			CategoryID:   "ic_01ae7bd7bfd21ca0ab81e1357e",
			SKU:          "BRG-6204-2RS",
			AttributeIDs: []string{"at_01c9493ec0c46bb0ed12708ae4"},
			Description:  augno.String("Deep groove ball bearing, 20x47x14mm"),
			Notes:        augno.String("OEM-equivalent; verify shielding type before substitution."),
			UnitCost: augno.RateInputParam{
				DenominatorUnitID: "un_01966263f74a5a0cae356000a1",
				NumeratorUnitID:   "un_01966263f74a5a0cae356000a1",
				Value:             "9.40",
			},
			UnitPrice: augno.RateInputParam{
				DenominatorUnitID: "un_01966263f74a5a0cae356000a1",
				NumeratorUnitID:   "un_01966263f74a5a0cae356000a1",
				Value:             "14.99",
			},
		},
		Include: []string{"item"},
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogPartGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Parts.Get(
		context.TODO(),
		"pt_018d7bab53e864351f4c693a21",
		augno.CatalogPartGetParams{
			Include: []string{"item"},
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

func TestCatalogPartUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Parts.Update(
		context.TODO(),
		"pt_018d7bab53e864351f4c693a21",
		augno.CatalogPartUpdateParams{
			Include: []string{"item"},
			UpdatePartRequest: augno.UpdatePartRequestParam{
				Description: augno.String("Deep groove ball bearing, 20x47x14mm"),
				Notes:       augno.String("Superseded by low-friction variant; keep for legacy assemblies."),
				SKU:         augno.String("BRG-6204-2RS"),
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

func TestCatalogPartListWithOptionalParams(t *testing.T) {
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
	_, err := client.Catalog.Parts.List(context.TODO(), augno.CatalogPartListParams{
		AttributeIDs: []string{"string"},
		CategoryIDs:  []string{"string"},
		Cursor:       augno.String("cursor"),
		EndDate:      augno.Time(time.Now()),
		Include:      []string{"item"},
		Limit:        augno.Int(0),
		Q:            augno.String("q"),
		StartDate:    augno.Time(time.Now()),
	})
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCatalogPartDelete(t *testing.T) {
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
	_, err := client.Catalog.Parts.Delete(context.TODO(), "pt_018d7bab53e864351f4c693a21")
	if err != nil {
		var apierr *augno.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
