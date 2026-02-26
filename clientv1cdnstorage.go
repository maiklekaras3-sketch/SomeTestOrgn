// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/SomeTestOrg/stnlss_test-go/internal/apijson"
	"github.com/SomeTestOrg/stnlss_test-go/internal/requestconfig"
	"github.com/SomeTestOrg/stnlss_test-go/option"
	"github.com/SomeTestOrg/stnlss_test-go/packages/param"
	"github.com/SomeTestOrg/stnlss_test-go/packages/respjson"
)

// ClientV1CdnStorageService contains methods and other services that help with
// interacting with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1CdnStorageService] method instead.
type ClientV1CdnStorageService struct {
	Options []option.RequestOption
	User    ClientV1CdnStorageUserService
}

// NewClientV1CdnStorageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClientV1CdnStorageService(opts ...option.RequestOption) (r ClientV1CdnStorageService) {
	r = ClientV1CdnStorageService{}
	r.Options = opts
	r.User = NewClientV1CdnStorageUserService(opts...)
	return
}

// Create new CDN storage where you can store your files.
func (r *ClientV1CdnStorageService) New(ctx context.Context, params ClientV1CdnStorageNewParams, opts ...option.RequestOption) (res *CdnStorageResponse, err error) {
	if !param.IsOmitted(params.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", params.XAuthEmail)))
	}
	if !param.IsOmitted(params.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", params.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/cdn/storage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Provide full list of CDN storages which is accessible for you.
func (r *ClientV1CdnStorageService) List(ctx context.Context, query ClientV1CdnStorageListParams, opts ...option.RequestOption) (res *[]CdnStorageResponse, err error) {
	if !param.IsOmitted(query.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", query.XAuthEmail)))
	}
	if !param.IsOmitted(query.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", query.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/cdn/storage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete your CDN storage.
func (r *ClientV1CdnStorageService) Delete(ctx context.Context, storageID string, body ClientV1CdnStorageDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", body.XAuthEmail)))
	}
	if !param.IsOmitted(body.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", body.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if storageID == "" {
		err = errors.New("missing required storage_id parameter")
		return
	}
	path := fmt.Sprintf("client/v1/cdn/storage/%s", url.PathEscape(storageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Refresh CDN storage from remote server.
func (r *ClientV1CdnStorageService) Refresh(ctx context.Context, storageID string, body ClientV1CdnStorageRefreshParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", body.XAuthEmail)))
	}
	if !param.IsOmitted(body.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", body.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if storageID == "" {
		err = errors.New("missing required storage_id parameter")
		return
	}
	path := fmt.Sprintf("client/v1/cdn/storage/%s/refresh", url.PathEscape(storageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type CdnStorageResponse struct {
	ID               string                       `json:"id"`
	BytesTotal       int64                        `json:"bytes_total"`
	CdnUser          CdnUserResponse              `json:"cdn_user"`
	DeletedAt        string                       `json:"deleted_at"`
	FilesCount       int64                        `json:"files_count"`
	LastFileChangeAt string                       `json:"last_file_change_at"`
	LastRefreshAt    string                       `json:"last_refresh_at"`
	Name             string                       `json:"name"`
	NeedsRefresh     bool                         `json:"needs_refresh"`
	Paths            []CdnStorageResponsePath     `json:"paths"`
	PurgeAt          string                       `json:"purge_at"`
	StartRefreshAt   string                       `json:"start_refresh_at"`
	SyncedDc         []CdnStorageResponseSyncedDc `json:"synced_dc"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		BytesTotal       respjson.Field
		CdnUser          respjson.Field
		DeletedAt        respjson.Field
		FilesCount       respjson.Field
		LastFileChangeAt respjson.Field
		LastRefreshAt    respjson.Field
		Name             respjson.Field
		NeedsRefresh     respjson.Field
		Paths            respjson.Field
		PurgeAt          respjson.Field
		StartRefreshAt   respjson.Field
		SyncedDc         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnStorageResponse) RawJSON() string { return r.JSON.raw }
func (r *CdnStorageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdnStorageResponsePath struct {
	ID         int64  `json:"id"`
	DomainName string `json:"domain_name"`
	// computed: subdomain.domain/path_prefix
	FullPath   string `json:"full_path"`
	PathPrefix string `json:"path_prefix"`
	StorageID  string `json:"storage_id"`
	Subdomain  string `json:"subdomain"`
	ZoneID     string `json:"zone_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DomainName  respjson.Field
		FullPath    respjson.Field
		PathPrefix  respjson.Field
		StorageID   respjson.Field
		Subdomain   respjson.Field
		ZoneID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnStorageResponsePath) RawJSON() string { return r.JSON.raw }
func (r *CdnStorageResponsePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdnStorageResponseSyncedDc struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	RefreshedAt string `json:"refreshed_at"`
	Status      bool   `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Name        respjson.Field
		RefreshedAt respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnStorageResponseSyncedDc) RawJSON() string { return r.JSON.raw }
func (r *CdnStorageResponseSyncedDc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CdnUserResponse struct {
	Password  string `json:"password"`
	SftpHost  string `json:"sftp_host"`
	SftpPort  int64  `json:"sftp_port"`
	StorageID string `json:"storage_id"`
	Username  string `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Password    respjson.Field
		SftpHost    respjson.Field
		SftpPort    respjson.Field
		StorageID   respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdnUserResponse) RawJSON() string { return r.JSON.raw }
func (r *CdnUserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1CdnStorageNewParams struct {
	Name       string            `json:"name" api:"required"`
	XAuthEmail string            `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string            `header:"X-Auth-Key" api:"required" json:"-"`
	ZoneID     param.Opt[string] `json:"zone_id,omitzero"`
	paramObj
}

func (r ClientV1CdnStorageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ClientV1CdnStorageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientV1CdnStorageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientV1CdnStorageListParams struct {
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	paramObj
}

type ClientV1CdnStorageDeleteParams struct {
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	paramObj
}

type ClientV1CdnStorageRefreshParams struct {
	XAuthEmail string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string `header:"X-Auth-Key" api:"required" json:"-"`
	paramObj
}
