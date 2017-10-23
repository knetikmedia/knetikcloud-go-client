# MobileDeviceResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalProperties** | [**map[string]Property**](Property.md) | A map of additional properties, keyed on the property name.  Must match the names and types defined in the template if one is specified | [optional] [default to null]
**CreatedDate** | **int64** | The date the device log was created | [optional] [default to null]
**Description** | **string** | The description of the device | [optional] [default to null]
**DeviceType** | **string** | The type of device. Use mobile to specifically register mobile devices. This particular type will be used to send and receive notifications | [optional] [default to null]
**Id** | **string** | The unique ID for this device | [optional] [default to null]
**Location** | **string** | The physical location of the device, coordinates or named place (office, living room, etc) | [optional] [default to null]
**MacAddress** | **string** | The MAC (media access control) address of the device | [optional] [default to null]
**Make** | **string** | The make of the device | [optional] [default to null]
**Model** | **string** | The model of the device | [optional] [default to null]
**Name** | **string** | The name of the device | [optional] [default to null]
**Os** | **string** | The OS (operating system) on the device | [optional] [default to null]
**Owner** | [***SimpleUserResource**](SimpleUserResource.md) | The user that owns the device | [optional] [default to null]
**Serial** | **string** | The serial number of the device | [optional] [default to null]
**Tags** | **[]string** | Random tags to facilitate search | [optional] [default to null]
**Template** | **string** | Use to describe and validate custom properties (custom schema). May be null and no validation of additional_properties will be done | [optional] [default to null]
**UpdatedDate** | **int64** | The date the device log was updated | [optional] [default to null]
**Users** | [**[]SimpleUserResource**](SimpleUserResource.md) | The users currently using the device | [optional] [default to null]
**Authorization** | **string** | The authorization code for push notifications provided by the provider platform (APNS, GCM, etc). | [optional] [default to null]
**Imei** | **string** |  | [optional] [default to null]
**NotificationPlatform** | **string** | The platform used for push notifications. Only Apple and Android are supported at the moment. | [optional] [default to null]
**Number** | **string** | The phone number associated with this device if applicable, in international format | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


