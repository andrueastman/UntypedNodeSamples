package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ServicesJsonRequestBuilder builds and executes requests for operations under \{provider-id}\services.json
type ServicesJsonRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServicesJsonRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicesJsonRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServicesJsonRequestBuilderInternal instantiates a new ServicesJsonRequestBuilder and sets the default values.
func NewServicesJsonRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicesJsonRequestBuilder) {
    m := &ServicesJsonRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/{provider%2Did}/services.json", pathParameters),
    }
    return m
}
// NewServicesJsonRequestBuilder instantiates a new ServicesJsonRequestBuilder and sets the default values.
func NewServicesJsonRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicesJsonRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicesJsonRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all serviceNames in the directory for a particular providerName
// Deprecated: This method is obsolete. Use GetAsServicesGetResponse instead.
// returns a ServicesJsonServicesResponseable when successful
func (m *ServicesJsonRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicesJsonRequestBuilderGetRequestConfiguration)(ServicesJsonServicesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServicesJsonServicesResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServicesJsonServicesResponseable), nil
}
// GetAsServicesGetResponse list all serviceNames in the directory for a particular providerName
// returns a ServicesJsonServicesGetResponseable when successful
func (m *ServicesJsonRequestBuilder) GetAsServicesGetResponse(ctx context.Context, requestConfiguration *ServicesJsonRequestBuilderGetRequestConfiguration)(ServicesJsonServicesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServicesJsonServicesGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServicesJsonServicesGetResponseable), nil
}
// ToGetRequestInformation list all serviceNames in the directory for a particular providerName
// returns a *RequestInformation when successful
func (m *ServicesJsonRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServicesJsonRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ServicesJsonRequestBuilder when successful
func (m *ServicesJsonRequestBuilder) WithUrl(rawUrl string)(*ServicesJsonRequestBuilder) {
    return NewServicesJsonRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
