# LeaderboardEntryResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rank** | **int64** | The position of the user in the leaderboard. Null means non-compete or disqualification | [optional] [default to null]
**Score** | **int64** | The raw score in this leaderboard. Null means non-compete or disqualification | [optional] [default to null]
**UpdatedDate** | **int64** | The date this score was recorded or updated. Unix timestamp in seconds | [optional] [default to null]
**User** | [**SimpleUserResource**](SimpleUserResource.md) | The player for this entry | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


