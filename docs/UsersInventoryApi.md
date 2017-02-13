# \UsersInventoryApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddItemToUserInventory**](UsersInventoryApi.md#AddItemToUserInventory) | **Post** /users/{id}/inventory | Adds an item to the user inventory
[**CheckUserEntitlementItem**](UsersInventoryApi.md#CheckUserEntitlementItem) | **Get** /users/{user_id}/entitlements/{item_id}/check | Check for access to an item without consuming
[**CreateEntitlementItem**](UsersInventoryApi.md#CreateEntitlementItem) | **Post** /entitlements | Create an entitlement item
[**CreateEntitlementTemplate**](UsersInventoryApi.md#CreateEntitlementTemplate) | **Post** /entitlements/templates | Create an entitlement template
[**DeleteEntitlementItem**](UsersInventoryApi.md#DeleteEntitlementItem) | **Delete** /entitlements/{entitlement_id} | Delete an entitlement item
[**DeleteEntitlementTemplate**](UsersInventoryApi.md#DeleteEntitlementTemplate) | **Delete** /entitlements/templates/{id} | Delete an entitlement template
[**GetEntitlementItem**](UsersInventoryApi.md#GetEntitlementItem) | **Get** /entitlements/{entitlement_id} | Get a single entitlement item
[**GetEntitlementItems**](UsersInventoryApi.md#GetEntitlementItems) | **Get** /entitlements | List and search entitlement items
[**GetEntitlementTemplate**](UsersInventoryApi.md#GetEntitlementTemplate) | **Get** /entitlements/templates/{id} | Get a single entitlement template
[**GetEntitlementTemplates**](UsersInventoryApi.md#GetEntitlementTemplates) | **Get** /entitlements/templates | List and search entitlement templates
[**GetUserInventories**](UsersInventoryApi.md#GetUserInventories) | **Get** /users/{id}/inventory | List the user inventory entries for a given user
[**GetUserInventory**](UsersInventoryApi.md#GetUserInventory) | **Get** /users/{user_id}/inventory/{id} | Get an inventory entry
[**GetUserInventoryLog**](UsersInventoryApi.md#GetUserInventoryLog) | **Get** /users/{user_id}/inventory/{id}/log | List the log entries for this inventory entry
[**GetUsersInventory**](UsersInventoryApi.md#GetUsersInventory) | **Get** /inventories | List the user inventory entries for all users
[**GrantUserEntitlement**](UsersInventoryApi.md#GrantUserEntitlement) | **Post** /users/{user_id}/entitlements | Grant an entitlement
[**UpdateEntitlementItem**](UsersInventoryApi.md#UpdateEntitlementItem) | **Put** /entitlements/{entitlement_id} | Update an entitlement item
[**UpdateEntitlementTemplate**](UsersInventoryApi.md#UpdateEntitlementTemplate) | **Put** /entitlements/templates/{id} | Update an entitlement template
[**UpdateUserInventoryBehaviorData**](UsersInventoryApi.md#UpdateUserInventoryBehaviorData) | **Put** /users/{user_id}/inventory/{id}/behavior-data | Set the behavior data for an inventory entry
[**UpdateUserInventoryExpires**](UsersInventoryApi.md#UpdateUserInventoryExpires) | **Put** /users/{user_id}/inventory/{id}/expires | Set the expiration date
[**UpdateUserInventoryStatus**](UsersInventoryApi.md#UpdateUserInventoryStatus) | **Put** /users/{user_id}/inventory/{id}/status | Set the status for an inventory entry
[**UseUserEntitlementItem**](UsersInventoryApi.md#UseUserEntitlementItem) | **Post** /users/{user_id}/entitlements/{item_id}/use | Use an item


# **AddItemToUserInventory**
> InvoiceResource AddItemToUserInventory($id, $userInventoryAddRequest)

Adds an item to the user inventory

The inventory is fulfilled asynchronously UNLESS the invoice is explicitely skipped. Depending on the use case, it might require the client to verify that the entitlement was added after the fact or configure a BRE rule to get a notification in real time


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the user | 
 **userInventoryAddRequest** | [**UserInventoryAddRequest**](UserInventoryAddRequest.md)| The user inventory add request object | [optional] 

### Return type

[**InvoiceResource**](InvoiceResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckUserEntitlementItem**
> CheckUserEntitlementItem($userId, $itemId, $sku)

Check for access to an item without consuming

Useful for pre-check and accounts for all various buisness rules


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| The id of the user to check for or &#39;me&#39; for logged in user | 
 **itemId** | **int32**| The id of the item | 
 **sku** | **string**| The specific sku of an entitlement list addition to check entitlement for. This is of very limited and specific use and should generally be left out | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEntitlementItem**
> EntitlementItem CreateEntitlementItem($entitlementItem)

Create an entitlement item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitlementItem** | [**EntitlementItem**](EntitlementItem.md)| The entitlement item object | [optional] 

### Return type

[**EntitlementItem**](EntitlementItem.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEntitlementTemplate**
> ItemTemplateResource CreateEntitlementTemplate($template)

Create an entitlement template

Entitlement templates define a type of entitlement and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **template** | [**ItemTemplateResource**](ItemTemplateResource.md)| The entitlement template to be created | [optional] 

### Return type

[**ItemTemplateResource**](ItemTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEntitlementItem**
> DeleteEntitlementItem($entitlementId)

Delete an entitlement item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitlementId** | **int32**| The id of the entitlement | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEntitlementTemplate**
> DeleteEntitlementTemplate($id, $cascade)

Delete an entitlement template

If cascade = 'detach', it will force delete the template even if it's attached to other objects


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| The value needed to delete used templates | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntitlementItem**
> EntitlementItem GetEntitlementItem($entitlementId)

Get a single entitlement item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitlementId** | **int32**| The id of the entitlement | 

### Return type

[**EntitlementItem**](EntitlementItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntitlementItems**
> PageResourceEntitlementItem GetEntitlementItems($size, $page, $order)

List and search entitlement items


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceEntitlementItem**](PageResource«EntitlementItem».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntitlementTemplate**
> ItemTemplateResource GetEntitlementTemplate($id)

Get a single entitlement template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 

### Return type

[**ItemTemplateResource**](ItemTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntitlementTemplates**
> PageResourceItemTemplateResource GetEntitlementTemplates($size, $page, $order)

List and search entitlement templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceItemTemplateResource**](PageResource«ItemTemplateResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserInventories**
> PageResourceUserInventoryResource GetUserInventories($id, $inactive, $size, $page, $filterItemName, $filterMinDate, $filterMaxDate)

List the user inventory entries for a given user


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the user | 
 **inactive** | **bool**| If true, accepts inactive user inventories | [optional] [default to false]
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **filterItemName** | **string**| Filter by items whose name starts with a string | [optional] 
 **filterMinDate** | **int64**| Filter for inventory added after the specified date, unix timestamp in seconds | [optional] 
 **filterMaxDate** | **int64**| Filter for inventory added before the specified date, unix timestamp in seconds | [optional] 

### Return type

[**PageResourceUserInventoryResource**](PageResource«UserInventoryResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserInventory**
> UserInventoryResource GetUserInventory($userId, $id)

Get an inventory entry


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the inventory owner or &#39;me&#39; for the logged in user | 
 **id** | **int32**| The id of the user inventory | 

### Return type

[**UserInventoryResource**](UserInventoryResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserInventoryLog**
> PageResourceUserItemLogResource GetUserInventoryLog($userId, $id, $size, $page)

List the log entries for this inventory entry


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| The id of the inventory owner or &#39;me&#39; for the logged in user | 
 **id** | **int32**| The id of the user inventory | 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageResourceUserItemLogResource**](PageResource«UserItemLogResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersInventory**
> PageResourceUserInventoryResource GetUsersInventory($inactive, $size, $page, $filterItemName, $filterUsername, $filterGroup, $filterDate)

List the user inventory entries for all users


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inactive** | **bool**| If true, accepts inactive user inventories | [optional] [default to false]
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **filterItemName** | **string**| Filter by items whose name starts with a string | [optional] 
 **filterUsername** | **string**| Filter by entries owned by the user with the specified username | [optional] 
 **filterGroup** | **string**| Filter by entries owned by the users in a given group, by unique name | [optional] 
 **filterDate** | **string**| A comma separated string without spaces.  First value is the operator to search on, second value is the log start date, a unix timestamp in seconds. Can be repeated for a range, eg: GT,123,LT,456  Allowed operators: (GT, LT, EQ, GOE, LOE). | [optional] 

### Return type

[**PageResourceUserInventoryResource**](PageResource«UserInventoryResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrantUserEntitlement**
> GrantUserEntitlement($userId, $grantRequest)

Grant an entitlement


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user to grant the entitlement to | 
 **grantRequest** | [**EntitlementGrantRequest**](EntitlementGrantRequest.md)| grantRequest | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEntitlementItem**
> UpdateEntitlementItem($entitlementId, $entitlementItem)

Update an entitlement item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitlementId** | **int32**| The id of the entitlement | 
 **entitlementItem** | [**EntitlementItem**](EntitlementItem.md)| The entitlement item object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEntitlementTemplate**
> UpdateEntitlementTemplate($id, $template)

Update an entitlement template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **template** | [**ItemTemplateResource**](ItemTemplateResource.md)| The updated template | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserInventoryBehaviorData**
> UpdateUserInventoryBehaviorData($userId, $id, $data)

Set the behavior data for an inventory entry


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 
 **id** | **int32**| The id of the user inventory | 
 **data** | [**interface{}**](interface{}.md)| The data map | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserInventoryExpires**
> UpdateUserInventoryExpires($userId, $id, $timestamp)

Set the expiration date

Will change the current grace period for a subscription but not the bill date (possibly even ending before having the chance to re-bill)


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| user_id | 
 **id** | **int32**| The id of the user inventory | 
 **timestamp** | **int64**| The new expiration date as a unix timestamp in seconds. May be null (no body). | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserInventoryStatus**
> UpdateUserInventoryStatus($userId, $id, $inventoryStatus)

Set the status for an inventory entry


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32**| The id of the user | 
 **id** | **int32**| The id of the user inventory | 
 **inventoryStatus** | **string**| The inventory status object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UseUserEntitlementItem**
> UseUserEntitlementItem($userId, $itemId, $sku, $info)

Use an item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| The id of the user to check for or &#39;me&#39; for logged in user | 
 **itemId** | **int32**| The id of the item | 
 **sku** | **string**| The specific sku of an entitlement_list addition to check entitlement for. This is of very limited and specific use and should generally be left out | [optional] 
 **info** | **string**| Any additional info to add to the log about this use | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

