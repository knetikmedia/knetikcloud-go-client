# \UsersFriendshipsApi

All URIs are relative to *https://integration.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFriendUsingPOST**](UsersFriendshipsApi.md#AddFriendUsingPOST) | **Post** /users/{user_id}/friends/{id} | Add a friend
[**ConnectUsingTokenUsingPOST**](UsersFriendshipsApi.md#ConnectUsingTokenUsingPOST) | **Post** /users/{user_id}/friends/tokens | Redeem friendship token
[**GetFriendsUsingGET**](UsersFriendshipsApi.md#GetFriendsUsingGET) | **Get** /users/{user_id}/friends | Get friends list
[**GetInviteTokenUsingGET**](UsersFriendshipsApi.md#GetInviteTokenUsingGET) | **Get** /users/{user_id}/invite-token | Returns the invite token
[**GetInvitesUsingGET**](UsersFriendshipsApi.md#GetInvitesUsingGET) | **Get** /users/{user_id}/invites | Get pending invites
[**RemoveFriendUsingDELETE**](UsersFriendshipsApi.md#RemoveFriendUsingDELETE) | **Delete** /users/{user_id}/friends/{id} | Remove or decline a friend


# **AddFriendUsingPOST**
> AddFriendUsingPOST($userId, $id)

Add a friend

As a user, either creates or confirm a pending request. As an admin, call this endpoint twice while inverting the IDs to create a confirmed friendship.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| The id of the user or &#39;me&#39; if logged in | 
 **id** | **int32**| The id of the user to befriend | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConnectUsingTokenUsingPOST**
> ConnectUsingTokenUsingPOST($userId, $token)

Redeem friendship token

Immediately connects the requested user with the user mapped by the provided invite token


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| The id of the user or &#39;me&#39; if logged in | 
 **token** | **string**| The invite token | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFriendsUsingGET**
> PageResourceSimpleUserResource GetFriendsUsingGET($userId, $size, $page, $order)

Get friends list


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| The id of the user or &#39;me&#39; | 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceSimpleUserResource**](PageResource«SimpleUserResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInviteTokenUsingGET**
> string GetInviteTokenUsingGET($userId)

Returns the invite token

This is a unique invite token that allows direct connection to the request user.  Exposing that token presents privacy issues if the token is leaked. Use friend request flow instead if confirmation is required


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| The id of the user or &#39;me&#39; if logged in | 

### Return type

**string**

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvitesUsingGET**
> PageResourceSimpleUserResource GetInvitesUsingGET($userId, $size, $page, $order)

Get pending invites

Invites that the specified user received


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| The id of the user or &#39;me&#39; | 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceSimpleUserResource**](PageResource«SimpleUserResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveFriendUsingDELETE**
> RemoveFriendUsingDELETE($userId, $id)

Remove or decline a friend


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| The id of the user or &#39;me&#39; if logged in | 
 **id** | **int32**| The id of the user to befriend | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

