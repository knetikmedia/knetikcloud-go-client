# \Store_BundlesApi

All URIs are relative to *https://jsapi-integration.us-east-1.elasticbeanstalk.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBundleItem**](Store_BundlesApi.md#CreateBundleItem) | **Post** /store/bundles | Create a bundle item
[**CreateBundleTemplate**](Store_BundlesApi.md#CreateBundleTemplate) | **Post** /store/bundles/templates | Create a bundle template
[**DeleteBundleItem**](Store_BundlesApi.md#DeleteBundleItem) | **Delete** /store/bundles/{id} | Delete a bundle item
[**DeleteBundleTemplate**](Store_BundlesApi.md#DeleteBundleTemplate) | **Delete** /store/bundles/templates/{id} | Delete a bundle template
[**GetBundleItem**](Store_BundlesApi.md#GetBundleItem) | **Get** /store/bundles/{id} | Get a single bundle item
[**GetBundleTemplate**](Store_BundlesApi.md#GetBundleTemplate) | **Get** /store/bundles/templates/{id} | Get a single bundle template
[**GetBundleTemplates**](Store_BundlesApi.md#GetBundleTemplates) | **Get** /store/bundles/templates | List and search bundle templates
[**UpdateBundleItem**](Store_BundlesApi.md#UpdateBundleItem) | **Put** /store/bundles/{id} | Update a bundle item
[**UpdateBundleTemplate**](Store_BundlesApi.md#UpdateBundleTemplate) | **Put** /store/bundles/templates/{id} | Update a bundle template


# **CreateBundleItem**
> BundleItem CreateBundleItem(ctx, ctx, optional)
Create a bundle item

The SKU for the bundle itself must be unique and there can only be one SKU.  Extra notes for price_override:  The price of all the items (multiplied by the quantity) must equal the price of the bundle.  With individual prices set, items will be processed individually and can be refunded as such.  However, if all prices are set to null, the price of the bundle will be used and will be treated as one item. <br><br><b>Permissions Needed:</b> BUNDLES_ADMIN

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
 **cascade** | **bool**| Whether to cascade group changes, such as in the limited gettable behavior. A 400 error will return otherwise if the group is already in use with different values. | [default to false]
 **bundleItem** | [**BundleItem**](BundleItem.md)| The bundle item object | 

### Return type

[**BundleItem**](BundleItem.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBundleTemplate**
> ItemTemplateResource CreateBundleTemplate(ctx, ctx, optional)
Create a bundle template

Bundle Templates define a type of bundle and the properties they have. <br><br><b>Permissions Needed:</b> BUNDLES_ADMIN

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
 **bundleTemplateResource** | [**ItemTemplateResource**](ItemTemplateResource.md)| The new bundle template | 

### Return type

[**ItemTemplateResource**](ItemTemplateResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBundleItem**
> DeleteBundleItem(ctx, ctx, id)
Delete a bundle item

<b>Permissions Needed:</b> BUNDLES_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int32**| The id of the bundle | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBundleTemplate**
> DeleteBundleTemplate(ctx, ctx, id, optional)
Delete a bundle template

<b>Permissions Needed:</b> BUNDLES_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the template | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| force deleting the template if it&#39;s attached to other objects, cascade &#x3D; detach | 

### Return type

 (empty response body)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleItem**
> BundleItem GetBundleItem(ctx, ctx, id)
Get a single bundle item

<b>Permissions Needed:</b> ANY

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int32**| The id of the bundle | 

### Return type

[**BundleItem**](BundleItem.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleTemplate**
> ItemTemplateResource GetBundleTemplate(ctx, ctx, id)
Get a single bundle template

Bundle Templates define a type of bundle and the properties they have. <br><br><b>Permissions Needed:</b> ANY

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the template | 

### Return type

[**ItemTemplateResource**](ItemTemplateResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleTemplates**
> PageResourceItemTemplateResource GetBundleTemplates(ctx, ctx, optional)
List and search bundle templates

<b>Permissions Needed:</b> ANY

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
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to id:ASC]

### Return type

[**PageResourceItemTemplateResource**](PageResource«ItemTemplateResource».md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBundleItem**
> BundleItem UpdateBundleItem(ctx, ctx, id, optional)
Update a bundle item

<b>Permissions Needed:</b> BUNDLES_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int32**| The id of the bundle | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the bundle | 
 **cascade** | **bool**| Whether to cascade group changes, such as in the limited gettable behavior. A 400 error will return otherwise if the group is already in use with different values. | [default to false]
 **bundleItem** | [**BundleItem**](BundleItem.md)| The bundle item object | 

### Return type

[**BundleItem**](BundleItem.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBundleTemplate**
> ItemTemplateResource UpdateBundleTemplate(ctx, ctx, id, optional)
Update a bundle template

<b>Permissions Needed:</b> BUNDLES_ADMIN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the template | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **bundleTemplateResource** | [**ItemTemplateResource**](ItemTemplateResource.md)| The bundle template resource object | 

### Return type

[**ItemTemplateResource**](ItemTemplateResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

