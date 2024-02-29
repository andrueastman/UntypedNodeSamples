package metricsjson

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2 "github.com/andrueastman/untypednodesamples/goclient/models"
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
func (m *MetricsJsonRequestBuilder) Get(ctx context.Context, requestConfiguration *MetricsJsonRequestBuilderGetRequestConfiguration)(iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2.Metricsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2.CreateMetricsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iae9fb0262d34e3cd98c53508ace559228b424528fd79025950263229053f76e2.Metricsable), nil
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
