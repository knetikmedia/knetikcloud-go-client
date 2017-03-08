# PollResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Whether the poll is active | [default to null]
**AdditionalProperties** | [**map[string]Property**](Property.md) | A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type | [optional] [default to null]
**Answers** | [**[]PollAnswerResource**](PollAnswerResource.md) | The answers to the poll | [default to null]
**Category** | [**NestedCategory**](NestedCategory.md) | The category for the poll | [default to null]
**CreatedDate** | **int64** | The date/time this resource was created in seconds since unix epoch | [optional] [default to null]
**Id** | **string** | The id of the poll | [optional] [default to null]
**Tags** | **[]string** | The tags for the poll | [optional] [default to null]
**Template** | **string** | A poll template this poll is validated against (private). May be null and no validation of additional_properties will be done | [optional] [default to null]
**Text** | **string** | The text of the poll | [default to null]
**Type_** | **string** | The media type of the poll | [default to null]
**UpdatedDate** | **int64** | The date/time this resource was last updated in seconds since unix epoch | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


