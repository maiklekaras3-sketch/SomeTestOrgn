// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/maiklekaras3-sketch/SomeTestOrgn"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/internal/testutil"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/option"
)

func TestClientV1ZoneNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Client.V1.Zones.New(context.TODO(), stnlsstest.ClientV1ZoneNewParams{
		Account: stnlsstest.ClientV1ZoneNewParamsAccount{
			ID: "023e105f4ecef8ad9ca31a8372d0c353",
		},
		Name:           "example.com",
		XAuthEmail:     "X-Auth-Email",
		XAuthKey:       "X-Auth-Key",
		DNSFileContent: []int64{0},
		DNSMethod:      stnlsstest.ClientV1ZoneNewParamsDNSMethodAuto,
		GroupID:        stnlsstest.Int(991),
		SourceZoneID:   stnlsstest.String("source_zone_id"),
		Type:           stnlsstest.ClientV1ZoneNewParamsTypeFull,
	})
	if err != nil {
		var apierr *stnlsstest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestClientV1ZoneListWithOptionalParams(t *testing.T) {
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
	_, err := client.Client.V1.Zones.List(context.TODO(), stnlsstest.ClientV1ZoneListParams{
		XAuthEmail: "X-Auth-Email",
		XAuthKey:   "X-Auth-Key",
		Name:       stnlsstest.String("name"),
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
