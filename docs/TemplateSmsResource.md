# TemplateSmsResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | The phone number to attribute the outgoing message to. Optional if the config text.out_number is set. | [optional] [default to null]
**Recipients** | **[]int32** | A list of user ids to send the message to. | [default to null]
**Template** | **string** | A mustache template | [default to null]
**TemplateVars** | [***interface{}**](interface{}.md) | A map of values to fill in the template | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


