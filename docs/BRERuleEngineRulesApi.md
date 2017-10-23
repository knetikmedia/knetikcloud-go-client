# \BRERuleEngineRulesApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBRERule**](BRERuleEngineRulesApi.md#CreateBRERule) | **Post** /bre/rules | Create a rule
[**DeleteBRERule**](BRERuleEngineRulesApi.md#DeleteBRERule) | **Delete** /bre/rules/{id} | Delete a rule
[**GetBREExpressionAsString**](BRERuleEngineRulesApi.md#GetBREExpressionAsString) | **Post** /bre/rules/expression-as-string | Returns a string representation of the provided expression
[**GetBRERule**](BRERuleEngineRulesApi.md#GetBRERule) | **Get** /bre/rules/{id} | Get a single rule
[**GetBRERules**](BRERuleEngineRulesApi.md#GetBRERules) | **Get** /bre/rules | List rules
[**SetBRERule**](BRERuleEngineRulesApi.md#SetBRERule) | **Put** /bre/rules/{id}/enabled | Enable or disable a rule
[**UpdateBRERule**](BRERuleEngineRulesApi.md#UpdateBRERule) | **Put** /bre/rules/{id} | Update a rule


# **CreateBRERule**
> BreRule CreateBRERule(ctx, ctx, optional)
Create a rule

Rules define which actions to run when a given event verifies the specified condition. Full list of predicates and other type of expressions can be found at GET /bre/expressions/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **breRule** | [**BreRule**](BreRule.md)| The BRE rule object | 

### Return type

[**BreRule**](BreRule.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBRERule**
> DeleteBRERule(ctx, ctx, id)
Delete a rule

May fail if there are existing rules against it. Cannot delete core rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the rule | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBREExpressionAsString**
> string GetBREExpressionAsString(ctx, ctx, optional)
Returns a string representation of the provided expression

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expression** | [**Expressionobject**](Expressionobject.md)| The expression | 

### Return type

**string**

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBRERule**
> BreRule GetBRERule(ctx, ctx, id)
Get a single rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the rule | 

### Return type

[**BreRule**](BreRule.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBRERules**
> PageResourceBreRule GetBRERules(ctx, ctx, optional)
List rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterName** | **string**| Filter for rules containing the given name | 
 **filterEnabled** | **bool**| Filter for rules by active status, null for both | [default to null]
 **filterSystem** | **bool**| Filter for rules that are system rules when true, or not when false. Leave off for both mixed | 
 **filterTrigger** | **string**| Filter for rules that are for the trigger with the given name | 
 **filterAction** | **string**| Filter for rules that use the action with the given name | 
 **filterCondition** | **string**| Filter for rules that have a condition containing the given string | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceBreRule**](PageResource«BreRule».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetBRERule**
> SetBRERule(ctx, ctx, id, optional)
Enable or disable a rule

This is helpful for turning off systems rules which cannot be deleted or modified otherwise

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the rule | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the rule | 
 **enabled** | [**BooleanResource**](BooleanResource.md)| The boolean value | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBRERule**
> BreRule UpdateBRERule(ctx, ctx, id, optional)
Update a rule

Cannot update system rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the rule | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the rule | 
 **breRule** | [**BreRule**](BreRule.md)| The BRE rule object | 

### Return type

[**BreRule**](BreRule.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

