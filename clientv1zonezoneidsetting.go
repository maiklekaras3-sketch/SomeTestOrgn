// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stnlsstest

import (
	"github.com/stainless-sdks/stnlss_test-go/option"
)

// ClientV1ZoneZoneIDSettingService contains methods and other services that help
// with interacting with the stnlss_test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientV1ZoneZoneIDSettingService] method instead.
type ClientV1ZoneZoneIDSettingService struct {
	Options   []option.RequestOption
	SettingID ClientV1ZoneZoneIDSettingSettingIDService
}

// NewClientV1ZoneZoneIDSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewClientV1ZoneZoneIDSettingService(opts ...option.RequestOption) (r ClientV1ZoneZoneIDSettingService) {
	r = ClientV1ZoneZoneIDSettingService{}
	r.Options = opts
	r.SettingID = NewClientV1ZoneZoneIDSettingSettingIDService(opts...)
	return
}
