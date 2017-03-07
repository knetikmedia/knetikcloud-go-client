# ActivityOccurrenceResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | **int64** | The id of the activity | [default to null]
**ChallengeActivityId** | **int64** | The id of the challenge activity (as part of the event, required if eventId set) | [optional] [default to null]
**CreatedDate** | **int64** | The date this occurrence was created, unix timestamp in seconds | [optional] [default to null]
**Entitlement** | [**ActivityEntitlementResource**](ActivityEntitlementResource.md) | The entitlement item required to enter the occurrence. Required if not part of an event. Must come from the set of entitlement items listed in the activity | [optional] [default to null]
**EventId** | **int64** | The id of the event | [optional] [default to null]
**Id** | **int64** | The id of the activity occurrence | [optional] [default to null]
**RewardStatus** | **string** | Indicate if the rewards have been given out already | [optional] [default to null]
**Settings** | [**[]SelectedSettingResource**](SelectedSettingResource.md) | The list of settings and their options available for this activity. Should be null on create if and only if part of an event | [optional] [default to null]
**Simulated** | **bool** | Whether this occurrence will be played as a simulation. | [optional] [default to null]
**StartDate** | **int64** | The date this occurrence was started, unix timestamp in seconds. null if not yet started | [optional] [default to null]
**Status** | **string** | The current status of the occurrence (default: OPEN) | [optional] [default to null]
**UpdatedDate** | **int64** | The date this occurrence was last updated, unix timestamp in seconds | [optional] [default to null]
**Users** | [**[]ActivityUserResource**](ActivityUserResource.md) | The list of users playing in this occurrence. Can only be set directly with ACTIVITIES_ADMIN permission | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


