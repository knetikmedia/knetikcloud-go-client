# \BRERuleEngineExpressionsApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBREExpressions**](BRERuleEngineExpressionsApi.md#GetBREExpressions) | **Get** /bre/expressions/lookup | Get a list of &#39;lookup&#39; type expressions


# **GetBREExpressions**
> []LookupTypeResource GetBREExpressions(ctx, ctx, )
Get a list of 'lookup' type expressions

These are expression types that take a second expression as input and produce a value. These can be used in addition to the standard types, like parameter, global and constant (see BRE documentation for details).

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]LookupTypeResource**](LookupTypeResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

