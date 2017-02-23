# \BRERuleEngineExpressionsApi

All URIs are relative to *https://sandbox.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBREExpressions**](BRERuleEngineExpressionsApi.md#GetBREExpressions) | **Get** /bre/expressions/lookup | Get a list of &#39;lookup&#39; type expressions


# **GetBREExpressions**
> []LookupTypeResource GetBREExpressions()

Get a list of 'lookup' type expressions

These are expression types that take a second expression as input and produce a value. These can be used in addition to the standard types, like parameter, global and constant (see BRE documentation for details).


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]LookupTypeResource**](LookupTypeResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

