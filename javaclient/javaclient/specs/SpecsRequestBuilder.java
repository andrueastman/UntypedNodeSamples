package javaclient.specs;

import com.microsoft.kiota.BaseRequestBuilder;
import com.microsoft.kiota.RequestAdapter;
import java.util.HashMap;
import java.util.Objects;
import javaclient.specs.item.WithProviderItemRequestBuilder;
/**
 * Builds and executes requests for operations under /specs
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public class SpecsRequestBuilder extends BaseRequestBuilder {
    /**
     * Gets an item from the javaclient.specs.item collection
     * @param provider Unique identifier of the item
     * @return a {@link WithProviderItemRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public WithProviderItemRequestBuilder byProvider(@jakarta.annotation.Nonnull final String provider) {
        Objects.requireNonNull(provider);
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("provider", provider);
        return new WithProviderItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * Instantiates a new {@link SpecsRequestBuilder} and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public SpecsRequestBuilder(@jakarta.annotation.Nonnull final HashMap<String, Object> pathParameters, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/specs", pathParameters);
    }
    /**
     * Instantiates a new {@link SpecsRequestBuilder} and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public SpecsRequestBuilder(@jakarta.annotation.Nonnull final String rawUrl, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/specs", rawUrl);
    }
}
