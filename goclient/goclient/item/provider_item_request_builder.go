package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    id6760cbde1c5ad1214c40ac9ca069164ef0f31f8ecb790d183bca1f74a71e629 "github.com/andrueastman/untypednodesamples/goclient/goclient/models"
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
func (m *ProviderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ProviderItemRequestBuilderGetRequestConfiguration)(id6760cbde1c5ad1214c40ac9ca069164ef0f31f8ecb790d183bca1f74a71e629.APIsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id6760cbde1c5ad1214c40ac9ca069164ef0f31f8ecb790d183bca1f74a71e629.CreateAPIsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id6760cbde1c5ad1214c40ac9ca069164ef0f31f8ecb790d183bca1f74a71e629.APIsable), nil
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
