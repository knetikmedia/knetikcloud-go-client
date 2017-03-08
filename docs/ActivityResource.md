# ActivityResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | **int64** | The date/time this resource was created in seconds since unix epoch | [optional] [default to null]
**Entitlements** | [**[]ActivityEntitlementResource**](ActivityEntitlementResource.md) | The list of items that can be used for entitlement (wager amounts/etc) | [optional] [default to null]
**Id** | **int64** | The unique ID for that resource | [optional] [default to null]
**Launch** | **string** | Details about how to launch the activity | [optional] [default to null]
**LongDescription** | **string** | The user friendly name of that resource. Defaults to blank string | [optional] [default to null]
**Name** | **string** | The user friendly name of that resource | [default to null]
**RewardSet** | [**RewardSetResource**](RewardSetResource.md) | The rewards to give at the end of each occurence of the activity. When creating/updating only id is used. Reward set must be pre-existing | [optional] [default to null]
**Settings** | [**[]AvailableSettingResource**](AvailableSettingResource.md) | The list of settings and their options available for this activity. Not populated when getting listing | [optional] [default to null]
**ShortDescription** | **string** | The user friendly name of that resource. Defaults to blank string | [optional] [default to null]
**Template** | **bool** | Whether this activity is a template for other activities. Default: false | [optional] [default to null]
**Type_** | **string** | The type of the activity | [default to null]
**UniqueKey** | **string** | The unique key (for static reference in code) of the activity | [optional] [default to null]
**UpdatedDate** | **int64** | The date/time this resource was last updated in seconds since unix epoch | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


