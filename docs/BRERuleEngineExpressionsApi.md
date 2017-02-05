# \BRERuleEngineExpressionsApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLookupTypesUsingGET**](BRERuleEngineExpressionsApi.md#GetLookupTypesUsingGET) | **Get** /bre/expressions/lookup | Get a list of &#39;lookup&#39; type expressions


# **GetLookupTypesUsingGET**
> []LookupTypeResource GetLookupTypesUsingGET()

Get a list of 'lookup' type expressions

These are expression types that take a second expression as input and produce a value. These can be used in addition to the standard types, like parameter, global and constant (see BRE documentation for details).


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]LookupTypeResource**](LookupTypeResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

