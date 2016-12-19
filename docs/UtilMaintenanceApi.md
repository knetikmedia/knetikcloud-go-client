# \UtilMaintenanceApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsingDELETE1**](UtilMaintenanceApi.md#DeleteUsingDELETE1) | **Delete** /maintenance | Remove maintenance info
[**GetUsingGET1**](UtilMaintenanceApi.md#GetUsingGET1) | **Get** /maintenance | Get current maintenance info
[**PostUsingPOST**](UtilMaintenanceApi.md#PostUsingPOST) | **Post** /maintenance | Set current maintenance info
[**PostUsingPUT**](UtilMaintenanceApi.md#PostUsingPUT) | **Put** /maintenance | Set current maintenance info


# **DeleteUsingDELETE1**
> DeleteUsingDELETE1()

Remove maintenance info


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsingGET1**
> Maintenance GetUsingGET1()

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
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsingPOST**
> PostUsingPOST($maintenance)

Set current maintenance info


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenance** | [**Maintenance**](Maintenance.md)| The Maintenance Object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUsingPUT**
> PostUsingPUT($maintenance)

Set current maintenance info


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenance** | [**Maintenance**](Maintenance.md)| The Maintenance Object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

