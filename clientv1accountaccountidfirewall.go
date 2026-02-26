// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest

import (
	"github.com/SomeTestOrg/stnlss_test-go/option"
)

// ClientV1AccountAccountIDFirewallService contains methods and other services that
// help with interacting with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1AccountAccountIDFirewallService] method instead.
type ClientV1AccountAccountIDFirewallService struct {
	Options     []option.RequestOption
	AccessRules ClientV1AccountAccountIDFirewallAccessRuleService
}

// NewClientV1AccountAccountIDFirewallService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewClientV1AccountAccountIDFirewallService(opts ...option.RequestOption) (r ClientV1AccountAccountIDFirewallService) {
	r = ClientV1AccountAccountIDFirewallService{}
	r.Options = opts
	r.AccessRules = NewClientV1AccountAccountIDFirewallAccessRuleService(opts...)
	return
}
