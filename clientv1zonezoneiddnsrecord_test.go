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

func TestClientV1ZoneZoneIDDNSRecordListWithOptionalParams(t *testing.T) {
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
	_, err := client.Client.V1.Zones.ZoneID.DNSRecords.List(context.TODO(), stnlsstest.ClientV1ZoneZoneIDDNSRecordListParams{
		XAuthEmail: "X-Auth-Email",
		XAuthKey:   "X-Auth-Key",
		Match:      stnlsstest.String("match"),
		Page:       stnlsstest.Int(1),
		PerPage:    stnlsstest.Int(5),
	})
	if err != nil {
		var apierr *stnlsstest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestClientV1ZoneZoneIDDNSRecordBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Client.V1.Zones.ZoneID.DNSRecords.Batch(context.TODO(), stnlsstest.ClientV1ZoneZoneIDDNSRecordBatchParams{
		XAuthEmail: "X-Auth-Email",
		XAuthKey:   "X-Auth-Key",
		Deletes: []stnlsstest.ClientV1ZoneZoneIDDNSRecordBatchParamsDelete{{
			ID: stnlsstest.String("00003f34-28bf-4cfe-a795-e7d64664fc0e"),
		}},
		Patches: []stnlsstest.ClientV1ZoneZoneIDDNSRecordBatchParamsPatch{{
			ID:      stnlsstest.String("00003f34-28bf-4cfe-a795-e7d64664fc0e"),
			Comment: stnlsstest.String("This is a DNS record comment"),
			Content: stnlsstest.String("192.168.1.1"),
			Name:    stnlsstest.String("example.com"),
			Proxied: stnlsstest.Bool(true),
			Ttl:     stnlsstest.Int(3600),
			Type:    stnlsstest.String("A"),
		}},
		Posts: []stnlsstest.ClientV1ZoneZoneIDDNSRecordBatchParamsPost{{
			ID:      stnlsstest.String("00003f34-28bf-4cfe-a795-e7d64664fc0e"),
			Comment: stnlsstest.String("This is a DNS record comment"),
			Content: stnlsstest.String("192.168.1.1"),
			Name:    stnlsstest.String("example.com"),
			Proxied: stnlsstest.Bool(true),
			Ttl:     stnlsstest.Int(3600),
			Type:    stnlsstest.String("A"),
		}},
		Puts: []stnlsstest.ClientV1ZoneZoneIDDNSRecordBatchParamsPut{{
			ID:      stnlsstest.String("00003f34-28bf-4cfe-a795-e7d64664fc0e"),
			Comment: stnlsstest.String("This is a DNS record comment"),
			Content: stnlsstest.String("192.168.1.1"),
			Name:    stnlsstest.String("example.com"),
			Proxied: stnlsstest.Bool(true),
			Ttl:     stnlsstest.Int(3600),
			Type:    stnlsstest.String("A"),
		}},
	})
	if err != nil {
		var apierr *stnlsstest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
