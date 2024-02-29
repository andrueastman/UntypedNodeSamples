package javaclient;

import com.microsoft.kiota.ApiClientBuilder;
import com.microsoft.kiota.BaseRequestBuilder;
import com.microsoft.kiota.RequestAdapter;
import com.microsoft.kiota.serialization.FormParseNodeFactory;
import com.microsoft.kiota.serialization.FormSerializationWriterFactory;
import com.microsoft.kiota.serialization.JsonParseNodeFactory;
import com.microsoft.kiota.serialization.JsonSerializationWriterFactory;
import com.microsoft.kiota.serialization.MultipartSerializationWriterFactory;
import com.microsoft.kiota.serialization.ParseNodeFactoryRegistry;
import com.microsoft.kiota.serialization.SerializationWriterFactoryRegistry;
import com.microsoft.kiota.serialization.TextParseNodeFactory;
import com.microsoft.kiota.serialization.TextSerializationWriterFactory;
import java.util.HashMap;
import java.util.Objects;
import javaclient.item.ProviderItemRequestBuilder;
import javaclient.listjson.ListJsonRequestBuilder;
import javaclient.metricsjson.MetricsJsonRequestBuilder;
import javaclient.providersjson.ProvidersJsonRequestBuilder;
import javaclient.specs.SpecsRequestBuilder;
/**
 * The main entry point of the SDK, exposes the configuration and the fluent API.
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public class ApiClient extends BaseRequestBuilder {
    /**
     * The listJson property
     * @return a {@link ListJsonRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public ListJsonRequestBuilder listJson() {
        return new ListJsonRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The metricsJson property
     * @return a {@link MetricsJsonRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public MetricsJsonRequestBuilder metricsJson() {
        return new MetricsJsonRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The providersJson property
     * @return a {@link ProvidersJsonRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public ProvidersJsonRequestBuilder providersJson() {
        return new ProvidersJsonRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The specs property
     * @return a {@link SpecsRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public SpecsRequestBuilder specs() {
        return new SpecsRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * Gets an item from the javaclient.item collection
     * @param providerId Unique identifier of the item
     * @return a {@link ProviderItemRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public ProviderItemRequestBuilder byProviderId(@jakarta.annotation.Nonnull final String providerId) {
        Objects.requireNonNull(providerId);
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("provider%2Did", providerId);
        return new ProviderItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * Instantiates a new {@link ApiClient} and sets the default values.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public ApiClient(@jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}");
        this.pathParameters = new HashMap<>();
        ApiClientBuilder.registerDefaultSerializer(() -> new JsonSerializationWriterFactory());
        ApiClientBuilder.registerDefaultSerializer(() -> new TextSerializationWriterFactory());
        ApiClientBuilder.registerDefaultSerializer(() -> new FormSerializationWriterFactory());
        ApiClientBuilder.registerDefaultSerializer(() -> new MultipartSerializationWriterFactory());
        ApiClientBuilder.registerDefaultDeserializer(() -> new JsonParseNodeFactory());
        ApiClientBuilder.registerDefaultDeserializer(() -> new FormParseNodeFactory());
        ApiClientBuilder.registerDefaultDeserializer(() -> new TextParseNodeFactory());
        if (requestAdapter.getBaseUrl() == null || requestAdapter.getBaseUrl().isEmpty()) {
            requestAdapter.setBaseUrl("https://api.apis.guru/v2");
        }
        pathParameters.put("baseurl", requestAdapter.getBaseUrl());
    }
}
