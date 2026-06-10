// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package augno_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/augno/augno-go"
	"github.com/augno/augno-go/internal"
	"github.com/augno/augno-go/option"
)

type closureTransport struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (t *closureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.fn(req)
}

func TestUserAgentHeader(t *testing.T) {
	var userAgent string
	client := augno.NewClient(
		option.WithBearerToken("My Bearer Token"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					userAgent = req.Header.Get("User-Agent")
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
			},
		}),
	)
	_, _ = client.Catalog.Items.List(context.Background(), augno.CatalogItemListParams{})
	if userAgent != fmt.Sprintf("Augno/Go %s", internal.PackageVersion) {
		t.Errorf("Expected User-Agent to be correct, but got: %#v", userAgent)
	}
}

func TestRetryAfter(t *testing.T) {
	attempts := 0
	client := augno.NewClient(
		option.WithBearerToken("My Bearer Token"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					attempts += 1
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
	)
	_, err := client.Catalog.Items.List(context.Background(), augno.CatalogItemListParams{})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	if attempts != 3 {
		t.Errorf("Expected %d attempts, got %d", 3, attempts)
	}
}

func TestRetryAfterMs(t *testing.T) {
	attempts := 0
	client := augno.NewClient(
		option.WithBearerToken("My Bearer Token"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					attempts++
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After-Ms"): []string{"100"},
						},
					}, nil
				},
			},
		}),
	)
	_, err := client.Catalog.Items.List(context.Background(), augno.CatalogItemListParams{})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}
	if want := 3; attempts != want {
		t.Errorf("Expected %d attempts, got %d", want, attempts)
	}
}

func TestContextCancel(t *testing.T) {
	client := augno.NewClient(
		option.WithBearerToken("My Bearer Token"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := client.Catalog.Items.List(cancelCtx, augno.CatalogItemListParams{})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}
}

func TestContextCancelDelay(t *testing.T) {
	client := augno.NewClient(
		option.WithBearerToken("My Bearer Token"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer cancel()
	_, err := client.Catalog.Items.List(cancelCtx, augno.CatalogItemListParams{})
	if err == nil {
		t.Error("expected there to be a cancel error")
	}
}

func TestContextDeadline(t *testing.T) {
	testTimeout := time.After(3 * time.Second)
	testDone := make(chan struct{})

	deadline := time.Now().Add(100 * time.Millisecond)
	deadlineCtx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		client := augno.NewClient(
			option.WithBearerToken("My Bearer Token"),
			option.WithHTTPClient(&http.Client{
				Transport: &closureTransport{
					fn: func(req *http.Request) (*http.Response, error) {
						<-req.Context().Done()
						return nil, req.Context().Err()
					},
				},
			}),
		)
		_, err := client.Catalog.Items.List(deadlineCtx, augno.CatalogItemListParams{})
		if err == nil {
			t.Error("expected there to be a deadline error")
		}
		close(testDone)
	}()

	select {
	case <-testTimeout:
		t.Fatal("client didn't finish in time")
	case <-testDone:
		if diff := time.Since(deadline); diff < -30*time.Millisecond || 30*time.Millisecond < diff {
			t.Fatalf("client did not return within 30ms of context deadline, got %s", diff)
		}
	}
}
