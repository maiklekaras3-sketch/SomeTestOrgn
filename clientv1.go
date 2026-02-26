// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest

import (
	"github.com/SomeTestOrg/stnlss_test-go/option"
)

// ClientV1Service contains methods and other services that help with interacting
// with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1Service] method instead.
type ClientV1Service struct {
	Options  []option.RequestOption
	Accounts ClientV1AccountService
	Cdn      ClientV1CdnService
	Zones    ClientV1ZoneService
}

// NewClientV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClientV1Service(opts ...option.RequestOption) (r ClientV1Service) {
	r = ClientV1Service{}
	r.Options = opts
	r.Accounts = NewClientV1AccountService(opts...)
	r.Cdn = NewClientV1CdnService(opts...)
	r.Zones = NewClientV1ZoneService(opts...)
	return
}
