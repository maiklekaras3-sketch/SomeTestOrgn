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

// ClientV1ZoneZoneIDSettingSettingIDService contains methods and other services
// that help with interacting with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1ZoneZoneIDSettingSettingIDService] method instead.
type ClientV1ZoneZoneIDSettingSettingIDService struct {
	Options []option.RequestOption
}

// NewClientV1ZoneZoneIDSettingSettingIDService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewClientV1ZoneZoneIDSettingSettingIDService(opts ...option.RequestOption) (r ClientV1ZoneZoneIDSettingSettingIDService) {
	r = ClientV1ZoneZoneIDSettingSettingIDService{}
	r.Options = opts
	return
}

// Fetch a single zone setting by name
func (r *ClientV1ZoneZoneIDSettingSettingIDService) Get(ctx context.Context, query ClientV1ZoneZoneIDSettingSettingIDGetParams, opts ...option.RequestOption) (res *ZoneSettingAPIResponse, err error) {
	if !param.IsOmitted(query.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", query.XAuthEmail)))
	}
	if !param.IsOmitted(query.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", query.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/zones/:zoneId/settings/:settingId"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a single zone setting by the identifier
func (r *ClientV1ZoneZoneIDSettingSettingIDService) Update(ctx context.Context, params ClientV1ZoneZoneIDSettingSettingIDUpdateParams, opts ...option.RequestOption) (res *ZoneSettingAPIResponse, err error) {
	if !param.IsOmitted(params.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", params.XAuthEmail)))
	}
	if !param.IsOmitted(params.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", params.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/zones/:zoneId/settings/:settingId"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type ZoneSettingAPIResponse struct {
	// Errors contains a list of errors that occurred during the request processing. If
	// the request was successful, this will be empty.
	Errors []ResponseInfo `json:"errors"`
	// Messages contains informational messages regarding the request's outcome. This
	// will be empty if no messages were generated.
	Messages []ResponseInfo `json:"messages"`
	// Result Holds the result of the zone setting operation, which includes the zone
	// setting ID and its value.
	Result ZoneSettingAPIResponseResult `json:"result"`
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
func (r ZoneSettingAPIResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneSettingAPIResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result Holds the result of the zone setting operation, which includes the zone
// setting ID and its value.
type ZoneSettingAPIResponseResult struct {
	// Zone setting ID Unique identifier for the zone setting.
	ID string `json:"id"`
	// Setting value The current value of the zone setting.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneSettingAPIResponseResult) RawJSON() string { return r.JSON.raw }
func (r *ZoneSettingAPIResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1ZoneZoneIDSettingSettingIDGetParams struct {
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	paramObj
}

type ClientV1ZoneZoneIDSettingSettingIDUpdateParams struct {
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	// Setting value The new value for the zone setting. ID of zone setting will be
	// taken from url path
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r ClientV1ZoneZoneIDSettingSettingIDUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ClientV1ZoneZoneIDSettingSettingIDUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientV1ZoneZoneIDSettingSettingIDUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
