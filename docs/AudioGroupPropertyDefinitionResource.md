# AudioGroupPropertyDefinitionResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the property | [optional] [default to null]
**FieldList** | [***PropertyFieldListResource**](PropertyFieldListResource.md) | A list of the fields on both the property definition and property of this type | [optional] [default to null]
**FriendlyName** | **string** | The friendly front-facing name of the property | [optional] [default to null]
**Name** | **string** | The name of the property | [default to null]
**OptionLabelPath** | **string** | The JSON path to the option label | [optional] [default to null]
**OptionValuePath** | **string** | The JSON path to the option value | [optional] [default to null]
**OptionsUrl** | **string** | URL of service containing the property options (assumed JSON array) | [optional] [default to null]
**Required** | **bool** | Whether the property is required | [default to null]
**Type_** | **string** | The type of the property. Used for polymorphic type recognition and thus must match an expected type with additional properties. | [default to null]
**FileType** | **string** | If provided, a file type that the property must match | [optional] [default to null]
**MaxCount** | **int32** | If provided, the maximum number of files in the group | [optional] [default to null]
**MaxFileSize** | **int64** | If provided, the maximum allowed size per file in bytes | [optional] [default to null]
**MinCount** | **int32** | If provided, the minimum number of files in the group | [optional] [default to null]
**MaxLength** | **int32** | If provided, the maximum length of the audio | [optional] [default to null]
**MinLength** | **int32** | If provided, the minimum length of the audio | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


