# GroupMemberResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalProperties** | [**map[string]Property**](Property.md) | A map of additional properties, keyed on the property name (private). Must match the names and types defined in the template for this type, or be an extra not from the template | [optional] [default to null]
**AvatarUrl** | **string** | The url of the user&#39;s avatar image | [optional] [default to null]
**DisplayName** | **string** | The public username of the user | [optional] [default to null]
**Id** | **int32** | The id of the user | [default to null]
**Order** | **string** | The position of the member in the group if applicable. Read notes for details | [optional] [default to null]
**Status** | **string** | The member&#39;s access level. Default: member | [optional] [default to null]
**Template** | **string** | A template this member additional properties are validated against (private). May be null and no validation of properties will be done | [optional] [default to null]
**Username** | **string** | The username of the user | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


