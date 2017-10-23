# ObjectResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behaviors** | [**[]Behavior**](Behavior.md) | The behaviors linked to the item, describing various options and interactions. May not be included in item lists | [optional] [default to null]
**Category** | **string** | A category for filtering items | [optional] [default to null]
**CreatedDate** | **int64** | The date the item was created, unix timestamp in seconds | [optional] [default to null]
**Data** | [***interface{}**](interface{}.md) | A map of additional data. The template will define requirements for the object | [optional] [default to null]
**Id** | **int32** | The id of the item | [optional] [default to null]
**LongDescription** | **string** | A long description of the item | [optional] [default to null]
**Name** | **string** | The name of the item | [default to null]
**ShortDescription** | **string** | A short description of the item, max 255 chars | [optional] [default to null]
**Sort** | **int32** | A number to use in sorting items.  Default 500 | [optional] [default to null]
**Tags** | **[]string** | List of tags used for filtering items | [optional] [default to null]
**UniqueKey** | **string** | The unique key for the item | [optional] [default to null]
**UpdatedDate** | **int64** | The date the item was last updated, unix timestamp in seconds | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


