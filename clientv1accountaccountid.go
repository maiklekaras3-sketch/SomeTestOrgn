// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest

import (
	"github.com/SomeTestOrg/stnlss_test-go/option"
)

// ClientV1AccountAccountIDService contains methods and other services that help
// with interacting with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1AccountAccountIDService] method instead.
type ClientV1AccountAccountIDService struct {
	Options  []option.RequestOption
	Firewall ClientV1AccountAccountIDFirewallService
}

// NewClientV1AccountAccountIDService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewClientV1AccountAccountIDService(opts ...option.RequestOption) (r ClientV1AccountAccountIDService) {
	r = ClientV1AccountAccountIDService{}
	r.Options = opts
	r.Firewall = NewClientV1AccountAccountIDFirewallService(opts...)
	return
}
