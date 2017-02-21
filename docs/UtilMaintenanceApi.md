# \UtilMaintenanceApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMaintenance**](UtilMaintenanceApi.md#DeleteMaintenance) | **Delete** /maintenance | Delete maintenance info
[**GetMaintenance**](UtilMaintenanceApi.md#GetMaintenance) | **Get** /maintenance | Get current maintenance info
[**SetMaintenance**](UtilMaintenanceApi.md#SetMaintenance) | **Post** /maintenance | Set current maintenance info
[**UpdateMaintenance**](UtilMaintenanceApi.md#UpdateMaintenance) | **Put** /maintenance | Update current maintenance info


# **DeleteMaintenance**
> DeleteMaintenance()

Delete maintenance info


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMaintenance**
> Maintenance GetMaintenance()

Get current maintenance info

Get current maintenance info. 404 if no maintenance.


### Parameters
This endpoint does not need any parameter.

### Return type

[**Maintenance**](Maintenance.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetMaintenance**
> SetMaintenance($maintenance)

Set current maintenance info


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenance** | [**Maintenance**](Maintenance.md)| The maintenance object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMaintenance**
> UpdateMaintenance($maintenance)

Update current maintenance info


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenance** | [**Maintenance**](Maintenance.md)| The maintenance object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

