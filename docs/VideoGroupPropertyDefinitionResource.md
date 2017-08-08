# VideoGroupPropertyDefinitionResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldList** | [***PropertyFieldListResource**](PropertyFieldListResource.md) | A list of the fields on both the property definition and property of this type | [optional] [default to null]
**Name** | **string** | The name of the property | [default to null]
**Required** | **bool** | Whether the property is required | [default to null]
**Type_** | **string** | The type of the property. Used for polymorphic type recognition and thus must match an expected type with additional properties. | [default to null]
**FileType** | **string** | If provided, a file type that the property must match | [optional] [default to null]
**MaxCount** | **int32** | If provided, the maximum number of files in the group | [optional] [default to null]
**MaxFileSize** | **int64** | If provided, the maximum allowed size per file in bytes | [optional] [default to null]
**MinCount** | **int32** | If provided, the minimum number of files in the group | [optional] [default to null]
**MaxHeight** | **int32** | If provided, the maximum height of each video | [optional] [default to null]
**MaxLength** | **int32** | If provided, the maximum length of each video | [optional] [default to null]
**MaxWidth** | **int32** | If provided, the maximum width of each video | [optional] [default to null]
**MinHeight** | **int32** | If provided, the minimum height of each video | [optional] [default to null]
**MinLength** | **int32** | If provided, the minimum length of each video | [optional] [default to null]
**MinWidth** | **int32** | If provided, the minimum width of each video | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


