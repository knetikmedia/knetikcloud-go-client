# \LocationsApi

All URIs are relative to *https://sandbox.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCountries**](LocationsApi.md#GetCountries) | **Get** /location/countries | Get a list of countries
[**GetCountryByGeoLocation**](LocationsApi.md#GetCountryByGeoLocation) | **Get** /location/geolocation/country | Get the iso3 code of your country
[**GetCountryStates**](LocationsApi.md#GetCountryStates) | **Get** /location/countries/{country_code_iso3}/states | Get a list of a country&#39;s states
[**GetCurrencyByGeoLocation**](LocationsApi.md#GetCurrencyByGeoLocation) | **Get** /location/geolocation/currency | Get the currency information of your country


# **GetCountries**
> []CountryResource GetCountries()

Get a list of countries


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]CountryResource**](CountryResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountryByGeoLocation**
> string GetCountryByGeoLocation()

Get the iso3 code of your country

Determined by geo ip location


### Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountryStates**
> []StateResource GetCountryStates($countryCodeIso3)

Get a list of a country's states


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 

### Return type

[**[]StateResource**](StateResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrencyByGeoLocation**
> CurrencyResource GetCurrencyByGeoLocation()

Get the currency information of your country

Determined by geo ip location, currency to country mapping and a fallback setting


### Parameters
This endpoint does not need any parameter.

### Return type

[**CurrencyResource**](CurrencyResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

