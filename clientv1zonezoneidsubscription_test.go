// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/stnlss_test-go"
	"github.com/stainless-sdks/stnlss_test-go/internal/testutil"
	"github.com/stainless-sdks/stnlss_test-go/option"
)

func TestClientV1ZoneZoneIDSubscriptionNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := stnlsstest.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Client.V1.Zones.ZoneID.Subscription.New(context.TODO(), stnlsstest.ClientV1ZoneZoneIDSubscriptionNewParams{
		XAuthEmail: "X-Auth-Email",
		XAuthKey:   "X-Auth-Key",
		ID:         stnlsstest.String("enterprise"),
	})
	if err != nil {
		var apierr *stnlsstest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestClientV1ZoneZoneIDSubscriptionGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := stnlsstest.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Client.V1.Zones.ZoneID.Subscription.Get(context.TODO(), stnlsstest.ClientV1ZoneZoneIDSubscriptionGetParams{
		XAuthEmail: "X-Auth-Email",
		XAuthKey:   "X-Auth-Key",
	})
	if err != nil {
		var apierr *stnlsstest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestClientV1ZoneZoneIDSubscriptionUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := stnlsstest.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Client.V1.Zones.ZoneID.Subscription.Update(context.TODO(), stnlsstest.ClientV1ZoneZoneIDSubscriptionUpdateParams{
		XAuthEmail: "X-Auth-Email",
		XAuthKey:   "X-Auth-Key",
		ID:         stnlsstest.String("enterprise"),
	})
	if err != nil {
		var apierr *stnlsstest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
