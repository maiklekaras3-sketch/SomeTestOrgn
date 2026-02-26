// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest

import (
	"github.com/stainless-sdks/stnlss_test-go/option"
)

// ClientV1AccountService contains methods and other services that help with
// interacting with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1AccountService] method instead.
type ClientV1AccountService struct {
	Options   []option.RequestOption
	AccountID ClientV1AccountAccountIDService
}

// NewClientV1AccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClientV1AccountService(opts ...option.RequestOption) (r ClientV1AccountService) {
	r = ClientV1AccountService{}
	r.Options = opts
	r.AccountID = NewClientV1AccountAccountIDService(opts...)
	return
}
