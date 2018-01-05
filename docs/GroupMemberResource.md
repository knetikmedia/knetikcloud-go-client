# GroupMemberResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalProperties** | [**map[string]Property**](Property.md) | A map of additional properties, keyed on the property name (private). Must match the names and types defined in the template for this type, or be an extra not from the template | [optional] [default to null]
**Group** | [***SimpleGroupResource**](SimpleGroupResource.md) | The group. Id is the unique name | [optional] [default to null]
**Implicit** | **bool** | Whether this membership is explicit (the user was added directly to the group) or implicit (the user was added only to one or more child groups) | [optional] [default to null]
**MembershipId** | **int64** | The id of the membership entry | [optional] [default to null]
**Order** | **string** | The position of the member in the group if applicable. Read notes for details | [optional] [default to null]
**Status** | **string** | The member&#39;s access level. Default: member | [optional] [default to null]
**Template** | **string** | A template this member additional properties are validated against (private). May be null and no validation of properties will be done | [optional] [default to null]
**User** | [***SimpleUserResource**](SimpleUserResource.md) | The user | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


