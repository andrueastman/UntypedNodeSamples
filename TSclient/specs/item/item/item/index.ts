/* tslint:disable */
/* eslint-disable */
// Generated by Microsoft Kiota
import { createAPIFromDiscriminatorValue, type API } from '../../../../models/';
import { type BaseRequestBuilder, type Parsable, type ParsableFactory, type RequestConfiguration, type RequestInformation, type RequestsMetadata } from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /specs/{provider}/{apiJson-id}/{api}.json
 */
export interface WithApiJsonItemRequestBuilder extends BaseRequestBuilder<WithApiJsonItemRequestBuilder> {
    /**
     * Returns the API entry for one specific version of an API where there is a serviceName.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {Promise<API>}
     */
     get(requestConfiguration?: RequestConfiguration<object> | undefined) : Promise<API | undefined>;
    /**
     * Returns the API entry for one specific version of an API where there is a serviceName.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {RequestInformation}
     */
     toGetRequestInformation(requestConfiguration?: RequestConfiguration<object> | undefined) : RequestInformation;
}
/**
 * Uri template for the request builder.
 */
export const WithApiJsonItemRequestBuilderUriTemplate = "{+baseurl}/specs/{provider}/{apiJson%2Did}/{api}.json";
/**
 * Metadata for all the requests in the request builder.
 */
export const WithApiJsonItemRequestBuilderRequestsMetadata: RequestsMetadata = {
    get: {
        uriTemplate: WithApiJsonItemRequestBuilderUriTemplate,
        responseBodyContentType: "application/json",
        adapterMethodName: "send",
        responseBodyFactory:  createAPIFromDiscriminatorValue,
    },
};
/* tslint:enable */
/* eslint-enable */
