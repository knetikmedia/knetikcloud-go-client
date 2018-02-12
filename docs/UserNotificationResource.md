# UserNotificationResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [***interface{}**](interface{}.md) | The data to send and fill templates | [optional] [default to null]
**NotificationId** | **string** | The id of the notification | [default to null]
**NotificationTypeId** | **string** | The id of the notification type | [default to null]
**Recipient** | **string** | The id of the recipient, dependent on the recipient_type. The user&#39;s id or the topic&#39;s id | [default to null]
**RecipientType** | **string** | The type of recipient for the notification. Either a user, or all users in a topic | [default to null]
**RetrieveDate** | **int64** | The date this notification was first retrieved | [optional] [default to null]
**SendDate** | **int64** | The date this notification was sent | [optional] [default to null]
**Status** | **string** | The user&#39;s status for this notification | [optional] [default to null]
**UserId** | **int32** | The id of the user | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


