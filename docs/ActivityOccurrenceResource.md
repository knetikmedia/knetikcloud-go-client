# ActivityOccurrenceResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | **int64** | The id of the activity | [default to null]
**ChallengeActivityId** | **int64** | The id of the challenge activity (as part of the event, required if eventId set) | [optional] [default to null]
**CreatedDate** | **int64** | The date this occurrence was created, unix timestamp in seconds | [optional] [default to null]
**Entitlement** | [***ActivityEntitlementResource**](ActivityEntitlementResource.md) | The entitlement item required to enter the occurrence. Required if not part of an event. Must come from the set of entitlement items listed in the activity | [optional] [default to null]
**EventId** | **int64** | The id of the event | [optional] [default to null]
**Id** | **int64** | The id of the activity occurrence | [optional] [default to null]
**RewardStatus** | **string** | Indicate if the rewards have been given out already | [optional] [default to null]
**Settings** | [**[]SelectedSettingResource**](SelectedSettingResource.md) | The values selected from the available settings defined for the activity. Ex: difficulty: hard. Can be left out if the activity is played during an event and the settings are already set at the event level. Ex: every monday, difficulty: hard, number of questions: 10, category: sport. Otherwise, the set must exactly match those of the activity. | [optional] [default to null]
**Simulated** | **bool** | Whether this occurrence will be ran as a simulation. Simulations will not be rewarded. Useful for bot play or trials | [optional] [default to null]
**StartDate** | **int64** | The date this occurrence was started, unix timestamp in seconds. null if not yet started | [optional] [default to null]
**Status** | **string** | The current status of the occurrence (default: OPEN) | [optional] [default to null]
**UpdatedDate** | **int64** | The date this occurrence was last updated, unix timestamp in seconds | [optional] [default to null]
**Users** | [**[]ActivityUserResource**](ActivityUserResource.md) | The list of users participating in this occurrence. Can only be set directly with ACTIVITIES_ADMIN permission | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


