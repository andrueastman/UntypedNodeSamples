/* tslint:disable */
/* eslint-disable */
// Generated by Microsoft Kiota
import { type WithProviderItemRequestBuilder, WithProviderItemRequestBuilderNavigationMetadata } from './item/';
import { type BaseRequestBuilder, type KeysToExcludeForNavigationMetadata, type NavigationMetadata } from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /specs
 */
export interface SpecsRequestBuilder extends BaseRequestBuilder<SpecsRequestBuilder> {
    /**
     * Gets an item from the graphtypescriptv4.utilities.specs.item collection
     * @param provider Unique identifier of the item
     * @returns {WithProviderItemRequestBuilder}
     */
     byProvider(provider: string) : WithProviderItemRequestBuilder;
}
/**
 * Uri template for the request builder.
 */
export const SpecsRequestBuilderUriTemplate = "{+baseurl}/specs";
/**
 * Metadata for all the navigation properties in the request builder.
 */
export const SpecsRequestBuilderNavigationMetadata: Record<Exclude<keyof SpecsRequestBuilder, KeysToExcludeForNavigationMetadata>, NavigationMetadata> = {
    byProvider: {
        navigationMetadata: WithProviderItemRequestBuilderNavigationMetadata,
        pathParametersMappings: ["provider"],
    },
};
/* tslint:enable */
/* eslint-enable */