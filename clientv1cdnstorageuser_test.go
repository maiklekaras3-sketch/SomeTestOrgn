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

func TestClientV1CdnStorageUserGet(t *testing.T) {
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
	_, err := client.Client.V1.Cdn.Storage.User.Get(
		context.TODO(),
		"storage_id",
		stnlsstest.ClientV1CdnStorageUserGetParams{
			XAuthEmail: "X-Auth-Email",
			XAuthKey:   "X-Auth-Key",
		},
	)
	if err != nil {
		var apierr *stnlsstest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
