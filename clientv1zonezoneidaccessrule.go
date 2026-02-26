// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/stnlss_test-go/internal/apijson"
	"github.com/stainless-sdks/stnlss_test-go/internal/requestconfig"
	"github.com/stainless-sdks/stnlss_test-go/option"
	"github.com/stainless-sdks/stnlss_test-go/packages/param"
)

// ClientV1ZoneZoneIDAccessRuleService contains methods and other services that
// help with interacting with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1ZoneZoneIDAccessRuleService] method instead.
type ClientV1ZoneZoneIDAccessRuleService struct {
	Options []option.RequestOption
}

// NewClientV1ZoneZoneIDAccessRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewClientV1ZoneZoneIDAccessRuleService(opts ...option.RequestOption) (r ClientV1ZoneZoneIDAccessRuleService) {
	r = ClientV1ZoneZoneIDAccessRuleService{}
	r.Options = opts
	return
}

// Creates a new IP Access rule for an account or zone. The rule will apply to all
// zones in the account or zone.
func (r *ClientV1ZoneZoneIDAccessRuleService) New(ctx context.Context, params ClientV1ZoneZoneIDAccessRuleNewParams, opts ...option.RequestOption) (res *string, err error) {
	if !param.IsOmitted(params.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", params.XAuthEmail)))
	}
	if !param.IsOmitted(params.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", params.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/zones/:zoneId/access_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ClientV1ZoneZoneIDAccessRuleNewParams struct {
	// Action taken.
	//
	// Any of "nginx_ban", "allow", "ip_tables_ban", "challenge".
	Action ClientV1ZoneZoneIDAccessRuleNewParamsAction `json:"action,omitzero" api:"required"`
	// Configuration Type of checking.
	//
	// Any of "ips", "countries".
	ConfigurationType ClientV1ZoneZoneIDAccessRuleNewParamsConfigurationType `json:"configuration_type,omitzero" api:"required"`
	// The value of type.
	ConfigurationValue string `json:"configuration_value" api:"required"`
	XAuthEmail         string `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey           string `header:"X-Auth-Key" api:"required" json:"-"`
	// Notes for access rule.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Source.
	Source param.Opt[string] `json:"source,omitzero"`
	// Scope of application.
	//
	// Any of "zone", "account".
	Scope ClientV1ZoneZoneIDAccessRuleNewParamsScope `json:"scope,omitzero"`
	paramObj
}

func (r ClientV1ZoneZoneIDAccessRuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ClientV1ZoneZoneIDAccessRuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientV1ZoneZoneIDAccessRuleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action taken.
type ClientV1ZoneZoneIDAccessRuleNewParamsAction string

const (
	ClientV1ZoneZoneIDAccessRuleNewParamsActionNginxBan    ClientV1ZoneZoneIDAccessRuleNewParamsAction = "nginx_ban"
	ClientV1ZoneZoneIDAccessRuleNewParamsActionAllow       ClientV1ZoneZoneIDAccessRuleNewParamsAction = "allow"
	ClientV1ZoneZoneIDAccessRuleNewParamsActionIPTablesBan ClientV1ZoneZoneIDAccessRuleNewParamsAction = "ip_tables_ban"
	ClientV1ZoneZoneIDAccessRuleNewParamsActionChallenge   ClientV1ZoneZoneIDAccessRuleNewParamsAction = "challenge"
)

// Configuration Type of checking.
type ClientV1ZoneZoneIDAccessRuleNewParamsConfigurationType string

const (
	ClientV1ZoneZoneIDAccessRuleNewParamsConfigurationTypeIPs       ClientV1ZoneZoneIDAccessRuleNewParamsConfigurationType = "ips"
	ClientV1ZoneZoneIDAccessRuleNewParamsConfigurationTypeCountries ClientV1ZoneZoneIDAccessRuleNewParamsConfigurationType = "countries"
)

// Scope of application.
type ClientV1ZoneZoneIDAccessRuleNewParamsScope string

const (
	ClientV1ZoneZoneIDAccessRuleNewParamsScopeZone    ClientV1ZoneZoneIDAccessRuleNewParamsScope = "zone"
	ClientV1ZoneZoneIDAccessRuleNewParamsScopeAccount ClientV1ZoneZoneIDAccessRuleNewParamsScope = "account"
)
