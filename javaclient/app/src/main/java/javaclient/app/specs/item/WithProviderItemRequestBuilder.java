package javaclient.client.specs.item;

import com.microsoft.kiota.BaseRequestBuilder;
import com.microsoft.kiota.RequestAdapter;
import java.util.HashMap;
import java.util.Objects;
import javaclient.client.specs.item.item.ApiJsonItemRequestBuilder;
/**
 * Builds and executes requests for operations under /specs/{provider}
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public class WithProviderItemRequestBuilder extends BaseRequestBuilder {
    /**
     * Gets an item from the javaclient.client.specs.item.item collection
     * @param apiJsonId Unique identifier of the item
     * @return a {@link ApiJsonItemRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public ApiJsonItemRequestBuilder byApiJsonId(@jakarta.annotation.Nonnull final String apiJsonId) {
        Objects.requireNonNull(apiJsonId);
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("apiJson%2Did", apiJsonId);
        return new ApiJsonItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * Instantiates a new {@link WithProviderItemRequestBuilder} and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public WithProviderItemRequestBuilder(@jakarta.annotation.Nonnull final HashMap<String, Object> pathParameters, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/specs/{provider}", pathParameters);
    }
    /**
     * Instantiates a new {@link WithProviderItemRequestBuilder} and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public WithProviderItemRequestBuilder(@jakarta.annotation.Nonnull final String rawUrl, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/specs/{provider}", rawUrl);
    }
}
