# \BRERuleEngineVariablesApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBREVariableTypes**](BRERuleEngineVariablesApi.md#GetBREVariableTypes) | **Get** /bre/variable-types | Get a list of variable types available
[**GetBREVariableValues**](BRERuleEngineVariablesApi.md#GetBREVariableValues) | **Get** /bre/variable-types/{name}/values | List valid values for a type


# **GetBREVariableTypes**
> []VariableTypeResource GetBREVariableTypes()

Get a list of variable types available

Types include integer, string, user and invoice. These are used to qualify trigger parameters and action variables with strong typing.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]VariableTypeResource**](VariableTypeResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBREVariableValues**
> PageResourceSimpleReferenceResourceobject GetBREVariableValues($name, $filterName, $size, $page)

List valid values for a type

Used to lookup users to fill in a user constant for example. Only types marked as enumerable are suppoorted here.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the type | 
 **filterName** | **string**| Filter results by those with names starting with this string | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceSimpleReferenceResourceobject**](PageResource«SimpleReferenceResource«object»».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

