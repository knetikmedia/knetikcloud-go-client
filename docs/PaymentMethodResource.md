# PaymentMethodResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | **int64** | The date/time this resource was created in seconds since unix epoch | [optional] [default to null]
**Default_** | **bool** |  | [optional] [default to null]
**Disabled** | **bool** | Whether this payment method is disabled or not | [optional] [default to null]
**ExpirationDate** | **int64** | The expiration date for the payment method, expressed as seconds since epoch. Typically used for credit card payment methods | [optional] [default to null]
**ExpirationMonth** | **int32** | The expiration month (1 - 12) for the payment method. Typically used for credit card payment methods | [optional] [default to null]
**ExpirationYear** | **int32** | The expiration year for the payment method. Typically used for credit card payment methods | [optional] [default to null]
**Id** | **int64** | The unique ID of the resource | [optional] [default to null]
**Last4** | **string** | The last 4 digits of the account number for the payment method. Typically used for credit card payment methods | [optional] [default to null]
**Name** | **string** | The user friendly name of the resource | [default to null]
**PaymentMethodType** | [***PaymentMethodTypeResource**](PaymentMethodTypeResource.md) | The type of payment method. Must be a pre-existing value | [default to null]
**PaymentType** | **string** | The generic payment type. Default is card | [optional] [default to null]
**Sort** | **int32** | The sort value for the payment method | [optional] [default to null]
**Token** | **string** | The unique token for the payment method | [optional] [default to null]
**UniqueKey** | **string** | An optional unique identifier | [optional] [default to null]
**UpdatedDate** | **int64** | The date/time this resource was last updated in seconds since unix epoch | [optional] [default to null]
**UserId** | **int32** | The user&#39;s id. If null, indicates a shared payment method that any user can use (i.e., &#39;wallet&#39;) | [optional] [default to null]
**Verified** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


