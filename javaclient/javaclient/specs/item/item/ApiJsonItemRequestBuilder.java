package javaclient.specs.item.item;

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
import javaclient.models.API;
import javaclient.specs.item.item.item.WithApiJsonItemRequestBuilder;
/**
 * Builds and executes requests for operations under /specs/{provider}/{apiJson-id}
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public class ApiJsonItemRequestBuilder extends BaseRequestBuilder {
    /**
     * Gets an item from the javaclient.specs.item.item.item collection
     * @param apiJson Unique identifier of the item
     * @return a {@link WithApiJsonItemRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public WithApiJsonItemRequestBuilder byApiJson(@jakarta.annotation.Nonnull final String apiJson) {
        Objects.requireNonNull(apiJson);
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("api", apiJson);
        return new WithApiJsonItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * Instantiates a new {@link ApiJsonItemRequestBuilder} and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public ApiJsonItemRequestBuilder(@jakarta.annotation.Nonnull final HashMap<String, Object> pathParameters, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/specs/{provider}/{apiJson%2Did}", pathParameters);
    }
    /**
     * Instantiates a new {@link ApiJsonItemRequestBuilder} and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public ApiJsonItemRequestBuilder(@jakarta.annotation.Nonnull final String rawUrl, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/specs/{provider}/{apiJson%2Did}", rawUrl);
    }
    /**
     * Returns the API entry for one specific version of an API where there is no serviceName.
     * @return a {@link API}
     */
    @jakarta.annotation.Nullable
    public API get() {
        return get(null);
    }
    /**
     * Returns the API entry for one specific version of an API where there is no serviceName.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a {@link API}
     */
    @jakarta.annotation.Nullable
    public API get(@jakarta.annotation.Nullable final java.util.function.Consumer<GetRequestConfiguration> requestConfiguration) {
        final RequestInformation requestInfo = toGetRequestInformation(requestConfiguration);
        return this.requestAdapter.send(requestInfo, null, API::createFromDiscriminatorValue);
    }
    /**
     * Returns the API entry for one specific version of an API where there is no serviceName.
     * @return a {@link RequestInformation}
     */
    @jakarta.annotation.Nonnull
    public RequestInformation toGetRequestInformation() {
        return toGetRequestInformation(null);
    }
    /**
     * Returns the API entry for one specific version of an API where there is no serviceName.
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
     * @return a {@link ApiJsonItemRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public ApiJsonItemRequestBuilder withUrl(@jakarta.annotation.Nonnull final String rawUrl) {
        Objects.requireNonNull(rawUrl);
        return new ApiJsonItemRequestBuilder(rawUrl, requestAdapter);
    }
    /**
     * Configuration for the request such as headers, query parameters, and middleware options.
     */
    @jakarta.annotation.Generated("com.microsoft.kiota")
    public class GetRequestConfiguration extends BaseRequestConfiguration {
    }
}
