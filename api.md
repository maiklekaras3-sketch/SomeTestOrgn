# Client

## V1

### Accounts

#### AccountID

##### Firewall

###### AccessRules

Methods:

- <code title="post /client/v1/accounts/:accountID/firewall/access_rules/rules">client.Client.V1.Accounts.AccountID.Firewall.AccessRules.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1AccountAccountIDFirewallAccessRuleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1AccountAccountIDFirewallAccessRuleNewParams">ClientV1AccountAccountIDFirewallAccessRuleNewParams</a>) (\*<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Cdn

#### Storage

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#CdnStorageResponse">CdnStorageResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#CdnUserResponse">CdnUserResponse</a>

Methods:

- <code title="post /client/v1/cdn/storage">client.Client.V1.Cdn.Storage.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1CdnStorageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1CdnStorageNewParams">ClientV1CdnStorageNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#CdnStorageResponse">CdnStorageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /client/v1/cdn/storage">client.Client.V1.Cdn.Storage.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1CdnStorageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1CdnStorageListParams">ClientV1CdnStorageListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#CdnStorageResponse">CdnStorageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /client/v1/cdn/storage/{storage_id}">client.Client.V1.Cdn.Storage.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1CdnStorageService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, storageID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1CdnStorageDeleteParams">ClientV1CdnStorageDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /client/v1/cdn/storage/{storage_id}/refresh">client.Client.V1.Cdn.Storage.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1CdnStorageService.Refresh">Refresh</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, storageID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1CdnStorageRefreshParams">ClientV1CdnStorageRefreshParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

##### User

Methods:

- <code title="get /client/v1/cdn/storage/{storage_id}/user">client.Client.V1.Cdn.Storage.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1CdnStorageUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, storageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1CdnStorageUserGetParams">ClientV1CdnStorageUserGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#CdnUserResponse">CdnUserResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### User

Methods:

- <code title="post /client/v1/cdn/user">client.Client.V1.Cdn.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1CdnUserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1CdnUserNewParams">ClientV1CdnUserNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#CdnUserResponse">CdnUserResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /client/v1/cdn/user/{username}/revoke">client.Client.V1.Cdn.User.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1CdnUserService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1CdnUserRevokeParams">ClientV1CdnUserRevokeParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#CdnUserResponse">CdnUserResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Zones

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ResponseInfo">ResponseInfo</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ResponseZone">ResponseZone</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ResultInfo">ResultInfo</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ZoneAPIResponse">ZoneAPIResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneListResponse">ClientV1ZoneListResponse</a>

Methods:

- <code title="post /client/v1/zones">client.Client.V1.Zones.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneNewParams">ClientV1ZoneNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ZoneAPIResponse">ZoneAPIResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /client/v1/zones">client.Client.V1.Zones.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneListParams">ClientV1ZoneListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneListResponse">ClientV1ZoneListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### ZoneID

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDDeleteResponse">ClientV1ZoneZoneIDDeleteResponse</a>

Methods:

- <code title="get /client/v1/zones/:zoneId">client.Client.V1.Zones.ZoneID.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDGetParams">ClientV1ZoneZoneIDGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ZoneAPIResponse">ZoneAPIResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /client/v1/zones/:zoneId">client.Client.V1.Zones.ZoneID.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDDeleteParams">ClientV1ZoneZoneIDDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDDeleteResponse">ClientV1ZoneZoneIDDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### AccessRules

Methods:

- <code title="post /client/v1/zones/:zoneId/access_rules">client.Client.V1.Zones.ZoneID.AccessRules.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDAccessRuleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDAccessRuleNewParams">ClientV1ZoneZoneIDAccessRuleNewParams</a>) (\*<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### DNSRecords

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ZoneDNSRecord">ZoneDNSRecord</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIddnsRecordListResponse">ClientV1ZoneZoneIddnsRecordListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIddnsRecordBatchResponse">ClientV1ZoneZoneIddnsRecordBatchResponse</a>

Methods:

- <code title="get /client/v1/zones/:zoneId/dns_records">client.Client.V1.Zones.ZoneID.DNSRecords.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDDNSRecordService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDDNSRecordListParams">ClientV1ZoneZoneIDDNSRecordListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIddnsRecordListResponse">ClientV1ZoneZoneIddnsRecordListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /client/v1/zones/:zoneId/dns_records/batch">client.Client.V1.Zones.ZoneID.DNSRecords.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDDNSRecordService.Batch">Batch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDDNSRecordBatchParams">ClientV1ZoneZoneIDDNSRecordBatchParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIddnsRecordBatchResponse">ClientV1ZoneZoneIddnsRecordBatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Settings

###### SettingID

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ZoneSettingAPIResponse">ZoneSettingAPIResponse</a>

Methods:

- <code title="get /client/v1/zones/:zoneId/settings/:settingId">client.Client.V1.Zones.ZoneID.Settings.SettingID.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDSettingSettingIDService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDSettingSettingIDGetParams">ClientV1ZoneZoneIDSettingSettingIDGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ZoneSettingAPIResponse">ZoneSettingAPIResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /client/v1/zones/:zoneId/settings/:settingId">client.Client.V1.Zones.ZoneID.Settings.SettingID.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDSettingSettingIDService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDSettingSettingIDUpdateParams">ClientV1ZoneZoneIDSettingSettingIDUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ZoneSettingAPIResponse">ZoneSettingAPIResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

##### Subscription

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#BaseAPIResponse">BaseAPIResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDSubscriptionGetResponse">ClientV1ZoneZoneIDSubscriptionGetResponse</a>

Methods:

- <code title="post /client/v1/zones/:zoneId/subscription">client.Client.V1.Zones.ZoneID.Subscription.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDSubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDSubscriptionNewParams">ClientV1ZoneZoneIDSubscriptionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#BaseAPIResponse">BaseAPIResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /client/v1/zones/:zoneId/subscription">client.Client.V1.Zones.ZoneID.Subscription.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDSubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDSubscriptionGetParams">ClientV1ZoneZoneIDSubscriptionGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDSubscriptionGetResponse">ClientV1ZoneZoneIDSubscriptionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /client/v1/zones/:zoneId/subscription">client.Client.V1.Zones.ZoneID.Subscription.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDSubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#ClientV1ZoneZoneIDSubscriptionUpdateParams">ClientV1ZoneZoneIDSubscriptionUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go">stnlsstest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/stnlss_test-go#BaseAPIResponse">BaseAPIResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
