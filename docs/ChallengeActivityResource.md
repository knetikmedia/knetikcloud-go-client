# ChallengeActivityResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | **int64** | The id of the activity | [default to null]
**AdditionalProperties** | [**map[string]Property**](Property.md) | A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type | [optional] [default to null]
**ChallengeId** | **int64** | The id of the challenge | [optional] [default to null]
**CoreSettings** | [***CoreChallengeActivitySettings**](CoreChallengeActivitySettings.md) | Defines core settings about the activity that affect how it can be created/played by users. Values may be left null to inherit from parent activity. | [optional] [default to null]
**Entitlement** | [***ActivityEntitlementResource**](ActivityEntitlementResource.md) | The entitlement item needed to participate in the activity as part of this event. Null indicates free entry. When creating/updating only id is used. Item must be pre-existing | [optional] [default to null]
**Id** | **int64** | The unique ID for this resource | [optional] [default to null]
**RewardSet** | [***RewardSetResource**](RewardSetResource.md) | The rewards to give at the end of each occurence of the activity. When creating/updating only id is used. Reward set must be pre-existing | [optional] [default to null]
**Settings** | [**[]SelectedSettingResource**](SelectedSettingResource.md) | The list of settings and the select options | [optional] [default to null]
**Template** | **string** | A challenge activity template this challenge activity is validated against (private). May be null and no validation of additional_properties will be done | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


