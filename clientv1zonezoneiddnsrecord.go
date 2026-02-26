// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/maiklekaras3-sketch/SomeTestOrgn/internal/apijson"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/internal/apiquery"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/internal/requestconfig"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/option"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/packages/param"
	"github.com/maiklekaras3-sketch/SomeTestOrgn/packages/respjson"
)

// ClientV1ZoneZoneIDDNSRecordService contains methods and other services that help
// with interacting with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1ZoneZoneIDDNSRecordService] method instead.
type ClientV1ZoneZoneIDDNSRecordService struct {
	Options []option.RequestOption
}

// NewClientV1ZoneZoneIDDNSRecordService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewClientV1ZoneZoneIDDNSRecordService(opts ...option.RequestOption) (r ClientV1ZoneZoneIDDNSRecordService) {
	r = ClientV1ZoneZoneIDDNSRecordService{}
	r.Options = opts
	return
}

// List, search, sort, and filter a zones' DNS records.
func (r *ClientV1ZoneZoneIDDNSRecordService) List(ctx context.Context, params ClientV1ZoneZoneIDDNSRecordListParams, opts ...option.RequestOption) (res *ClientV1ZoneZoneIddnsRecordListResponse, err error) {
	if !param.IsOmitted(params.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", params.XAuthEmail)))
	}
	if !param.IsOmitted(params.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", params.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/zones/:zoneId/dns_records"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Send a Batch of DNS Record API calls to be executed together.
func (r *ClientV1ZoneZoneIDDNSRecordService) Batch(ctx context.Context, params ClientV1ZoneZoneIDDNSRecordBatchParams, opts ...option.RequestOption) (res *ClientV1ZoneZoneIddnsRecordBatchResponse, err error) {
	if !param.IsOmitted(params.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", params.XAuthEmail)))
	}
	if !param.IsOmitted(params.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", params.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/zones/:zoneId/dns_records/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ZoneDNSRecord struct {
	// Record ID Unique identifier for the DNS record.
	ID string `json:"id"`
	// Notes about the record A comment or note associated with the DNS record.
	Comment string `json:"comment"`
	// DNS Record content based on the record type The content of the DNS record,
	// varying based on the record type.
	Content string `json:"content"`
	// Record name The name associated with the DNS record, typically representing a
	// domain or subdomain.
	Name string `json:"name"`
	// Whether proxied through our system A flag indicating if the DNS record is being
	// proxied through our system.
	Proxied bool `json:"proxied"`
	// Time to live The TTL (Time To Live) in seconds for the DNS record.
	Ttl int64 `json:"ttl"`
	// Record type The type of the DNS record (e.g., A, AAAA, CNAME, etc.).
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Comment     respjson.Field
		Content     respjson.Field
		Name        respjson.Field
		Proxied     respjson.Field
		Ttl         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneDNSRecord) RawJSON() string { return r.JSON.raw }
func (r *ZoneDNSRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1ZoneZoneIddnsRecordListResponse struct {
	// Errors contains a list of errors that occurred during the request processing. If
	// the request was successful, this will be empty.
	Errors []ResponseInfo `json:"errors"`
	// Messages contains informational messages regarding the request's outcome. This
	// will be empty if no messages were generated.
	Messages []ResponseInfo `json:"messages"`
	// Array of zone dns record objects returned in the response.
	Result []ZoneDNSRecord `json:"result"`
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
func (r ClientV1ZoneZoneIddnsRecordListResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientV1ZoneZoneIddnsRecordListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1ZoneZoneIddnsRecordBatchResponse struct {
	// Errors contains a list of errors that occurred during the request processing. If
	// the request was successful, this will be empty.
	Errors []ResponseInfo `json:"errors"`
	// Messages contains informational messages regarding the request's outcome. This
	// will be empty if no messages were generated.
	Messages []ResponseInfo `json:"messages"`
	// Result Holds the results of DNS record batch operations. It contains the
	// deletes, patches, posts, and puts that were processed.
	Result ClientV1ZoneZoneIddnsRecordBatchResponseResult `json:"result"`
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
func (r ClientV1ZoneZoneIddnsRecordBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientV1ZoneZoneIddnsRecordBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result Holds the results of DNS record batch operations. It contains the
// deletes, patches, posts, and puts that were processed.
type ClientV1ZoneZoneIddnsRecordBatchResponseResult struct {
	// Deletes contains the DNS records that were deleted in the batch operation. These
	// records are identified by their unique IDs and will no longer exist in the
	// system.
	Deletes []ZoneDNSRecord `json:"deletes"`
	// Patches contains the DNS records that were updated in the batch operation. These
	// records contain the modified fields that were updated during the operation.
	Patches []ZoneDNSRecord `json:"patches"`
	// Posts contains the DNS records that were created in the batch operation. These
	// records represent new DNS records that were added to the system.
	Posts []ZoneDNSRecord `json:"posts"`
	// Puts contains the DNS records that were fully replaced in the batch operation.
	// These records are replaced with new data, typically including all fields.
	Puts []ZoneDNSRecord `json:"puts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deletes     respjson.Field
		Patches     respjson.Field
		Posts       respjson.Field
		Puts        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientV1ZoneZoneIddnsRecordBatchResponseResult) RawJSON() string { return r.JSON.raw }
func (r *ClientV1ZoneZoneIddnsRecordBatchResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1ZoneZoneIDDNSRecordListParams struct {
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	// Search zone name with filter operator
	Match param.Opt[string] `query:"match,omitzero" json:"-"`
	// Page number of paginated results
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of zones per page
	PerPage param.Opt[int64] `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ClientV1ZoneZoneIDDNSRecordListParams]'s query parameters
// as `url.Values`.
func (r ClientV1ZoneZoneIDDNSRecordListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ClientV1ZoneZoneIDDNSRecordBatchParams struct {
	XAuthEmail string                                         `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string                                         `header:"X-Auth-Key" api:"required" json:"-"`
	Deletes    []ClientV1ZoneZoneIDDNSRecordBatchParamsDelete `json:"deletes,omitzero"`
	Patches    []ClientV1ZoneZoneIDDNSRecordBatchParamsPatch  `json:"patches,omitzero"`
	Posts      []ClientV1ZoneZoneIDDNSRecordBatchParamsPost   `json:"posts,omitzero"`
	Puts       []ClientV1ZoneZoneIDDNSRecordBatchParamsPut    `json:"puts,omitzero"`
	paramObj
}

func (r ClientV1ZoneZoneIDDNSRecordBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow ClientV1ZoneZoneIDDNSRecordBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientV1ZoneZoneIDDNSRecordBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1ZoneZoneIDDNSRecordBatchParamsDelete struct {
	// ID is the unique identifier of the DNS record to be deleted. This ID is used to
	// find and delete the specified record.
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r ClientV1ZoneZoneIDDNSRecordBatchParamsDelete) MarshalJSON() (data []byte, err error) {
	type shadow ClientV1ZoneZoneIDDNSRecordBatchParamsDelete
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientV1ZoneZoneIDDNSRecordBatchParamsDelete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1ZoneZoneIDDNSRecordBatchParamsPatch struct {
	// Record ID Unique identifier for the DNS record.
	ID param.Opt[string] `json:"id,omitzero"`
	// Notes about the record A comment or note associated with the DNS record.
	Comment param.Opt[string] `json:"comment,omitzero"`
	// DNS Record content based on the record type The content of the DNS record,
	// varying based on the record type.
	Content param.Opt[string] `json:"content,omitzero"`
	// Record name The name associated with the DNS record, typically representing a
	// domain or subdomain.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether proxied through our system A flag indicating if the DNS record is being
	// proxied through our system.
	Proxied param.Opt[bool] `json:"proxied,omitzero"`
	// Time to live The TTL (Time To Live) in seconds for the DNS record.
	Ttl param.Opt[int64] `json:"ttl,omitzero"`
	// Record type The type of the DNS record (e.g., A, AAAA, CNAME, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r ClientV1ZoneZoneIDDNSRecordBatchParamsPatch) MarshalJSON() (data []byte, err error) {
	type shadow ClientV1ZoneZoneIDDNSRecordBatchParamsPatch
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientV1ZoneZoneIDDNSRecordBatchParamsPatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1ZoneZoneIDDNSRecordBatchParamsPost struct {
	// Record ID Unique identifier for the DNS record.
	ID param.Opt[string] `json:"id,omitzero"`
	// Notes about the record A comment or note associated with the DNS record.
	Comment param.Opt[string] `json:"comment,omitzero"`
	// DNS Record content based on the record type The content of the DNS record,
	// varying based on the record type.
	Content param.Opt[string] `json:"content,omitzero"`
	// Record name The name associated with the DNS record, typically representing a
	// domain or subdomain.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether proxied through our system A flag indicating if the DNS record is being
	// proxied through our system.
	Proxied param.Opt[bool] `json:"proxied,omitzero"`
	// Time to live The TTL (Time To Live) in seconds for the DNS record.
	Ttl param.Opt[int64] `json:"ttl,omitzero"`
	// Record type The type of the DNS record (e.g., A, AAAA, CNAME, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r ClientV1ZoneZoneIDDNSRecordBatchParamsPost) MarshalJSON() (data []byte, err error) {
	type shadow ClientV1ZoneZoneIDDNSRecordBatchParamsPost
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientV1ZoneZoneIDDNSRecordBatchParamsPost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1ZoneZoneIDDNSRecordBatchParamsPut struct {
	// Record ID Unique identifier for the DNS record.
	ID param.Opt[string] `json:"id,omitzero"`
	// Notes about the record A comment or note associated with the DNS record.
	Comment param.Opt[string] `json:"comment,omitzero"`
	// DNS Record content based on the record type The content of the DNS record,
	// varying based on the record type.
	Content param.Opt[string] `json:"content,omitzero"`
	// Record name The name associated with the DNS record, typically representing a
	// domain or subdomain.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether proxied through our system A flag indicating if the DNS record is being
	// proxied through our system.
	Proxied param.Opt[bool] `json:"proxied,omitzero"`
	// Time to live The TTL (Time To Live) in seconds for the DNS record.
	Ttl param.Opt[int64] `json:"ttl,omitzero"`
	// Record type The type of the DNS record (e.g., A, AAAA, CNAME, etc.).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r ClientV1ZoneZoneIDDNSRecordBatchParamsPut) MarshalJSON() (data []byte, err error) {
	type shadow ClientV1ZoneZoneIDDNSRecordBatchParamsPut
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientV1ZoneZoneIDDNSRecordBatchParamsPut) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
