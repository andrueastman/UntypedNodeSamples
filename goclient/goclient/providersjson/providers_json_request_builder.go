package providersjson

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ProvidersJsonRequestBuilder builds and executes requests for operations under \providers.json
type ProvidersJsonRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ProvidersJsonRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProvidersJsonRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewProvidersJsonRequestBuilderInternal instantiates a new ProvidersJsonRequestBuilder and sets the default values.
func NewProvidersJsonRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProvidersJsonRequestBuilder) {
    m := &ProvidersJsonRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/providers.json", pathParameters),
    }
    return m
}
// NewProvidersJsonRequestBuilder instantiates a new ProvidersJsonRequestBuilder and sets the default values.
func NewProvidersJsonRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProvidersJsonRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProvidersJsonRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all the providers in the directory
// Deprecated: This method is obsolete. Use GetAsProvidersGetResponse instead.
// returns a ProvidersResponseable when successful
func (m *ProvidersJsonRequestBuilder) Get(ctx context.Context, requestConfiguration *ProvidersJsonRequestBuilderGetRequestConfiguration)(ProvidersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProvidersResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ProvidersResponseable), nil
}
// GetAsProvidersGetResponse list all the providers in the directory
// returns a ProvidersGetResponseable when successful
func (m *ProvidersJsonRequestBuilder) GetAsProvidersGetResponse(ctx context.Context, requestConfiguration *ProvidersJsonRequestBuilderGetRequestConfiguration)(ProvidersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProvidersGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ProvidersGetResponseable), nil
}
// ToGetRequestInformation list all the providers in the directory
// returns a *RequestInformation when successful
func (m *ProvidersJsonRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProvidersJsonRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ProvidersJsonRequestBuilder when successful
func (m *ProvidersJsonRequestBuilder) WithUrl(rawUrl string)(*ProvidersJsonRequestBuilder) {
    return NewProvidersJsonRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
