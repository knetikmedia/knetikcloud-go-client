# NotificationTypeResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | **int64** | The date the type was created, as a unix timestamp in seconds | [optional] [default to null]
**EmailBodyTemplateExternal** | **bool** | Whether the email body should be resolved. If true, the external email delivery system will be expected to handle the templating (Mandrill/Mailchimp). default: false | [optional] [default to null]
**EmailBodyTemplateId** | **string** | The id of a message template to resolve the email body. If null, no websocket message wil be sent | [optional] [default to null]
**EmailSubjectTemplateId** | **string** | The id of a message template to resolve the email subject. If null, no websocket message wil be sent | [optional] [default to null]
**Id** | **string** | The id of the notification type. Default: random | [optional] [default to null]
**Name** | **string** | The name of the notification type | [default to null]
**SmsTemplateId** | **string** | The id of a message template to resolve the SMS message. If null, no sms message wil be sent | [optional] [default to null]
**UpdatedDate** | **int64** | The date the type was last updated, as a unix timestamp in seconds | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


