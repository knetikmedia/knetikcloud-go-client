# \DevicesApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDeviceUsers**](DevicesApi.md#AddDeviceUsers) | **Post** /devices/{id}/users | Add device users
[**CreateDevice**](DevicesApi.md#CreateDevice) | **Post** /devices | Create a device
[**DeleteDevice**](DevicesApi.md#DeleteDevice) | **Delete** /devices/{id} | Delete a device
[**DeleteDeviceUser**](DevicesApi.md#DeleteDeviceUser) | **Delete** /devices/{id}/users/{user_id} | Delete a device user
[**DeleteDeviceUsers**](DevicesApi.md#DeleteDeviceUsers) | **Delete** /devices/{id}/users | Delete all device users
[**GetDevice**](DevicesApi.md#GetDevice) | **Get** /devices/{id} | Get a single device
[**GetDevices**](DevicesApi.md#GetDevices) | **Get** /devices | List and search devices
[**UpdateDevice**](DevicesApi.md#UpdateDevice) | **Put** /devices/{id} | Update a device


# **AddDeviceUsers**
> DeviceResource AddDeviceUsers(userResources, id)
Add device users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **userResources** | [**[]SimpleUserResource**](SimpleUserResource.md)| userResources | 
  **id** | **int32**| id | 

### Return type

[**DeviceResource**](DeviceResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDevice**
> DeviceResource CreateDevice(device)
Create a device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **device** | [**DeviceResource**](DeviceResource.md)| device | 

### Return type

[**DeviceResource**](DeviceResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDevice**
> DeleteDevice(id)
Delete a device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **id** | **int32**| id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeviceUser**
> DeleteDeviceUser(id, userId)
Delete a device user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **id** | **int32**| The id of the device | 
  **userId** | **int32**| The user id of the device user | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeviceUsers**
> DeleteDeviceUsers(id, optional)
Delete all device users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **id** | **int32**| The id of the device | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the device | 
 **filterId** | **string**| Filter for device users to delete with a user id in a given comma separated list of ids | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevice**
> DeviceResource GetDevice(id)
Get a single device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **id** | **int32**| id | 

### Return type

[**DeviceResource**](DeviceResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDevices**
> PageResourceDeviceResource GetDevices(optional)
List and search devices

Get a list of devices with optional filtering

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterMake** | **string**| Filter for devices with specified make | 
 **filterModel** | **string**| Filter for devices with specified model | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to id:ASC]

### Return type

[**PageResourceDeviceResource**](PageResource«DeviceResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDevice**
> DeviceResource UpdateDevice(device, id)
Update a device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **device** | [**DeviceResource**](DeviceResource.md)| device | 
  **id** | **int32**| id | 

### Return type

[**DeviceResource**](DeviceResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

