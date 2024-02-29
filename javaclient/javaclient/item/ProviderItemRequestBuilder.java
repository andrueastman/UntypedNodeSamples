package javaclient.item;

import com.microsoft.kiota.BaseRequestBuilder;
import com.microsoft.kiota.BaseRequestConfiguration;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.RequestAdapter;
import com.microsoft.kiota.RequestInformation;
import com.microsoft.kiota.RequestOption;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParsableFactory;
import java.util.Collection;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import javaclient.item.servicesjson.ServicesJsonRequestBuilder;
import javaclient.models.APIs;
/**
 * Builds and executes requests for operations under /{provider-id}
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public class ProviderItemRequestBuilder extends BaseRequestBuilder {
    /**
     * The servicesJson property
     * @return a {@link ServicesJsonRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public ServicesJsonRequestBuilder servicesJson() {
        return new ServicesJsonRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * Instantiates a new {@link ProviderItemRequestBuilder} and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public ProviderItemRequestBuilder(@jakarta.annotation.Nonnull final HashMap<String, Object> pathParameters, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/{provider%2Did}", pathParameters);
    }
    /**
     * Instantiates a new {@link ProviderItemRequestBuilder} and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public ProviderItemRequestBuilder(@jakarta.annotation.Nonnull final String rawUrl, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/{provider%2Did}", rawUrl);
    }
    /**
     * List all APIs in the directory for a particular providerNameReturns links to the individual API entry for each API.
     * @return a {@link APIs}
     */
    @jakarta.annotation.Nullable
    public APIs get() {
        return get(null);
    }
    /**
     * List all APIs in the directory for a particular providerNameReturns links to the individual API entry for each API.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a {@link APIs}
     */
    @jakarta.annotation.Nullable
    public APIs get(@jakarta.annotation.Nullable final java.util.function.Consumer<GetRequestConfiguration> requestConfiguration) {
        final RequestInformation requestInfo = toGetRequestInformation(requestConfiguration);
        return this.requestAdapter.send(requestInfo, null, APIs::createFromDiscriminatorValue);
    }
    /**
     * List all APIs in the directory for a particular providerNameReturns links to the individual API entry for each API.
     * @return a {@link RequestInformation}
     */
    @jakarta.annotation.Nonnull
    public RequestInformation toGetRequestInformation() {
        return toGetRequestInformation(null);
    }
    /**
     * List all APIs in the directory for a particular providerNameReturns links to the individual API entry for each API.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a {@link RequestInformation}
     */
    @jakarta.annotation.Nonnull
    public RequestInformation toGetRequestInformation(@jakarta.annotation.Nullable final java.util.function.Consumer<GetRequestConfiguration> requestConfiguration) {
        final RequestInformation requestInfo = new RequestInformation(HttpMethod.GET, urlTemplate, pathParameters);
        requestInfo.configure(requestConfiguration, GetRequestConfiguration::new);
        requestInfo.headers.tryAdd("Accept", "application/json");
        return requestInfo;
    }
    /**
     * Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
     * @param rawUrl The raw URL to use for the request builder.
     * @return a {@link ProviderItemRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public ProviderItemRequestBuilder withUrl(@jakarta.annotation.Nonnull final String rawUrl) {
        Objects.requireNonNull(rawUrl);
        return new ProviderItemRequestBuilder(rawUrl, requestAdapter);
    }
    /**
     * Configuration for the request such as headers, query parameters, and middleware options.
     */
    @jakarta.annotation.Generated("com.microsoft.kiota")
    public class GetRequestConfiguration extends BaseRequestConfiguration {
    }
}
