package goclient

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i09660f6760c580472f1bed66035949142e3d83caba3f9f4679e9044a906e4a2e "github.com/andrueastman/untypednodesamples/goclient/goclient/providersjson"
    i72d086d4f672d8ba7e7ba71e68189724b58bfa196e4f996422c55ca70ff86ebb "github.com/andrueastman/untypednodesamples/goclient/goclient/metricsjson"
    i7b02e18c76d7fd596404c542e7640988e55a7a4e8d71499c4c84131930cec3ae "github.com/andrueastman/untypednodesamples/goclient/goclient/item"
    i963d1cd5dbc7fdf91709e41663675ee1c8ae1490a22fb961144463145c017303 "github.com/andrueastman/untypednodesamples/goclient/goclient/listjson"
    iaecb83612fa302773a3a593b916318735d6fb22ae54e9b5107bed3d79cfe53f7 "github.com/andrueastman/untypednodesamples/goclient/goclient/specs"
)

// ApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type ApiClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByProviderId gets an item from the github.com/andrueastman/UntypedNodeSamples/GoClient/GoClient.item collection
// returns a *ProviderItemRequestBuilder when successful
func (m *ApiClient) ByProviderId(providerId string)(*i7b02e18c76d7fd596404c542e7640988e55a7a4e8d71499c4c84131930cec3ae.ProviderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if providerId != "" {
        urlTplParams["provider%2Did"] = providerId
    }
    return i7b02e18c76d7fd596404c542e7640988e55a7a4e8d71499c4c84131930cec3ae.NewProviderItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
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
func (m *ApiClient) ListJson()(*i963d1cd5dbc7fdf91709e41663675ee1c8ae1490a22fb961144463145c017303.ListJsonRequestBuilder) {
    return i963d1cd5dbc7fdf91709e41663675ee1c8ae1490a22fb961144463145c017303.NewListJsonRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MetricsJson the metricsJson property
// returns a *MetricsJsonRequestBuilder when successful
func (m *ApiClient) MetricsJson()(*i72d086d4f672d8ba7e7ba71e68189724b58bfa196e4f996422c55ca70ff86ebb.MetricsJsonRequestBuilder) {
    return i72d086d4f672d8ba7e7ba71e68189724b58bfa196e4f996422c55ca70ff86ebb.NewMetricsJsonRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ProvidersJson the providersJson property
// returns a *ProvidersJsonRequestBuilder when successful
func (m *ApiClient) ProvidersJson()(*i09660f6760c580472f1bed66035949142e3d83caba3f9f4679e9044a906e4a2e.ProvidersJsonRequestBuilder) {
    return i09660f6760c580472f1bed66035949142e3d83caba3f9f4679e9044a906e4a2e.NewProvidersJsonRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Specs the specs property
// returns a *SpecsRequestBuilder when successful
func (m *ApiClient) Specs()(*iaecb83612fa302773a3a593b916318735d6fb22ae54e9b5107bed3d79cfe53f7.SpecsRequestBuilder) {
    return iaecb83612fa302773a3a593b916318735d6fb22ae54e9b5107bed3d79cfe53f7.NewSpecsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
