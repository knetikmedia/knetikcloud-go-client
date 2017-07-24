# UserActivityResults

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyRewards** | [**[]RewardCurrencyResource**](RewardCurrencyResource.md) | Any currency rewarded to this user | [optional] [default to null]
**ItemRewards** | [**[]RewardItemResource**](RewardItemResource.md) | Any items rewarded to this user | [optional] [default to null]
**Rank** | **int64** | The position of the user in the leaderboard. Null means non-compete or disqualification | [optional] [default to null]
**Score** | **int64** | The raw score in this leaderboard. Null means non-compete or disqualification | [optional] [default to null]
**Tags** | **[]string** | Any tags for the metric. Each unique tag will translate into a unique leaderboard. Maximum 5 tags and 50 characters each | [optional] [default to null]
**Ties** | **int32** | The number of users tied at this rank, including this user. 1 means no tie | [optional] [default to null]
**UpdatedDate** | **int64** | The date this score was recorded or updated. Unix timestamp in seconds | [optional] [default to null]
**User** | [**SimpleUserResource**](SimpleUserResource.md) | The player for this entry | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


