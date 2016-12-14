# \LocationsApi

All URIs are relative to *https://devsandbox.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CitiesUsingGET**](LocationsApi.md#CitiesUsingGET) | **Get** /location/countries/{country_code_iso3}/states/{state_code}/cities | Get a list of a state&#39;s cities
[**CountriesUsingGET**](LocationsApi.md#CountriesUsingGET) | **Get** /location/countries | Get a list of countries
[**GetCurrencyByGeoLocationUsingGET**](LocationsApi.md#GetCurrencyByGeoLocationUsingGET) | **Get** /location/geolocation/currency | Get the currency information of your country
[**GetGeoLocationCountryUsingGET**](LocationsApi.md#GetGeoLocationCountryUsingGET) | **Get** /location/geolocation/country | Get the iso3 code of your country
[**StatesUsingGET**](LocationsApi.md#StatesUsingGET) | **Get** /location/countries/{country_code_iso3}/states | Get a list of a country&#39;s states


# **CitiesUsingGET**
> []CityResource CitiesUsingGET($countryCodeIso3, $stateCode)

Get a list of a state's cities


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 
 **stateCode** | **string**| The code of the state | 

### Return type

[**[]CityResource**](CityResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountriesUsingGET**
> []CountryResource CountriesUsingGET()

Get a list of countries


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]CountryResource**](CountryResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrencyByGeoLocationUsingGET**
> CurrencyResource GetCurrencyByGeoLocationUsingGET()

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
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGeoLocationCountryUsingGET**
> string GetGeoLocationCountryUsingGET()

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
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatesUsingGET**
> []StateResource StatesUsingGET($countryCodeIso3)

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
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

