// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/maiklekaras3-sketch/SomeTestOrgn/internal/apijson"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/internal/requestconfig"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/option"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/packages/param"
)

// ClientV1CdnUserService contains methods and other services that help with
// interacting with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1CdnUserService] method instead.
type ClientV1CdnUserService struct {
	Options []option.RequestOption
}

// NewClientV1CdnUserService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClientV1CdnUserService(opts ...option.RequestOption) (r ClientV1CdnUserService) {
	r = ClientV1CdnUserService{}
	r.Options = opts
	return
}

// Create new sftp credentials if they don't exist.
func (r *ClientV1CdnUserService) New(ctx context.Context, params ClientV1CdnUserNewParams, opts ...option.RequestOption) (res *CdnUserResponse, err error) {
	if !param.IsOmitted(params.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", params.XAuthEmail)))
	}
	if !param.IsOmitted(params.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", params.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/cdn/user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Revoke old sftp credentials and simultaneously provides new credentials.
func (r *ClientV1CdnUserService) Revoke(ctx context.Context, username string, body ClientV1CdnUserRevokeParams, opts ...option.RequestOption) (res *CdnUserResponse, err error) {
	if !param.IsOmitted(body.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", body.XAuthEmail)))
	}
	if !param.IsOmitted(body.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", body.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	if username == "" {
		err = errors.New("missing required username parameter")
		return
	}
	path := fmt.Sprintf("client/v1/cdn/user/%s/revoke", url.PathEscape(username))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type ClientV1CdnUserNewParams struct {
	StorageID  string `json:"storage_id" api:"required"`
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	paramObj
}

func (r ClientV1CdnUserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ClientV1CdnUserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientV1CdnUserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1CdnUserRevokeParams struct {
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	paramObj
}
