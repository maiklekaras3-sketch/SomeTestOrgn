// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/SomeTestOrg/stnlss_test-go/internal/apijson"
	"github.com/SomeTestOrg/stnlss_test-go/internal/requestconfig"
	"github.com/SomeTestOrg/stnlss_test-go/option"
	"github.com/SomeTestOrg/stnlss_test-go/packages/param"
	"github.com/SomeTestOrg/stnlss_test-go/packages/respjson"
)

// ClientV1ZoneZoneIDService contains methods and other services that help with
// interacting with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1ZoneZoneIDService] method instead.
type ClientV1ZoneZoneIDService struct {
	Options      []option.RequestOption
	AccessRules  ClientV1ZoneZoneIDAccessRuleService
	DNSRecords   ClientV1ZoneZoneIDDNSRecordService
	Settings     ClientV1ZoneZoneIDSettingService
	Subscription ClientV1ZoneZoneIDSubscriptionService
}

// NewClientV1ZoneZoneIDService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClientV1ZoneZoneIDService(opts ...option.RequestOption) (r ClientV1ZoneZoneIDService) {
	r = ClientV1ZoneZoneIDService{}
	r.Options = opts
	r.AccessRules = NewClientV1ZoneZoneIDAccessRuleService(opts...)
	r.DNSRecords = NewClientV1ZoneZoneIDDNSRecordService(opts...)
	r.Settings = NewClientV1ZoneZoneIDSettingService(opts...)
	r.Subscription = NewClientV1ZoneZoneIDSubscriptionService(opts...)
	return
}

// Zone Details
func (r *ClientV1ZoneZoneIDService) Get(ctx context.Context, query ClientV1ZoneZoneIDGetParams, opts ...option.RequestOption) (res *ZoneAPIResponse, err error) {
	if !param.IsOmitted(query.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", query.XAuthEmail)))
	}
	if !param.IsOmitted(query.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", query.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/zones/:zoneId"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an existing zone.
func (r *ClientV1ZoneZoneIDService) Delete(ctx context.Context, body ClientV1ZoneZoneIDDeleteParams, opts ...option.RequestOption) (res *ClientV1ZoneZoneIDDeleteResponse, err error) {
	if !param.IsOmitted(body.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", body.XAuthEmail)))
	}
	if !param.IsOmitted(body.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", body.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/zones/:zoneId"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ClientV1ZoneZoneIDDeleteResponse struct {
	// Errors contains a list of errors that occurred during the request processing. If
	// the request was successful, this will be empty.
	Errors []ResponseInfo `json:"errors"`
	// Messages contains informational messages regarding the request's outcome. This
	// will be empty if no messages were generated.
	Messages []ResponseInfo `json:"messages"`
	// Zone identifier object returned in the response.
	Result ClientV1ZoneZoneIDDeleteResponseResult `json:"result"`
	// Success indicates whether the request was successful.
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors      respjson.Field
		Messages    respjson.Field
		Result      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientV1ZoneZoneIDDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientV1ZoneZoneIDDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Zone identifier object returned in the response.
type ClientV1ZoneZoneIDDeleteResponseResult struct {
	// Zone id Unique identifier for the zone.
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientV1ZoneZoneIDDeleteResponseResult) RawJSON() string { return r.JSON.raw }
func (r *ClientV1ZoneZoneIDDeleteResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1ZoneZoneIDGetParams struct {
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	paramObj
}

type ClientV1ZoneZoneIDDeleteParams struct {
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	paramObj
}
