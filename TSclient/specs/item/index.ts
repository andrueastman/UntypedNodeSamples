/* tslint:disable */
/* eslint-disable */
// Generated by Microsoft Kiota
import { ApiJsonItemRequestBuilderNavigationMetadata, ApiJsonItemRequestBuilderRequestsMetadata, type ApiJsonItemRequestBuilder } from './item/';
import { type BaseRequestBuilder, type KeysToExcludeForNavigationMetadata, type NavigationMetadata } from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /specs/{provider}
 */
export interface WithProviderItemRequestBuilder extends BaseRequestBuilder<WithProviderItemRequestBuilder> {
    /**
     * Gets an item from the graphtypescriptv4.utilities.specs.item.item collection
     * @param apiJsonId Unique identifier of the item
     * @returns {ApiJsonItemRequestBuilder}
     */
     byApiJsonId(apiJsonId: string) : ApiJsonItemRequestBuilder;
}
/**
 * Uri template for the request builder.
 */
export const WithProviderItemRequestBuilderUriTemplate = "{+baseurl}/specs/{provider}";
/**
 * Metadata for all the navigation properties in the request builder.
 */
export const WithProviderItemRequestBuilderNavigationMetadata: Record<Exclude<keyof WithProviderItemRequestBuilder, KeysToExcludeForNavigationMetadata>, NavigationMetadata> = {
    byApiJsonId: {
        requestsMetadata: ApiJsonItemRequestBuilderRequestsMetadata,
        navigationMetadata: ApiJsonItemRequestBuilderNavigationMetadata,
        pathParametersMappings: ["apiJson%2Did"],
    },
};
/* tslint:enable */
/* eslint-enable */