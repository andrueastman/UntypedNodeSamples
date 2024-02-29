package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2 "github.com/andrueastman/untypednodesamples/goclient/models"
)

// ProviderItemRequestBuilder builds and executes requests for operations under \{provider-id}
type ProviderItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ProviderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProviderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewProviderItemRequestBuilderInternal instantiates a new ProviderItemRequestBuilder and sets the default values.
func NewProviderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProviderItemRequestBuilder) {
    m := &ProviderItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/{provider%2Did}", pathParameters),
    }
    return m
}
// NewProviderItemRequestBuilder instantiates a new ProviderItemRequestBuilder and sets the default values.
func NewProviderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProviderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProviderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all APIs in the directory for a particular providerNameReturns links to the individual API entry for each API.
// returns a APIsable when successful
func (m *ProviderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ProviderItemRequestBuilderGetRequestConfiguration)(iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2.APIsable, error) {
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
// ServicesJson the servicesJson property
// returns a *ServicesJsonRequestBuilder when successful
func (m *ProviderItemRequestBuilder) ServicesJson()(*ServicesJsonRequestBuilder) {
    return NewServicesJsonRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation list all APIs in the directory for a particular providerNameReturns links to the individual API entry for each API.
// returns a *RequestInformation when successful
func (m *ProviderItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProviderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ProviderItemRequestBuilder when successful
func (m *ProviderItemRequestBuilder) WithUrl(rawUrl string)(*ProviderItemRequestBuilder) {
    return NewProviderItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
