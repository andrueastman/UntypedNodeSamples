package specs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2 "github.com/andrueastman/untypednodesamples/goclient/models"
)

// ItemItemWithApiJsonItemRequestBuilder builds and executes requests for operations under \specs\{provider}\{apiJson-id}\{api}.json
type ItemItemWithApiJsonItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemWithApiJsonItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemWithApiJsonItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemWithApiJsonItemRequestBuilderInternal instantiates a new ItemItemWithApiJsonItemRequestBuilder and sets the default values.
func NewItemItemWithApiJsonItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemWithApiJsonItemRequestBuilder) {
    m := &ItemItemWithApiJsonItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/specs/{provider}/{apiJson%2Did}/{api}.json", pathParameters),
    }
    return m
}
// NewItemItemWithApiJsonItemRequestBuilder instantiates a new ItemItemWithApiJsonItemRequestBuilder and sets the default values.
func NewItemItemWithApiJsonItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemWithApiJsonItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemWithApiJsonItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the API entry for one specific version of an API where there is a serviceName.
// returns a APIable when successful
func (m *ItemItemWithApiJsonItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemWithApiJsonItemRequestBuilderGetRequestConfiguration)(iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2.APIable, error) {
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
// ToGetRequestInformation returns the API entry for one specific version of an API where there is a serviceName.
// returns a *RequestInformation when successful
func (m *ItemItemWithApiJsonItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemWithApiJsonItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemWithApiJsonItemRequestBuilder when successful
func (m *ItemItemWithApiJsonItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemWithApiJsonItemRequestBuilder) {
    return NewItemItemWithApiJsonItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
