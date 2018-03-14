# ServiceDeployedEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | **string** |  | [optional] [default to null]
**Customer** | **string** |  | [optional] [default to null]
**DoNotBroadcast** | **bool** |  | [optional] [default to null]
**Section** | **string** |  | [optional] [default to null]
**Source** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Specifics** | **string** |  | [optional] [default to null]
**Synchronous** | **bool** |  | [optional] [default to null]
**Timestamp** | **int64** |  | [optional] [default to null]
**Type_** | **string** | The type of the event. Used for polymorphic type recognition and thus must match an expected type | [default to null]
**Events** | [**[]BreTriggerResource**](BreTriggerResource.md) | The events fired by this service | [default to null]
**Resources** | [**[]ResourceTypeDescription**](ResourceTypeDescription.md) | The resources managed by this service | [default to null]
**ServiceName** | **string** | The unique name for the service. This serves as the unique identifier. Cannot be changed after creation | [default to null]
**SwaggerUrl** | **string** | The url of the swagger doc | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


