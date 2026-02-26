// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest

import (
	"github.com/SomeTestOrg/stnlss_test-go/option"
)

// ClientV1CdnService contains methods and other services that help with
// interacting with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1CdnService] method instead.
type ClientV1CdnService struct {
	Options []option.RequestOption
	Storage ClientV1CdnStorageService
	User    ClientV1CdnUserService
}

// NewClientV1CdnService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClientV1CdnService(opts ...option.RequestOption) (r ClientV1CdnService) {
	r = ClientV1CdnService{}
	r.Options = opts
	r.Storage = NewClientV1CdnStorageService(opts...)
	r.User = NewClientV1CdnUserService(opts...)
	return
}
