# \TaxesApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCountryTax**](TaxesApi.md#CreateCountryTax) | **Post** /tax/countries | Create a country tax
[**CreateStateTax**](TaxesApi.md#CreateStateTax) | **Post** /tax/countries/{country_code_iso3}/states | Create a state tax
[**DeleteCountryTax**](TaxesApi.md#DeleteCountryTax) | **Delete** /tax/countries/{country_code_iso3} | Delete an existing tax
[**DeleteStateTax**](TaxesApi.md#DeleteStateTax) | **Delete** /tax/countries/{country_code_iso3}/states/{state_code} | Delete an existing state tax
[**GetCountryTax**](TaxesApi.md#GetCountryTax) | **Get** /tax/countries/{country_code_iso3} | Get a single tax
[**GetCountryTaxes**](TaxesApi.md#GetCountryTaxes) | **Get** /tax/countries | List and search taxes
[**GetStateTax**](TaxesApi.md#GetStateTax) | **Get** /tax/countries/{country_code_iso3}/states/{state_code} | Get a single state tax
[**GetStateTaxesForCountries**](TaxesApi.md#GetStateTaxesForCountries) | **Get** /tax/states | List and search taxes across all countries
[**GetStateTaxesForCountry**](TaxesApi.md#GetStateTaxesForCountry) | **Get** /tax/countries/{country_code_iso3}/states | List and search taxes within a country
[**UpdateCountryTax**](TaxesApi.md#UpdateCountryTax) | **Put** /tax/countries/{country_code_iso3} | Create or update a tax
[**UpdateStateTax**](TaxesApi.md#UpdateStateTax) | **Put** /tax/countries/{country_code_iso3}/states/{state_code} | Create or update a state tax


# **CreateCountryTax**
> CountryTaxResource CreateCountryTax($taxResource)

Create a country tax


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxResource** | [**CountryTaxResource**](CountryTaxResource.md)| The tax object | [optional] 

### Return type

[**CountryTaxResource**](CountryTaxResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStateTax**
> StateTaxResource CreateStateTax($countryCodeIso3, $taxResource)

Create a state tax


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 
 **taxResource** | [**StateTaxResource**](StateTaxResource.md)| The tax object | [optional] 

### Return type

[**StateTaxResource**](StateTaxResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCountryTax**
> DeleteCountryTax($countryCodeIso3)

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
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStateTax**
> DeleteStateTax($countryCodeIso3, $stateCode)

Delete an existing state tax


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
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountryTax**
> CountryTaxResource GetCountryTax($countryCodeIso3)

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
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountryTaxes**
> PageResourceCountryTaxResource GetCountryTaxes($size, $page, $order)

List and search taxes

Get a list of taxes


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to name:ASC]

### Return type

[**PageResourceCountryTaxResource**](PageResource«CountryTaxResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStateTax**
> StateTaxResource GetStateTax($countryCodeIso3, $stateCode)

Get a single state tax


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
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStateTaxesForCountries**
> PageResourceStateTaxResource GetStateTaxesForCountries($size, $page, $order)

List and search taxes across all countries

Get a list of taxes


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] 

### Return type

[**PageResourceStateTaxResource**](PageResource«StateTaxResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStateTaxesForCountry**
> PageResourceStateTaxResource GetStateTaxesForCountry($countryCodeIso3, $size, $page, $order)

List and search taxes within a country

Get a list of taxes


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] 

### Return type

[**PageResourceStateTaxResource**](PageResource«StateTaxResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCountryTax**
> CountryTaxResource UpdateCountryTax($countryCodeIso3, $taxResource)

Create or update a tax


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 
 **taxResource** | [**CountryTaxResource**](CountryTaxResource.md)| The tax object | [optional] 

### Return type

[**CountryTaxResource**](CountryTaxResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStateTax**
> StateTaxResource UpdateStateTax($countryCodeIso3, $stateCode, $taxResource)

Create or update a state tax


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 
 **stateCode** | **string**| The code of the state | 
 **taxResource** | [**StateTaxResource**](StateTaxResource.md)| The tax object | [optional] 

### Return type

[**StateTaxResource**](StateTaxResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

