// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest_test

import (
	"context"
	"os"
	"testing"

	"github.com/SomeTestOrg/stnlss_test-go"
	"github.com/SomeTestOrg/stnlss_test-go/internal/testutil"
	"github.com/SomeTestOrg/stnlss_test-go/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Mock server tests are disabled")
	accessRule, err := client.Client.V1.Accounts.AccountID.Firewall.AccessRules.New(context.TODO(), stnlsstest.ClientV1AccountAccountIDFirewallAccessRuleNewParams{
		Configuration: stnlsstest.ClientV1AccountAccountIDFirewallAccessRuleNewParamsConfiguration{
			Target: "ip",
			Value:  "value",
		},
		Mode:       stnlsstest.ClientV1AccountAccountIDFirewallAccessRuleNewParamsModeChallenge,
		XAuthEmail: "X-Auth-Email",
		XAuthKey:   "X-Auth-Key",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", accessRule)
}
