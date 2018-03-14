# \Rule_Engine_GlobalsApi

All URIs are relative to *https://jsapi-integration.us-east-1.elasticbeanstalk.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBREGlobal**](Rule_Engine_GlobalsApi.md#CreateBREGlobal) | **Post** /bre/globals/definitions | Create a global definition
[**DeleteBREGlobal**](Rule_Engine_GlobalsApi.md#DeleteBREGlobal) | **Delete** /bre/globals/definitions/{id} | Delete a global
[**GetBREGlobal**](Rule_Engine_GlobalsApi.md#GetBREGlobal) | **Get** /bre/globals/definitions/{id} | Get a single global definition
[**GetBREGlobals**](Rule_Engine_GlobalsApi.md#GetBREGlobals) | **Get** /bre/globals/definitions | List global definitions
[**UpdateBREGlobal**](Rule_Engine_GlobalsApi.md#UpdateBREGlobal) | **Put** /bre/globals/definitions/{id} | Update a global definition


# **CreateBREGlobal**
> BreGlobalResource CreateBREGlobal(ctx, ctx, optional)
Create a global definition

Once created you can then use in a custom rule. Note that global definitions cannot be modified or deleted if in use. <br><br><b>Permissions Needed:</b> BRE_RULE_ENGINE_GLOBALS_ADMIN

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
 **breGlobalResource** | [**BreGlobalResource**](BreGlobalResource.md)| The BRE global resource object | 

### Return type

[**BreGlobalResource**](BreGlobalResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBREGlobal**
> DeleteBREGlobal(ctx, ctx, id)
Delete a global

May fail if there are existing rules against it. Cannot delete core globals. <br><br><b>Permissions Needed:</b> BRE_RULE_ENGINE_GLOBALS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the global definition | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBREGlobal**
> BreGlobalResource GetBREGlobal(ctx, ctx, id)
Get a single global definition

<b>Permissions Needed:</b> BRE_RULE_ENGINE_GLOBALS_USER

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the global definition | 

### Return type

[**BreGlobalResource**](BreGlobalResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBREGlobals**
> PageResourceBreGlobalResource GetBREGlobals(ctx, ctx, optional)
List global definitions

<b>Permissions Needed:</b> BRE_RULE_ENGINE_GLOBALS_USER

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
 **filterSystem** | **bool**| Filter for globals that are system globals when true, or not when false. Leave off for both mixed | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]

### Return type

[**PageResourceBreGlobalResource**](PageResource«BreGlobalResource».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBREGlobal**
> BreGlobalResource UpdateBREGlobal(ctx, ctx, id, optional)
Update a global definition

May fail if new parameters mismatch requirements of existing rules. Cannot update core globals. <br><br><b>Permissions Needed:</b> BRE_RULE_ENGINE_GLOBALS_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the global definition | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the global definition | 
 **breGlobalResource** | [**BreGlobalResource**](BreGlobalResource.md)| The BRE global resource object | 

### Return type

[**BreGlobalResource**](BreGlobalResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

