// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/maiklekaras3-sketch/SomeTestOrgn/internal/requestconfig"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/option"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/packages/param"
)

// ClientV1CdnStorageUserService contains methods and other services that help with
// interacting with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1CdnStorageUserService] method instead.
type ClientV1CdnStorageUserService struct {
	Options []option.RequestOption
}

// NewClientV1CdnStorageUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClientV1CdnStorageUserService(opts ...option.RequestOption) (r ClientV1CdnStorageUserService) {
	r = ClientV1CdnStorageUserService{}
	r.Options = opts
	return
}

// Provide sftp credentials for connection.
func (r *ClientV1CdnStorageUserService) Get(ctx context.Context, storageID string, query ClientV1CdnStorageUserGetParams, opts ...option.RequestOption) (res *CdnUserResponse, err error) {
	if !param.IsOmitted(query.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", query.XAuthEmail)))
	}
	if !param.IsOmitted(query.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", query.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	if storageID == "" {
		err = errors.New("missing required storage_id parameter")
		return
	}
	path := fmt.Sprintf("client/v1/cdn/storage/%s/user", url.PathEscape(storageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ClientV1CdnStorageUserGetParams struct {
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	paramObj
}
