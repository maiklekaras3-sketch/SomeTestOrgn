// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/maiklekaras3-sketch/SomeTestOrgn/internal/apijson"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/internal/requestconfig"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/option"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/packages/param"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/packages/respjson"
)

// ClientV1ZoneZoneIDSubscriptionService contains methods and other services that
// help with interacting with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1ZoneZoneIDSubscriptionService] method instead.
type ClientV1ZoneZoneIDSubscriptionService struct {
	Options []option.RequestOption
}

// NewClientV1ZoneZoneIDSubscriptionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewClientV1ZoneZoneIDSubscriptionService(opts ...option.RequestOption) (r ClientV1ZoneZoneIDSubscriptionService) {
	r = ClientV1ZoneZoneIDSubscriptionService{}
	r.Options = opts
	return
}

// Create a zone subscription, either plan or add-ons.
func (r *ClientV1ZoneZoneIDSubscriptionService) New(ctx context.Context, params ClientV1ZoneZoneIDSubscriptionNewParams, opts ...option.RequestOption) (res *BaseAPIResponse, err error) {
	if !param.IsOmitted(params.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", params.XAuthEmail)))
	}
	if !param.IsOmitted(params.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", params.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/zones/:zoneId/subscription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Lists zone subscription details.
func (r *ClientV1ZoneZoneIDSubscriptionService) Get(ctx context.Context, query ClientV1ZoneZoneIDSubscriptionGetParams, opts ...option.RequestOption) (res *ClientV1ZoneZoneIDSubscriptionGetResponse, err error) {
	if !param.IsOmitted(query.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", query.XAuthEmail)))
	}
	if !param.IsOmitted(query.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", query.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/zones/:zoneId/subscription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates zone subscriptions, either plan or add-ons.
func (r *ClientV1ZoneZoneIDSubscriptionService) Update(ctx context.Context, params ClientV1ZoneZoneIDSubscriptionUpdateParams, opts ...option.RequestOption) (res *BaseAPIResponse, err error) {
	if !param.IsOmitted(params.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", params.XAuthEmail)))
	}
	if !param.IsOmitted(params.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", params.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/zones/:zoneId/subscription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type BaseAPIResponse struct {
	// Errors contains a list of errors that occurred during the request processing. If
	// the request was successful, this will be empty.
	Errors []ResponseInfo `json:"errors"`
	// Messages contains informational messages regarding the request's outcome. This
	// will be empty if no messages were generated.
	Messages []ResponseInfo `json:"messages"`
	// Success indicates whether the request was successful.
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors      respjson.Field
		Messages    respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseAPIResponse) RawJSON() string { return r.JSON.raw }
func (r *BaseAPIResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1ZoneZoneIDSubscriptionGetResponse struct {
	// Errors contains a list of errors that occurred during the request processing. If
	// the request was successful, this will be empty.
	Errors []ResponseInfo `json:"errors"`
	// Messages contains informational messages regarding the request's outcome. This
	// will be empty if no messages were generated.
	Messages []ResponseInfo `json:"messages"`
	// Result Holds the result of the subscription operation, which includes the rate
	// plan for the subscription.
	Result ClientV1ZoneZoneIDSubscriptionGetResponseResult `json:"result"`
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
func (r ClientV1ZoneZoneIDSubscriptionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientV1ZoneZoneIDSubscriptionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result Holds the result of the subscription operation, which includes the rate
// plan for the subscription.
type ClientV1ZoneZoneIDSubscriptionGetResponseResult struct {
	// Rate plan associated with the subscription The rate plan chosen for the
	// subscription.
	RatePlan ClientV1ZoneZoneIDSubscriptionGetResponseResultRatePlan `json:"rate_plan"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RatePlan    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientV1ZoneZoneIDSubscriptionGetResponseResult) RawJSON() string { return r.JSON.raw }
func (r *ClientV1ZoneZoneIDSubscriptionGetResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rate plan associated with the subscription The rate plan chosen for the
// subscription.
type ClientV1ZoneZoneIDSubscriptionGetResponseResultRatePlan struct {
	// Rate plan ID Unique identifier for the rate plan.
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientV1ZoneZoneIDSubscriptionGetResponseResultRatePlan) RawJSON() string { return r.JSON.raw }
func (r *ClientV1ZoneZoneIDSubscriptionGetResponseResultRatePlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1ZoneZoneIDSubscriptionNewParams struct {
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	// Plan ID The unique identifier for the rate plan to be requested.
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r ClientV1ZoneZoneIDSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ClientV1ZoneZoneIDSubscriptionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientV1ZoneZoneIDSubscriptionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1ZoneZoneIDSubscriptionGetParams struct {
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	paramObj
}

type ClientV1ZoneZoneIDSubscriptionUpdateParams struct {
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	// Plan ID The unique identifier for the rate plan to be requested.
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r ClientV1ZoneZoneIDSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ClientV1ZoneZoneIDSubscriptionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientV1ZoneZoneIDSubscriptionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
