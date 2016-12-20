# \TaxesApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaxUsingPOST**](TaxesApi.md#CreateTaxUsingPOST) | **Post** /tax/countries/{country_code_iso3}/states | Create a tax
[**DeleteTaxUsingDELETE**](TaxesApi.md#DeleteTaxUsingDELETE) | **Delete** /tax/countries/{country_code_iso3} | Delete an existing tax
[**DeleteTaxUsingDELETE1**](TaxesApi.md#DeleteTaxUsingDELETE1) | **Delete** /tax/countries/{country_code_iso3}/states/{state_code} | Delete an existing tax
[**GetAllTaxesUsingGET**](TaxesApi.md#GetAllTaxesUsingGET) | **Get** /tax/states | List and search taxes across all countries
[**GetTaxUsingGET**](TaxesApi.md#GetTaxUsingGET) | **Get** /tax/countries/{country_code_iso3} | Get a single tax
[**GetTaxUsingGET1**](TaxesApi.md#GetTaxUsingGET1) | **Get** /tax/countries/{country_code_iso3}/states/{state_code} | Get a single tax
[**GetTaxesUsingGET**](TaxesApi.md#GetTaxesUsingGET) | **Get** /tax/countries | List and search taxes
[**GetTaxesUsingGET1**](TaxesApi.md#GetTaxesUsingGET1) | **Get** /tax/countries/{country_code_iso3}/states | List and search taxes within a country
[**UpdateTaxUsingPOST**](TaxesApi.md#UpdateTaxUsingPOST) | **Post** /tax/countries | Create a tax
[**UpdateTaxUsingPUT**](TaxesApi.md#UpdateTaxUsingPUT) | **Put** /tax/countries/{country_code_iso3} | Create or update a tax
[**UpdateTaxUsingPUT1**](TaxesApi.md#UpdateTaxUsingPUT1) | **Put** /tax/countries/{country_code_iso3}/states/{state_code} | Create or update a tax


# **CreateTaxUsingPOST**
> CreateTaxUsingPOST($countryCodeIso3, $taxResource)

Create a tax


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 
 **taxResource** | [**StateTaxResource**](StateTaxResource.md)| The tax object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTaxUsingDELETE**
> DeleteTaxUsingDELETE($countryCodeIso3)

Delete an existing tax


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTaxUsingDELETE1**
> DeleteTaxUsingDELETE1($countryCodeIso3, $stateCode)

Delete an existing tax


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 
 **stateCode** | **string**| The code of the state | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTaxesUsingGET**
> PageStateTaxResource GetAllTaxesUsingGET($size, $page, $order)

List and search taxes across all countries

Get a list of taxes


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to name:ASC]

### Return type

[**PageStateTaxResource**](Page«StateTaxResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaxUsingGET**
> CountryTaxResource GetTaxUsingGET($countryCodeIso3)

Get a single tax


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 

### Return type

[**CountryTaxResource**](CountryTaxResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaxUsingGET1**
> StateTaxResource GetTaxUsingGET1($countryCodeIso3, $stateCode)

Get a single tax


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 
 **stateCode** | **string**| The code of the state | 

### Return type

[**StateTaxResource**](StateTaxResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaxesUsingGET**
> PageCountryTaxResource GetTaxesUsingGET($size, $page, $order)

List and search taxes

Get a list of taxes


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to name:ASC]

### Return type

[**PageCountryTaxResource**](Page«CountryTaxResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaxesUsingGET1**
> PageStateTaxResource GetTaxesUsingGET1($countryCodeIso3, $size, $page, $order)

List and search taxes within a country

Get a list of taxes


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to name:ASC]

### Return type

[**PageStateTaxResource**](Page«StateTaxResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTaxUsingPOST**
> UpdateTaxUsingPOST($taxResource)

Create a tax


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxResource** | [**CountryTaxResource**](CountryTaxResource.md)| The tax object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTaxUsingPUT**
> UpdateTaxUsingPUT($countryCodeIso3, $taxResource)

Create or update a tax


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 
 **taxResource** | [**CountryTaxResource**](CountryTaxResource.md)| The tax object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTaxUsingPUT1**
> UpdateTaxUsingPUT1($countryCodeIso3, $stateCode, $taxResource)

Create or update a tax


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 
 **stateCode** | **string**| The code of the state | 
 **taxResource** | [**StateTaxResource**](StateTaxResource.md)| The tax object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

