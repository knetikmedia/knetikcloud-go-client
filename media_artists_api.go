/* 
 * Knetik Platform API Documentation latest 
 *
 * This is the spec for the Knetik API.  Use this in conjunction with the documentation found at https://knetikcloud.com.
 *
 * OpenAPI spec version: latest 
 * Contact: support@knetik.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

import (
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type Media_ArtistsApiService service


/* Media_ArtistsApiService Adds a new artist in the system
 Adds a new artist in the system. Use specific media contributions endpoint to add contributions. &lt;br&gt;&lt;br&gt;&lt;b&gt;Permissions Needed:&lt;/b&gt; ANY
 * @param ctx context.Context Authentication Context 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "artistResource" (ArtistResource) The new artist
 @return ArtistResource*/
func (a *Media_ArtistsApiService) AddArtist(ctx context.Context, localVarOptionals map[string]interface{}) (ArtistResource,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ArtistResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/media/artists"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["artistResource"].(ArtistResource); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* Media_ArtistsApiService Create an artist template
 Artist Templates define a type of artist and the properties they have. &lt;br&gt;&lt;br&gt;&lt;b&gt;Permissions Needed:&lt;/b&gt; TEMPLATE_ADMIN
 * @param ctx context.Context Authentication Context 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "artistTemplateResource" (TemplateResource) The artist template resource object
 @return TemplateResource*/
func (a *Media_ArtistsApiService) CreateArtistTemplate(ctx context.Context, localVarOptionals map[string]interface{}) (TemplateResource,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  TemplateResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/media/artists/templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["artistTemplateResource"].(TemplateResource); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* Media_ArtistsApiService Removes an artist from the system IF no resources are attached to it
 &lt;b&gt;Permissions Needed:&lt;/b&gt; ARTISTS_ADMIN
 * @param ctx context.Context Authentication Context 
 @param id The artist id
 @return */
func (a *Media_ArtistsApiService) DeleteArtist(ctx context.Context, id int64) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/media/artists/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }

	return localVarHttpResponse, err
}

/* Media_ArtistsApiService Delete an artist template
 If cascade &#x3D; &#39;detach&#39;, it will force delete the template even if it&#39;s attached to other objects. &lt;br&gt;&lt;br&gt;&lt;b&gt;Permissions Needed:&lt;/b&gt; TEMPLATE_ADMIN
 * @param ctx context.Context Authentication Context 
 @param id The id of the template
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "cascade" (string) The value needed to delete used templates
 @return */
func (a *Media_ArtistsApiService) DeleteArtistTemplate(ctx context.Context, id string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/media/artists/templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["cascade"], "string", "cascade"); err != nil {
		return nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["cascade"].(string); localVarOk {
		localVarQueryParams.Add("cascade", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }

	return localVarHttpResponse, err
}

/* Media_ArtistsApiService Loads a specific artist details
 &lt;b&gt;Permissions Needed:&lt;/b&gt; ANY
 * @param ctx context.Context Authentication Context 
 @param id The artist id
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "showContributions" (int32) The number of contributions to show fetch
 @return ArtistResource*/
func (a *Media_ArtistsApiService) GetArtist(ctx context.Context, id int64, localVarOptionals map[string]interface{}) (ArtistResource,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ArtistResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/media/artists/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["showContributions"], "int32", "showContributions"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["showContributions"].(int32); localVarOk {
		localVarQueryParams.Add("show_contributions", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* Media_ArtistsApiService Get a single artist template
 &lt;b&gt;Permissions Needed:&lt;/b&gt; TEMPLATE_ADMIN or ARTISTS_ADMIN
 * @param ctx context.Context Authentication Context 
 @param id The id of the template
 @return TemplateResource*/
func (a *Media_ArtistsApiService) GetArtistTemplate(ctx context.Context, id string) (TemplateResource,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  TemplateResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/media/artists/templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* Media_ArtistsApiService List and search artist templates
 &lt;b&gt;Permissions Needed:&lt;/b&gt; TEMPLATE_ADMIN or ARTISTS_ADMIN
 * @param ctx context.Context Authentication Context 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "size" (int32) The number of objects returned per page
     @param "page" (int32) The number of the page returned, starting with 1
     @param "order" (string) A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC]
 @return PageResourceTemplateResource*/
func (a *Media_ArtistsApiService) GetArtistTemplates(ctx context.Context, localVarOptionals map[string]interface{}) (PageResourceTemplateResource,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PageResourceTemplateResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/media/artists/templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["size"], "int32", "size"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["page"], "int32", "page"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["order"], "string", "order"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["size"].(int32); localVarOk {
		localVarQueryParams.Add("size", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["page"].(int32); localVarOk {
		localVarQueryParams.Add("page", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["order"].(string); localVarOk {
		localVarQueryParams.Add("order", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* Media_ArtistsApiService Search for artists
 &lt;b&gt;Permissions Needed:&lt;/b&gt; ANY
 * @param ctx context.Context Authentication Context 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "filterArtistsByName" (string) Filter for artists which name *STARTS* with the given string
     @param "size" (int32) The number of objects returned per page
     @param "page" (int32) The number of the page returned, starting with 1
     @param "order" (string) A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC]
 @return PageResourceArtistResource*/
func (a *Media_ArtistsApiService) GetArtists(ctx context.Context, localVarOptionals map[string]interface{}) (PageResourceArtistResource,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PageResourceArtistResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/media/artists"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["filterArtistsByName"], "string", "filterArtistsByName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["size"], "int32", "size"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["page"], "int32", "page"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["order"], "string", "order"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["filterArtistsByName"].(string); localVarOk {
		localVarQueryParams.Add("filter_artists_by_name", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["size"].(int32); localVarOk {
		localVarQueryParams.Add("size", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["page"].(int32); localVarOk {
		localVarQueryParams.Add("page", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["order"].(string); localVarOk {
		localVarQueryParams.Add("order", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* Media_ArtistsApiService Modifies an artist details
 &lt;b&gt;Permissions Needed:&lt;/b&gt; ARTISTS_ADMIN
 * @param ctx context.Context Authentication Context 
 @param id The artist id
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "artistResource" (ArtistResource) The new artist
 @return */
func (a *Media_ArtistsApiService) UpdateArtist(ctx context.Context, id int64, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/media/artists/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["artistResource"].(ArtistResource); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }

	return localVarHttpResponse, err
}

/* Media_ArtistsApiService Update an artist template
 &lt;b&gt;Permissions Needed:&lt;/b&gt; TEMPLATE_ADMIN
 * @param ctx context.Context Authentication Context 
 @param id The id of the template
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "artistTemplateResource" (TemplateResource) The artist template resource object
 @return TemplateResource*/
func (a *Media_ArtistsApiService) UpdateArtistTemplate(ctx context.Context, id string, localVarOptionals map[string]interface{}) (TemplateResource,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  TemplateResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/media/artists/templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["artistTemplateResource"].(TemplateResource); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

