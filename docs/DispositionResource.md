# DispositionResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | **string** | The context of that resource. Required when passed to /dispositions rather than context specific endpoint | [optional] [default to null]
**ContextId** | **string** | The context_id of that resource. Required when passed to /dispositions rather than context specific endpoint | [optional] [default to null]
**CreatedDate** | **int64** | The date/time this resource was created in seconds since unix epoch | [optional] [default to null]
**Id** | **int64** | The unique ID for that resource | [optional] [default to null]
**Name** | **string** | The name of the disposition, 1-20 characters. (ex: like/dislike/favorite, etc) | [default to null]
**User** | [***SimpleUserResource**](SimpleUserResource.md) | The user | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


