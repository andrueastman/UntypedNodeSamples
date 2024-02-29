package specs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithProviderItemRequestBuilder builds and executes requests for operations under \specs\{provider}
type WithProviderItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByApiJsonId gets an item from the github.com/andrueastman/UntypedNodeSamples/GoClient/.specs.item.item collection
// returns a *ItemApiJsonItemRequestBuilder when successful
func (m *WithProviderItemRequestBuilder) ByApiJsonId(apiJsonId string)(*ItemApiJsonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if apiJsonId != "" {
        urlTplParams["apiJson%2Did"] = apiJsonId
    }
    return NewItemApiJsonItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithProviderItemRequestBuilderInternal instantiates a new WithProviderItemRequestBuilder and sets the default values.
func NewWithProviderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithProviderItemRequestBuilder) {
    m := &WithProviderItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/specs/{provider}", pathParameters),
    }
    return m
}
// NewWithProviderItemRequestBuilder instantiates a new WithProviderItemRequestBuilder and sets the default values.
func NewWithProviderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithProviderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithProviderItemRequestBuilderInternal(urlParams, requestAdapter)
}
