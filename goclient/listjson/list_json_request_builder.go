package listjson

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2 "github.com/andrueastman/untypednodesamples/goclient/models"
)

// ListJsonRequestBuilder builds and executes requests for operations under \list.json
type ListJsonRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ListJsonRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListJsonRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewListJsonRequestBuilderInternal instantiates a new ListJsonRequestBuilder and sets the default values.
func NewListJsonRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListJsonRequestBuilder) {
    m := &ListJsonRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/list.json", pathParameters),
    }
    return m
}
// NewListJsonRequestBuilder instantiates a new ListJsonRequestBuilder and sets the default values.
func NewListJsonRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListJsonRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListJsonRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all APIs in the directory.Returns links to the OpenAPI definitions for each API in the directory.If API exist in multiple versions `preferred` one is explicitly marked.Some basic info from the OpenAPI definition is cached inside each object.This allows you to generate some simple views without needing to fetch the OpenAPI definition for each API.
// returns a APIsable when successful
func (m *ListJsonRequestBuilder) Get(ctx context.Context, requestConfiguration *ListJsonRequestBuilderGetRequestConfiguration)(iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2.APIsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2.CreateAPIsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2.APIsable), nil
}
// ToGetRequestInformation list all APIs in the directory.Returns links to the OpenAPI definitions for each API in the directory.If API exist in multiple versions `preferred` one is explicitly marked.Some basic info from the OpenAPI definition is cached inside each object.This allows you to generate some simple views without needing to fetch the OpenAPI definition for each API.
// returns a *RequestInformation when successful
func (m *ListJsonRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ListJsonRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ListJsonRequestBuilder when successful
func (m *ListJsonRequestBuilder) WithUrl(rawUrl string)(*ListJsonRequestBuilder) {
    return NewListJsonRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
