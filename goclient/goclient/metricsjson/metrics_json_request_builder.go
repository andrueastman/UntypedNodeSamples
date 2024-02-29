package metricsjson

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    id6760cbde1c5ad1214c40ac9ca069164ef0f31f8ecb790d183bca1f74a71e629 "github.com/andrueastman/untypednodesamples/goclient/goclient/models"
)

// MetricsJsonRequestBuilder builds and executes requests for operations under \metrics.json
type MetricsJsonRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MetricsJsonRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MetricsJsonRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMetricsJsonRequestBuilderInternal instantiates a new MetricsJsonRequestBuilder and sets the default values.
func NewMetricsJsonRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsJsonRequestBuilder) {
    m := &MetricsJsonRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/metrics.json", pathParameters),
    }
    return m
}
// NewMetricsJsonRequestBuilder instantiates a new MetricsJsonRequestBuilder and sets the default values.
func NewMetricsJsonRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MetricsJsonRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMetricsJsonRequestBuilderInternal(urlParams, requestAdapter)
}
// Get some basic metrics for the entire directory.Just stunning numbers to put on a front page and are intended purely for WoW effect :)
// returns a Metricsable when successful
func (m *MetricsJsonRequestBuilder) Get(ctx context.Context, requestConfiguration *MetricsJsonRequestBuilderGetRequestConfiguration)(id6760cbde1c5ad1214c40ac9ca069164ef0f31f8ecb790d183bca1f74a71e629.Metricsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id6760cbde1c5ad1214c40ac9ca069164ef0f31f8ecb790d183bca1f74a71e629.CreateMetricsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id6760cbde1c5ad1214c40ac9ca069164ef0f31f8ecb790d183bca1f74a71e629.Metricsable), nil
}
// ToGetRequestInformation some basic metrics for the entire directory.Just stunning numbers to put on a front page and are intended purely for WoW effect :)
// returns a *RequestInformation when successful
func (m *MetricsJsonRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MetricsJsonRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MetricsJsonRequestBuilder when successful
func (m *MetricsJsonRequestBuilder) WithUrl(rawUrl string)(*MetricsJsonRequestBuilder) {
    return NewMetricsJsonRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
