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
> CountryTaxResource CreateCountryTax(optional)
Create a country tax

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxResource** | [**CountryTaxResource**](CountryTaxResource.md)| The tax object | 

### Return type

[**CountryTaxResource**](CountryTaxResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStateTax**
> StateTaxResource CreateStateTax(countryCodeIso3, optional)
Create a state tax

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **countryCodeIso3** | **string**| The iso3 code of the country | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 
 **taxResource** | [**StateTaxResource**](StateTaxResource.md)| The tax object | 

### Return type

[**StateTaxResource**](StateTaxResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCountryTax**
> DeleteCountryTax(countryCodeIso3)
Delete an existing tax

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **countryCodeIso3** | **string**| The iso3 code of the country | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStateTax**
> DeleteStateTax(countryCodeIso3, stateCode)
Delete an existing state tax

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **countryCodeIso3** | **string**| The iso3 code of the country | 
  **stateCode** | **string**| The code of the state | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountryTax**
> CountryTaxResource GetCountryTax(countryCodeIso3)
Get a single tax

### Required Parameters

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
> PageResourceCountryTaxResource GetCountryTaxes(optional)
List and search taxes

Get a list of taxes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to name:ASC]

### Return type

[**PageResourceCountryTaxResource**](PageResource«CountryTaxResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStateTax**
> StateTaxResource GetStateTax(countryCodeIso3, stateCode)
Get a single state tax

### Required Parameters

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
> PageResourceStateTaxResource GetStateTaxesForCountries(optional)
List and search taxes across all countries

Get a list of taxes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | 

### Return type

[**PageResourceStateTaxResource**](PageResource«StateTaxResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStateTaxesForCountry**
> PageResourceStateTaxResource GetStateTaxesForCountry(countryCodeIso3, optional)
List and search taxes within a country

Get a list of taxes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **countryCodeIso3** | **string**| The iso3 code of the country | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | 

### Return type

[**PageResourceStateTaxResource**](PageResource«StateTaxResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCountryTax**
> CountryTaxResource UpdateCountryTax(countryCodeIso3, optional)
Create or update a tax

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **countryCodeIso3** | **string**| The iso3 code of the country | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 
 **taxResource** | [**CountryTaxResource**](CountryTaxResource.md)| The tax object | 

### Return type

[**CountryTaxResource**](CountryTaxResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStateTax**
> StateTaxResource UpdateStateTax(countryCodeIso3, stateCode, optional)
Create or update a state tax

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **countryCodeIso3** | **string**| The iso3 code of the country | 
  **stateCode** | **string**| The code of the state | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCodeIso3** | **string**| The iso3 code of the country | 
 **stateCode** | **string**| The code of the state | 
 **taxResource** | [**StateTaxResource**](StateTaxResource.md)| The tax object | 

### Return type

[**StateTaxResource**](StateTaxResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

