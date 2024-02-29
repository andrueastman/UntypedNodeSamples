package specs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2 "github.com/andrueastman/untypednodesamples/goclient/models"
)

// ItemApiJsonItemRequestBuilder builds and executes requests for operations under \specs\{provider}\{apiJson-id}
type ItemApiJsonItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemApiJsonItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemApiJsonItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByApiJson gets an item from the github.com/andrueastman/UntypedNodeSamples/GoClient/.specs.item.item.item collection
// returns a *ItemItemWithApiJsonItemRequestBuilder when successful
func (m *ItemApiJsonItemRequestBuilder) ByApiJson(apiJson string)(*ItemItemWithApiJsonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if apiJson != "" {
        urlTplParams["api"] = apiJson
    }
    return NewItemItemWithApiJsonItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemApiJsonItemRequestBuilderInternal instantiates a new ItemApiJsonItemRequestBuilder and sets the default values.
func NewItemApiJsonItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemApiJsonItemRequestBuilder) {
    m := &ItemApiJsonItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/specs/{provider}/{apiJson%2Did}", pathParameters),
    }
    return m
}
// NewItemApiJsonItemRequestBuilder instantiates a new ItemApiJsonItemRequestBuilder and sets the default values.
func NewItemApiJsonItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemApiJsonItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemApiJsonItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the API entry for one specific version of an API where there is no serviceName.
// returns a APIable when successful
func (m *ItemApiJsonItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemApiJsonItemRequestBuilderGetRequestConfiguration)(iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2.APIable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2.CreateAPIFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2.APIable), nil
}
// ToGetRequestInformation returns the API entry for one specific version of an API where there is no serviceName.
// returns a *RequestInformation when successful
func (m *ItemApiJsonItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemApiJsonItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemApiJsonItemRequestBuilder when successful
func (m *ItemApiJsonItemRequestBuilder) WithUrl(rawUrl string)(*ItemApiJsonItemRequestBuilder) {
    return NewItemApiJsonItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
