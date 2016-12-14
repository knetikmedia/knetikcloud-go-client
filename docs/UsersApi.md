# \UsersApi

All URIs are relative to *https://devsandbox.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserTemplateUsingPOST**](UsersApi.md#CreateUserTemplateUsingPOST) | **Post** /users/templates | Create a user template
[**DeleteUserTemplateUsingDELETE**](UsersApi.md#DeleteUserTemplateUsingDELETE) | **Delete** /users/templates/{id} | Delete a user template
[**DoPasswordResetUsingPUT**](UsersApi.md#DoPasswordResetUsingPUT) | **Put** /users/{id}/password-reset | Choose a new password after a reset
[**GetUserTemplateUsingGET**](UsersApi.md#GetUserTemplateUsingGET) | **Get** /users/templates/{id} | Get a single user template
[**GetUserTemplatesUsingGET**](UsersApi.md#GetUserTemplatesUsingGET) | **Get** /users/templates | List and search user templates
[**GetUserUsingGET**](UsersApi.md#GetUserUsingGET) | **Get** /users/{id} | Get a single user
[**GetUsersUsingGET**](UsersApi.md#GetUsersUsingGET) | **Get** /users | List and search users
[**InitiatePasswordResetUsingPOST**](UsersApi.md#InitiatePasswordResetUsingPOST) | **Post** /users/{id}/password-reset | Reset a user&#39;s password
[**RegisterUserUsingPOST**](UsersApi.md#RegisterUserUsingPOST) | **Post** /users | Register a new user
[**SetPasswordUsingPUT**](UsersApi.md#SetPasswordUsingPUT) | **Put** /users/{id}/password | Set a user&#39;s password
[**UpdateUserTemplateUsingPUT**](UsersApi.md#UpdateUserTemplateUsingPUT) | **Put** /users/templates/{id} | Update a user template
[**UpdateUserUsingPUT**](UsersApi.md#UpdateUserUsingPUT) | **Put** /users/{id} | Update a user&#39;s info


# **CreateUserTemplateUsingPOST**
> TemplateResource CreateUserTemplateUsingPOST($userTemplateResource)

Create a user template

User Templates define a type of user and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userTemplateResource** | [**TemplateResource**](TemplateResource.md)| The user template resource object | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserTemplateUsingDELETE**
> DeleteUserTemplateUsingDELETE($id, $cascade)

Delete a user template

If cascade = 'detach', it will force delete the template even if it's attached to other objects


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| The value needed to delete used templates | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoPasswordResetUsingPUT**
> DoPasswordResetUsingPUT($id, $newPasswordRequest)

Choose a new password after a reset

Finish resetting a user's password using the secret provided from the password-reset endpoint.  Password should be in plain text and will be encrypted on receipt. Use SSL for security.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the user | 
 **newPasswordRequest** | [**NewPasswordRequest**](NewPasswordRequest.md)| The new password request object | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserTemplateUsingGET**
> TemplateResource GetUserTemplateUsingGET($id)

Get a single user template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserTemplatesUsingGET**
> PageTemplateResource GetUserTemplatesUsingGET($size, $page, $order)

List and search user templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| a comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to 1]

### Return type

[**PageTemplateResource**](Page«TemplateResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserUsingGET**
> UserResource GetUserUsingGET($id)

Get a single user

Additional private info is included as USERS_ADMIN


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the user or &#39;me&#39; | 

### Return type

[**UserResource**](UserResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersUsingGET**
> PageUserBaseResource GetUsersUsingGET($filterRole, $filterDisplayname, $filterEmail, $filterFirstname, $filterFullname, $filterLastname, $filterUsername, $size, $page, $order)

List and search users

Additional private info is included as USERS_ADMIN


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterRole** | **string**| Filter for users whose roles has the provided role | [optional] 
 **filterDisplayname** | **string**| Filter for users whose display name starts with provided string. | [optional] 
 **filterEmail** | **string**| Filter for users whose email starts with provided string. Requires USERS_ADMIN permission | [optional] 
 **filterFirstname** | **string**| Filter for users whose first name starts with provided string. Requires USERS_ADMIN permission | [optional] 
 **filterFullname** | **string**| Filter for users whose full name starts with provided string. Requires USERS_ADMIN permission | [optional] 
 **filterLastname** | **string**| Filter for users whose last name starts with provided string. Requires USERS_ADMIN permission | [optional] 
 **filterUsername** | **string**| Filter for users whose username starts with the provided string. Requires USERS_ADMIN permission | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to 1]

### Return type

[**PageUserBaseResource**](Page«UserBaseResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitiatePasswordResetUsingPOST**
> InitiatePasswordResetUsingPOST($id)

Reset a user's password

A reset code will be generated and a 'forgot_password' BRE event will be fired with that code; this can be routed to the user as needed and submitted to the 'choose a new password' endpoint.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the user | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterUserUsingPOST**
> UserResource RegisterUserUsingPOST($userResource)

Register a new user

Password should be in plain text and will be encrypted on receipt. Use SSL for security


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userResource** | [**UserResource**](UserResource.md)| The user resource object | [optional] 

### Return type

[**UserResource**](UserResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPasswordUsingPUT**
> SetPasswordUsingPUT($id, $password)

Set a user's password

Password should be in plain text and will be encrypted on receipt. Use SSL for security.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the user | 
 **password** | **string**| The new plain text password | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserTemplateUsingPUT**
> UpdateUserTemplateUsingPUT($id, $userTemplateResource)

Update a user template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **userTemplateResource** | [**TemplateResource**](TemplateResource.md)| The user template resource object | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserUsingPUT**
> UpdateUserUsingPUT($id, $userResource)

Update a user's info

Password will not be edited on this endpoint, use password specific endpoints.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the user or &#39;me&#39; | 
 **userResource** | [**UserResource**](UserResource.md)| The user resource object | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

