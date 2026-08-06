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
	"github.com/augno/augno-go/packages/param"
)

func TestCatalogItemCategoryActionBulkUpsert(t *testing.T) {
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
	_, err := client.Catalog.ItemCategories.Actions.BulkUpsert(context.TODO(), augno.CatalogItemCategoryActionBulkUpsertParams{
		BulkUpsertItemCategoriesRequest: augno.BulkUpsertItemCategoriesRequestParam{
			ItemCategories: []augno.UpsertItemCategoryInputParam{{
				Name:          "Electronics",
				Notes:         param.Null[string](),
				PropertyNames: []string{"string"},
				Type:          augno.UpsertItemCategoryInputTypeMaterialCategory,
				UnitGroup: augno.ObjectIdentifierParam{
					ID:   "ug_andst6m79n41",
					Name: "name",
				},
			}},
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
