# UserItemLogResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The log entry id | [optional] [default to null]
**Info** | **string** | Additional information defined by the type | [optional] [default to null]
**Item** | [**SimpleReferenceResourceint**](SimpleReferenceResource«int».md) | The item interacted with | [optional] [default to null]
**LogDate** | **int64** | The date/time this event occurred in seconds since epoch | [optional] [default to null]
**Type_** | **string** | The type of event | [optional] [default to null]
**User** | [**SimpleUserResource**](SimpleUserResource.md) | The user making the interaction | [optional] [default to null]
**UserInventory** | **int32** | The id of the inventory entry this event is related to, if any | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


