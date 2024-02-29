package goclient

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i1c50d8d42f015f5d5899432b2871abef69bc5ea327f21aa9d3ed6871706258bb "github.com/andrueastman/untypednodesamples/goclient/metricsjson"
    i26313f826591dd405f2fa3870f7c7a3e7f1e2f43bf087379334ec7a39b15544d "github.com/andrueastman/untypednodesamples/goclient/specs"
    i36ca6f1d012e0be89ec67d85485365301d9ae942df8a28b1befc182ba95e1d78 "github.com/andrueastman/untypednodesamples/goclient/listjson"
    iecd88ee973ec927f64463b934d78069f4c05f80dd08eec6c0eccd31e0bb8170a "github.com/andrueastman/untypednodesamples/goclient/item"
    if1c3f206bb8fcff4e3799bf55187acb223f8337e0e5f9eb3350041b509cd8965 "github.com/andrueastman/untypednodesamples/goclient/providersjson"
)

// ApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type ApiClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByProviderId gets an item from the github.com/andrueastman/UntypedNodeSamples/GoClient/.item collection
// returns a *ProviderItemRequestBuilder when successful
func (m *ApiClient) ByProviderId(providerId string)(*iecd88ee973ec927f64463b934d78069f4c05f80dd08eec6c0eccd31e0bb8170a.ProviderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if providerId != "" {
        urlTplParams["provider%2Did"] = providerId
    }
    return iecd88ee973ec927f64463b934d78069f4c05f80dd08eec6c0eccd31e0bb8170a.NewProviderItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewApiClient instantiates a new ApiClient and sets the default values.
func NewApiClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiClient) {
    m := &ApiClient{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("https://api.apis.guru/v2")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
// ListJson the listJson property
// returns a *ListJsonRequestBuilder when successful
func (m *ApiClient) ListJson()(*i36ca6f1d012e0be89ec67d85485365301d9ae942df8a28b1befc182ba95e1d78.ListJsonRequestBuilder) {
    return i36ca6f1d012e0be89ec67d85485365301d9ae942df8a28b1befc182ba95e1d78.NewListJsonRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MetricsJson the metricsJson property
// returns a *MetricsJsonRequestBuilder when successful
func (m *ApiClient) MetricsJson()(*i1c50d8d42f015f5d5899432b2871abef69bc5ea327f21aa9d3ed6871706258bb.MetricsJsonRequestBuilder) {
    return i1c50d8d42f015f5d5899432b2871abef69bc5ea327f21aa9d3ed6871706258bb.NewMetricsJsonRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ProvidersJson the providersJson property
// returns a *ProvidersJsonRequestBuilder when successful
func (m *ApiClient) ProvidersJson()(*if1c3f206bb8fcff4e3799bf55187acb223f8337e0e5f9eb3350041b509cd8965.ProvidersJsonRequestBuilder) {
    return if1c3f206bb8fcff4e3799bf55187acb223f8337e0e5f9eb3350041b509cd8965.NewProvidersJsonRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Specs the specs property
// returns a *SpecsRequestBuilder when successful
func (m *ApiClient) Specs()(*i26313f826591dd405f2fa3870f7c7a3e7f1e2f43bf087379334ec7a39b15544d.SpecsRequestBuilder) {
    return i26313f826591dd405f2fa3870f7c7a3e7f1e2f43bf087379334ec7a39b15544d.NewSpecsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
