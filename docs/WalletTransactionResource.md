# WalletTransactionResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **float32** | The new balance of the wallet after the transaction | [optional] [default to null]
**CreateDate** | **int64** | The unix timestamp in seconds of the transaction | [optional] [default to null]
**CurrencyCode** | **string** | The code of the currency for the transaction | [optional] [default to null]
**Details** | **string** | The specific details of the transaction, such as a message from the admin that created it | [optional] [default to null]
**Id** | **int32** | The id of the transaction | [optional] [default to null]
**InvoiceId** | **int32** | The id of the invoice that spawned the transaction, if any | [optional] [default to null]
**IsRefunded** | **bool** | Whether the transaction has been refunded | [optional] [default to null]
**Response** | **string** | The response | [optional] [default to null]
**Source** | **string** | The root source of the transaction | [optional] [default to null]
**Successful** | **bool** | If the transaction was successful | [optional] [default to null]
**TransactionId** | **string** | The payment gateway (external) transaction ID | [optional] [default to null]
**Type_** | **string** | The general type of the transaction | [optional] [default to null]
**TypeHint** | **string** | The table name of the subclass | [optional] [default to null]
**User** | [***SimpleUserResource**](SimpleUserResource.md) | The owner of the wallet | [optional] [default to null]
**Value** | **float32** | The amount of the transaction, positive if a gain, negative if an expenditure | [optional] [default to null]
**WalletId** | **int32** | The id of the wallet this transaction affected | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


