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
)

// ClientV1AccountAccountIDFirewallAccessRuleService contains methods and other
// services that help with interacting with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1AccountAccountIDFirewallAccessRuleService] method instead.
type ClientV1AccountAccountIDFirewallAccessRuleService struct {
	Options []option.RequestOption
}

// NewClientV1AccountAccountIDFirewallAccessRuleService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewClientV1AccountAccountIDFirewallAccessRuleService(opts ...option.RequestOption) (r ClientV1AccountAccountIDFirewallAccessRuleService) {
	r = ClientV1AccountAccountIDFirewallAccessRuleService{}
	r.Options = opts
	return
}

// Creates a new IP Access rule for an account or zone. The rule will apply to all
// zones in the account or zone.
func (r *ClientV1AccountAccountIDFirewallAccessRuleService) New(ctx context.Context, params ClientV1AccountAccountIDFirewallAccessRuleNewParams, opts ...option.RequestOption) (res *string, err error) {
	if !param.IsOmitted(params.XAuthEmail) {
		opts = append(opts, option.WithHeader("X-Auth-Email", fmt.Sprintf("%v", params.XAuthEmail)))
	}
	if !param.IsOmitted(params.XAuthKey) {
		opts = append(opts, option.WithHeader("X-Auth-Key", fmt.Sprintf("%v", params.XAuthKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "client/v1/accounts/:accountID/firewall/access_rules/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type ClientV1AccountAccountIDFirewallAccessRuleNewParams struct {
	Configuration ClientV1AccountAccountIDFirewallAccessRuleNewParamsConfiguration `json:"configuration,omitzero" api:"required"`
	// Action mode.
	//
	// Any of "challenge", "block", "whitelist".
	Mode       ClientV1AccountAccountIDFirewallAccessRuleNewParamsMode `json:"mode,omitzero" api:"required"`
	XAuthEmail string                                                  `header:"X-Auth-Email" api:"required" json:"-"`
	XAuthKey   string                                                  `header:"X-Auth-Key" api:"required" json:"-"`
	// Notes for access rule.
	Notes param.Opt[string] `json:"notes,omitzero"`
	paramObj
}

func (r ClientV1AccountAccountIDFirewallAccessRuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ClientV1AccountAccountIDFirewallAccessRuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientV1AccountAccountIDFirewallAccessRuleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Target, Value are required.
type ClientV1AccountAccountIDFirewallAccessRuleNewParamsConfiguration struct {
	// Configuration target.
	//
	// Any of "ip", "country".
	Target string `json:"target,omitzero" api:"required"`
	// The value of type.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ClientV1AccountAccountIDFirewallAccessRuleNewParamsConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow ClientV1AccountAccountIDFirewallAccessRuleNewParamsConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientV1AccountAccountIDFirewallAccessRuleNewParamsConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ClientV1AccountAccountIDFirewallAccessRuleNewParamsConfiguration](
		"target", "ip", "country",
	)
}

// Action mode.
type ClientV1AccountAccountIDFirewallAccessRuleNewParamsMode string

const (
	ClientV1AccountAccountIDFirewallAccessRuleNewParamsModeChallenge ClientV1AccountAccountIDFirewallAccessRuleNewParamsMode = "challenge"
	ClientV1AccountAccountIDFirewallAccessRuleNewParamsModeBlock     ClientV1AccountAccountIDFirewallAccessRuleNewParamsMode = "block"
	ClientV1AccountAccountIDFirewallAccessRuleNewParamsModeWhitelist ClientV1AccountAccountIDFirewallAccessRuleNewParamsMode = "whitelist"
)
