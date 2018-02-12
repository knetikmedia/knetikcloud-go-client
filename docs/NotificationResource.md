# NotificationResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [***interface{}**](interface{}.md) | The data to send to websockets. Also used to fill in the templates | [optional] [default to null]
**NotificationId** | **string** | The id of this individual notification. Default: random | [optional] [default to null]
**NotificationTypeId** | **string** | The id of the notification type which will define message templates | [default to null]
**Recipient** | **string** | The id of the recipient, dependent on the recipient_type. The user&#39;s id or the topic&#39;s id | [default to null]
**RecipientType** | **string** | The type of recipient for the notification. Either a user, or all users in a topic | [default to null]
**SendDate** | **int64** | The date this notification was sent | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


