// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/SomeTestOrg/stnlss_test-go/internal/apijson"
	"github.com/SomeTestOrg/stnlss_test-go/internal/apiquery"
	"github.com/SomeTestOrg/stnlss_test-go/internal/requestconfig"
	"github.com/SomeTestOrg/stnlss_test-go/option"
	"github.com/SomeTestOrg/stnlss_test-go/packages/param"
	"github.com/SomeTestOrg/stnlss_test-go/packages/respjson"
)

// ClientV1ZoneService contains methods and other services that help with
// interacting with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1ZoneService] method instead.
type ClientV1ZoneService struct {
	Options []option.RequestOption
	ZoneID  ClientV1ZoneZoneIDService
}

// NewClientV1ZoneService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClientV1ZoneService(opts ...option.RequestOption) (r ClientV1ZoneService) {
	r = ClientV1ZoneService{}
	r.Options = opts
	r.ZoneID = NewClientV1ZoneZoneIDService(opts...)
	return
}

// Create Zone
func (r *ClientV1ZoneService) New(ctx context.Context, params ClientV1ZoneNewParams, opts ...option.RequestOption) (res *ZoneAPIResponse, err error) {
	if !param.IsOmitted(params.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", params.XAuthEmail)))
	}
	if !param.IsOmitted(params.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", params.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Lists, searches, sorts, and filters your zones. Listing zones across more than
// 500 accounts is currently not allowed.
func (r *ClientV1ZoneService) List(ctx context.Context, params ClientV1ZoneListParams, opts ...option.RequestOption) (res *ClientV1ZoneListResponse, err error) {
	if !param.IsOmitted(params.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", params.XAuthEmail)))
	}
	if !param.IsOmitted(params.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", params.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type ResponseInfo struct {
	// Numeric code representing the message or error.
	Code int64 `json:"code"`
	// Human-readable message describing the response.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseInfo) RawJSON() string { return r.JSON.raw }
func (r *ResponseInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseZone struct {
	// Unique identifier for the zone.
	ID string `json:"id"`
	// Associated account information.
	Account ResponseZoneAccount `json:"account"`
	// Group identifier for the zone
	GroupID int64 `json:"group_id"`
	// Domain name of the zone.
	Name string `json:"name"`
	// Nameservers associated with the zone
	NameServers []string `json:"name_servers"`
	// Original Nameservers given you by DNS provider
	OriginalNameServers []string `json:"original_name_servers"`
	// Zone owner details
	Owner ResponseZoneOwner `json:"owner"`
	// Zone type
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Account             respjson.Field
		GroupID             respjson.Field
		Name                respjson.Field
		NameServers         respjson.Field
		OriginalNameServers respjson.Field
		Owner               respjson.Field
		Type                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseZone) RawJSON() string { return r.JSON.raw }
func (r *ResponseZone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Associated account information.
type ResponseZoneAccount struct {
	// ID is the unique identifier for the account. This ID is used to reference the
	// account in various operations.
	ID string `json:"id"`
	// Name is the human-readable name of the account. This name is used to identify
	// the account in the user interface or in responses.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseZoneAccount) RawJSON() string { return r.JSON.raw }
func (r *ResponseZoneAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Zone owner details
type ResponseZoneOwner struct {
	// ID is the unique identifier for the owner. This ID is used to reference the
	// owner in various operations and records.
	ID string `json:"id"`
	// Name is the human-readable name of the owner. It identifies the owner, which
	// could be either an organization or an individual.
	Name string `json:"name"`
	// Type indicates the type of the owner. Possible values are "organization" or
	// "individual", denoting whether the owner is a company or a person.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseZoneOwner) RawJSON() string { return r.JSON.raw }
func (r *ResponseZoneOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResultInfo struct {
	// Number of items returned in the current page.
	Count int64 `json:"count"`
	// Current page number.
	Page int64 `json:"page"`
	// Number of items requested per page.
	PerPage int64 `json:"per_page"`
	// Total number of items available.
	TotalCount int64 `json:"total_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Page        respjson.Field
		PerPage     respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResultInfo) RawJSON() string { return r.JSON.raw }
func (r *ResultInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAPIResponse struct {
	// Errors contains a list of errors that occurred during the request processing. If
	// the request was successful, this will be empty.
	Errors []ResponseInfo `json:"errors"`
	// Messages contains informational messages regarding the request's outcome. This
	// will be empty if no messages were generated.
	Messages []ResponseInfo `json:"messages"`
	// Zone object returned in the response.
	Result ResponseZone `json:"result"`
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
func (r ZoneAPIResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneAPIResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1ZoneListResponse struct {
	// Errors contains a list of errors that occurred during the request processing. If
	// the request was successful, this will be empty.
	Errors []ResponseInfo `json:"errors"`
	// Messages contains informational messages regarding the request's outcome. This
	// will be empty if no messages were generated.
	Messages []ResponseInfo `json:"messages"`
	// Array of zone objects returned in the response.
	Result []ResponseZone `json:"result"`
	// ResultInfo contains pagination details if applicable, such as the number of
	// results per page, current page, and total count.
	ResultInfo ResultInfo `json:"result_info"`
	// Success indicates whether the request was successful.
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors      respjson.Field
		Messages    respjson.Field
		Result      respjson.Field
		ResultInfo  respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientV1ZoneListResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientV1ZoneListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1ZoneNewParams struct {
	// Account represents the account information associated with the zone.
	Account ClientV1ZoneNewParamsAccount `json:"account,omitzero" api:"required"`
	// Name represents the domain name of the zone. Maximum length: 253 characters.
	Name       string `json:"name" api:"required"`
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	// DomainGroupId represents domain group which zone will be added
	GroupID        param.Opt[int64]  `json:"group_id,omitzero"`
	SourceZoneID   param.Opt[string] `json:"source_zone_id,omitzero"`
	DNSFileContent []int64           `json:"dns_file_content,omitzero"`
	// Any of "auto", "file", "manual", "source".
	DNSMethod ClientV1ZoneNewParamsDNSMethod `json:"dns_method,omitzero"`
	// Type represents the type of the zone, which can be one of: "full", "partial",
	// "secondary", or "internal".
	//
	// Any of "full", "partial", "secondary", "internal".
	Type ClientV1ZoneNewParamsType `json:"type,omitzero"`
	paramObj
}

func (r ClientV1ZoneNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ClientV1ZoneNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientV1ZoneNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account represents the account information associated with the zone.
//
// The property ID is required.
type ClientV1ZoneNewParamsAccount struct {
	ID string `json:"id" api:"required"`
	paramObj
}

func (r ClientV1ZoneNewParamsAccount) MarshalJSON() (data []byte, err error) {
	type shadow ClientV1ZoneNewParamsAccount
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientV1ZoneNewParamsAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1ZoneNewParamsDNSMethod string

const (
	ClientV1ZoneNewParamsDNSMethodAuto   ClientV1ZoneNewParamsDNSMethod = "auto"
	ClientV1ZoneNewParamsDNSMethodFile   ClientV1ZoneNewParamsDNSMethod = "file"
	ClientV1ZoneNewParamsDNSMethodManual ClientV1ZoneNewParamsDNSMethod = "manual"
	ClientV1ZoneNewParamsDNSMethodSource ClientV1ZoneNewParamsDNSMethod = "source"
)

// Type represents the type of the zone, which can be one of: "full", "partial",
// "secondary", or "internal".
type ClientV1ZoneNewParamsType string

const (
	ClientV1ZoneNewParamsTypeFull      ClientV1ZoneNewParamsType = "full"
	ClientV1ZoneNewParamsTypePartial   ClientV1ZoneNewParamsType = "partial"
	ClientV1ZoneNewParamsTypeSecondary ClientV1ZoneNewParamsType = "secondary"
	ClientV1ZoneNewParamsTypeInternal  ClientV1ZoneNewParamsType = "internal"
)

type ClientV1ZoneListParams struct {
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	// Search zone name with filter operator
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Page number of paginated results
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of zones per page
	PerPage param.Opt[int64] `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ClientV1ZoneListParams]'s query parameters as `url.Values`.
func (r ClientV1ZoneListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
