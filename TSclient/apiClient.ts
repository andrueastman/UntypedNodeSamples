/* tslint:disable */
/* eslint-disable */
// Generated by Microsoft Kiota
import { ProviderItemRequestBuilderNavigationMetadata, ProviderItemRequestBuilderRequestsMetadata, type ProviderItemRequestBuilder } from './item/';
import { ListJsonRequestBuilderRequestsMetadata, type ListJsonRequestBuilder } from './listJson/';
import { MetricsJsonRequestBuilderRequestsMetadata, type MetricsJsonRequestBuilder } from './metricsJson/';
import { ProvidersJsonRequestBuilderRequestsMetadata, type ProvidersJsonRequestBuilder } from './providersJson/';
import { SpecsRequestBuilderNavigationMetadata, type SpecsRequestBuilder } from './specs/';
import { apiClientProxifier, registerDefaultDeserializer, registerDefaultSerializer, type BaseRequestBuilder, type KeysToExcludeForNavigationMetadata, type NavigationMetadata, type RequestAdapter } from '@microsoft/kiota-abstractions';
import { FormParseNodeFactory, FormSerializationWriterFactory } from '@microsoft/kiota-serialization-form';
import { JsonParseNodeFactory, JsonSerializationWriterFactory } from '@microsoft/kiota-serialization-json';
import { MultipartSerializationWriterFactory } from '@microsoft/kiota-serialization-multipart';
import { TextParseNodeFactory, TextSerializationWriterFactory } from '@microsoft/kiota-serialization-text';

/**
 * The main entry point of the SDK, exposes the configuration and the fluent API.
 */
export interface ApiClient extends BaseRequestBuilder<ApiClient> {
    /**
     * The listJson property
     */
    get listJson(): ListJsonRequestBuilder;
    /**
     * The metricsJson property
     */
    get metricsJson(): MetricsJsonRequestBuilder;
    /**
     * The providersJson property
     */
    get providersJson(): ProvidersJsonRequestBuilder;
    /**
     * The specs property
     */
    get specs(): SpecsRequestBuilder;
    /**
     * Gets an item from the graphtypescriptv4.utilities.item collection
     * @param providerId Unique identifier of the item
     * @returns {ProviderItemRequestBuilder}
     */
     byProviderId(providerId: string) : ProviderItemRequestBuilder;
}
/**
 * Instantiates a new {@link ApiClient} and sets the default values.
 * @param requestAdapter The request adapter to use to execute the requests.
 */
export function createApiClient(requestAdapter: RequestAdapter) {
    registerDefaultSerializer(JsonSerializationWriterFactory);
    registerDefaultSerializer(TextSerializationWriterFactory);
    registerDefaultSerializer(FormSerializationWriterFactory);
    registerDefaultSerializer(MultipartSerializationWriterFactory);
    registerDefaultDeserializer(JsonParseNodeFactory);
    registerDefaultDeserializer(TextParseNodeFactory);
    registerDefaultDeserializer(FormParseNodeFactory);
    if (requestAdapter.baseUrl === undefined || requestAdapter.baseUrl === "") {
        requestAdapter.baseUrl = "https://api.apis.guru/v2";
    }
    const pathParameters: Record<string, unknown> = {
        "baseurl": requestAdapter.baseUrl,
    };
    return apiClientProxifier<ApiClient>(requestAdapter, pathParameters, ApiClientNavigationMetadata, undefined);
}
/**
 * Uri template for the request builder.
 */
export const ApiClientUriTemplate = "{+baseurl}";
/**
 * Metadata for all the navigation properties in the request builder.
 */
export const ApiClientNavigationMetadata: Record<Exclude<keyof ApiClient, KeysToExcludeForNavigationMetadata>, NavigationMetadata> = {
    byProviderId: {
        requestsMetadata: ProviderItemRequestBuilderRequestsMetadata,
        navigationMetadata: ProviderItemRequestBuilderNavigationMetadata,
        pathParametersMappings: ["provider%2Did"],
    },
    listJson: {
        requestsMetadata: ListJsonRequestBuilderRequestsMetadata,
    },
    metricsJson: {
        requestsMetadata: MetricsJsonRequestBuilderRequestsMetadata,
    },
    providersJson: {
        requestsMetadata: ProvidersJsonRequestBuilderRequestsMetadata,
    },
    specs: {
        navigationMetadata: SpecsRequestBuilderNavigationMetadata,
    },
};
/* tslint:enable */
/* eslint-enable */