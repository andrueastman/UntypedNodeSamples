package specs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SpecsRequestBuilder builds and executes requests for operations under \specs
type SpecsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByProvider gets an item from the github.com/andrueastman/UntypedNodeSamples/GoClient/GoClient.specs.item collection
// returns a *WithProviderItemRequestBuilder when successful
func (m *SpecsRequestBuilder) ByProvider(provider string)(*WithProviderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if provider != "" {
        urlTplParams["provider"] = provider
    }
    return NewWithProviderItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewSpecsRequestBuilderInternal instantiates a new SpecsRequestBuilder and sets the default values.
func NewSpecsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SpecsRequestBuilder) {
    m := &SpecsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/specs", pathParameters),
    }
    return m
}
// NewSpecsRequestBuilder instantiates a new SpecsRequestBuilder and sets the default values.
func NewSpecsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SpecsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSpecsRequestBuilderInternal(urlParams, requestAdapter)
}
