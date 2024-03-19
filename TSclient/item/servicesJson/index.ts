/* tslint:disable */
/* eslint-disable */
// Generated by Microsoft Kiota
import { type AdditionalDataHolder, type BaseRequestBuilder, type Parsable, type ParsableFactory, type ParseNode, type RequestConfiguration, type RequestInformation, type RequestsMetadata, type SerializationWriter } from '@microsoft/kiota-abstractions';

/**
 * Creates a new instance of the appropriate class based on discriminator value
 * @param parseNode The parse node to use to read the discriminator value and create the object
 * @returns {ServicesGetResponse}
 */
export function createServicesGetResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : ((instance?: Parsable) => Record<string, (node: ParseNode) => void>) {
    return deserializeIntoServicesGetResponse;
}
/**
 * The deserialization information for the current model
 * @returns {Record<string, (node: ParseNode) => void>}
 */
export function deserializeIntoServicesGetResponse(servicesGetResponse: Partial<ServicesGetResponse> | undefined = {}) : Record<string, (node: ParseNode) => void> {
    return {
        "data": n => { servicesGetResponse.data = n.getCollectionOfPrimitiveValues<string>(); },
    }
}
/**
 * Serializes information the current object
 * @param writer Serialization writer to use to serialize this model
 */
export function serializeServicesGetResponse(writer: SerializationWriter, servicesGetResponse: Partial<ServicesGetResponse> | undefined = {}) : void {
    writer.writeCollectionOfPrimitiveValues<string>("data", servicesGetResponse.data);
    writer.writeAdditionalData(servicesGetResponse.additionalData);
}
export interface ServicesGetResponse extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * The data property
     */
    data?: string[];
}
/**
 * Builds and executes requests for operations under /{provider-id}/services.json
 */
export interface ServicesJsonRequestBuilder extends BaseRequestBuilder<ServicesJsonRequestBuilder> {
    /**
     * List all serviceNames in the directory for a particular providerName
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {Promise<ServicesGetResponse>}
     */
     get(requestConfiguration?: RequestConfiguration<object> | undefined) : Promise<ServicesGetResponse | undefined>;
    /**
     * List all serviceNames in the directory for a particular providerName
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {RequestInformation}
     */
     toGetRequestInformation(requestConfiguration?: RequestConfiguration<object> | undefined) : RequestInformation;
}
/**
 * Uri template for the request builder.
 */
export const ServicesJsonRequestBuilderUriTemplate = "{+baseurl}/{provider%2Did}/services.json";
/**
 * Metadata for all the requests in the request builder.
 */
export const ServicesJsonRequestBuilderRequestsMetadata: RequestsMetadata = {
    get: {
        uriTemplate: ServicesJsonRequestBuilderUriTemplate,
        responseBodyContentType: "application/json",
        adapterMethodName: "send",
        responseBodyFactory:  createServicesGetResponseFromDiscriminatorValue,
    },
};
/* tslint:enable */
/* eslint-enable */