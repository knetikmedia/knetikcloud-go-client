# TemplateEmailResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | Address to attribute the outgoing message to. Optional if the config email.out_address is set. | [optional] [default to null]
**Recipients** | **[]int32** | A list of user ids to send the message to. | [default to null]
**TemplateKey** | **string** | The key for the template | [default to null]
**TemplateVars** | [**[]KeyValuePairstringstring**](KeyValuePair«string,string».md) | A list of variables to fill in the template | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


