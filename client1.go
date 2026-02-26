// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest

import (
	"github.com/maiklekaras3-sketch/SomeTestOrgn/option"
)

// ClientService contains methods and other services that help with interacting
// with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientService] method instead.
type ClientService struct {
	Options []option.RequestOption
	V1      ClientV1Service
}

// NewClientService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewClientService(opts ...option.RequestOption) (r ClientService) {
	r = ClientService{}
	r.Options = opts
	r.V1 = NewClientV1Service(opts...)
	return
}
