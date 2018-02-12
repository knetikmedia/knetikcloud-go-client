# \LocationsApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCountries**](LocationsApi.md#GetCountries) | **Get** /location/countries | Get a list of countries
[**GetCountryByGeoLocation**](LocationsApi.md#GetCountryByGeoLocation) | **Get** /location/geolocation/country | Get the iso3 code of your country
[**GetCountryStates**](LocationsApi.md#GetCountryStates) | **Get** /location/countries/{country_code_iso3}/states | Get a list of a country&#39;s states
[**GetCurrencyByGeoLocation**](LocationsApi.md#GetCurrencyByGeoLocation) | **Get** /location/geolocation/currency | Get the currency information of your country


# **GetCountries**
> []CountryResource GetCountries(ctx, ctx, )
Get a list of countries

<b>Permissions Needed:</b> ANY

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]CountryResource**](CountryResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountryByGeoLocation**
> string GetCountryByGeoLocation(ctx, ctx, )
Get the iso3 code of your country

Determined by geo ip location. <br><br><b>Permissions Needed:</b> ANY

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountryStates**
> []StateResource GetCountryStates(ctx, ctx, countryCodeIso3)
Get a list of a country's states

<b>Permissions Needed:</b> ANY

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **countryCodeIso3** | **string**| The iso3 code of the country | 

### Return type

[**[]StateResource**](StateResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrencyByGeoLocation**
> CurrencyResource GetCurrencyByGeoLocation(ctx, ctx, )
Get the currency information of your country

Determined by geo ip location, currency to country mapping and a fallback setting. <br><br><b>Permissions Needed:</b> ANY

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CurrencyResource**](CurrencyResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

