# MetricResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityOccurenceId** | **int64** | The id of the activity occurence where this score/metric occurred | [default to null]
**Tags** | **[]string** | Any tags for the metric. Each unique tag will translate into a unique leaderboard. Maximum 10 tags and 50 characters each | [optional] [default to null]
**UserId** | **int32** | The id of the user this metric is for. Default to caller and requires METRICS_ADMIN permission to specify another | [optional] [default to null]
**Value** | **int64** | The value/score of the metric | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


