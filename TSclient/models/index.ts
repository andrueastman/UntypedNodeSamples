/* tslint:disable */
/* eslint-disable */
// Generated by Microsoft Kiota
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

/**
 * Meta information about API
 */
export interface API extends Parsable {
    /**
     * Timestamp when the API was first added to the directory
     */
    added?: Date;
    /**
     * Recommended version
     */
    preferred?: string;
    /**
     * List of supported versions of the API
     */
    versions?: API_versions;
}
/**
 * List of supported versions of the API
 */
export interface API_versions extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
}
/**
 * List of API details.It is a JSON object with API IDs(`<provider>[:<service>]`) as keys.
 */
export interface APIs extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
}
/**
 * Creates a new instance of the appropriate class based on discriminator value
 * @param parseNode The parse node to use to read the discriminator value and create the object
 * @returns {API_versions}
 */
export function createAPI_versionsFromDiscriminatorValue(parseNode: ParseNode | undefined) : ((instance?: Parsable) => Record<string, (node: ParseNode) => void>) {
    return deserializeIntoAPI_versions;
}
/**
 * Creates a new instance of the appropriate class based on discriminator value
 * @param parseNode The parse node to use to read the discriminator value and create the object
 * @returns {API}
 */
export function createAPIFromDiscriminatorValue(parseNode: ParseNode | undefined) : ((instance?: Parsable) => Record<string, (node: ParseNode) => void>) {
    return deserializeIntoAPI;
}
/**
 * Creates a new instance of the appropriate class based on discriminator value
 * @param parseNode The parse node to use to read the discriminator value and create the object
 * @returns {APIs}
 */
export function createAPIsFromDiscriminatorValue(parseNode: ParseNode | undefined) : ((instance?: Parsable) => Record<string, (node: ParseNode) => void>) {
    return deserializeIntoAPIs;
}
/**
 * Creates a new instance of the appropriate class based on discriminator value
 * @param parseNode The parse node to use to read the discriminator value and create the object
 * @returns {Metrics_thisWeek}
 */
export function createMetrics_thisWeekFromDiscriminatorValue(parseNode: ParseNode | undefined) : ((instance?: Parsable) => Record<string, (node: ParseNode) => void>) {
    return deserializeIntoMetrics_thisWeek;
}
/**
 * Creates a new instance of the appropriate class based on discriminator value
 * @param parseNode The parse node to use to read the discriminator value and create the object
 * @returns {Metrics}
 */
export function createMetricsFromDiscriminatorValue(parseNode: ParseNode | undefined) : ((instance?: Parsable) => Record<string, (node: ParseNode) => void>) {
    return deserializeIntoMetrics;
}
/**
 * The deserialization information for the current model
 * @returns {Record<string, (node: ParseNode) => void>}
 */
export function deserializeIntoAPI(aPI: Partial<API> | undefined = {}) : Record<string, (node: ParseNode) => void> {
    return {
        "added": n => { aPI.added = n.getDateValue(); },
        "preferred": n => { aPI.preferred = n.getStringValue(); },
        "versions": n => { aPI.versions = n.getObjectValue<API_versions>(createAPI_versionsFromDiscriminatorValue); },
    }
}
/**
 * The deserialization information for the current model
 * @returns {Record<string, (node: ParseNode) => void>}
 */
export function deserializeIntoAPI_versions(aPI_versions: Partial<API_versions> | undefined = {}) : Record<string, (node: ParseNode) => void> {
    return {
    }
}
/**
 * The deserialization information for the current model
 * @returns {Record<string, (node: ParseNode) => void>}
 */
export function deserializeIntoAPIs(aPIs: Partial<APIs> | undefined = {}) : Record<string, (node: ParseNode) => void> {
    return {
    }
}
/**
 * The deserialization information for the current model
 * @returns {Record<string, (node: ParseNode) => void>}
 */
export function deserializeIntoMetrics(metrics: Partial<Metrics> | undefined = {}) : Record<string, (node: ParseNode) => void> {
    return {
        "datasets": n => { metrics.datasets = n.getObjectValue<UntypedNode>(createUntypedNodeFromDiscriminatorValue); },
        "fixedPct": n => { metrics.fixedPct = n.getNumberValue(); },
        "fixes": n => { metrics.fixes = n.getNumberValue(); },
        "invalid": n => { metrics.invalid = n.getNumberValue(); },
        "issues": n => { metrics.issues = n.getNumberValue(); },
        "numAPIs": n => { metrics.numAPIs = n.getNumberValue(); },
        "numDrivers": n => { metrics.numDrivers = n.getNumberValue(); },
        "numEndpoints": n => { metrics.numEndpoints = n.getNumberValue(); },
        "numProviders": n => { metrics.numProviders = n.getNumberValue(); },
        "numSpecs": n => { metrics.numSpecs = n.getNumberValue(); },
        "stars": n => { metrics.stars = n.getNumberValue(); },
        "thisWeek": n => { metrics.thisWeek = n.getObjectValue<Metrics_thisWeek>(createMetrics_thisWeekFromDiscriminatorValue); },
        "unofficial": n => { metrics.unofficial = n.getNumberValue(); },
        "unreachable": n => { metrics.unreachable = n.getNumberValue(); },
    }
}
/**
 * The deserialization information for the current model
 * @returns {Record<string, (node: ParseNode) => void>}
 */
export function deserializeIntoMetrics_thisWeek(metrics_thisWeek: Partial<Metrics_thisWeek> | undefined = {}) : Record<string, (node: ParseNode) => void> {
    return {
        "added": n => { metrics_thisWeek.added = n.getNumberValue(); },
        "updated": n => { metrics_thisWeek.updated = n.getNumberValue(); },
    }
}
/**
 * List of basic metrics
 */
export interface Metrics extends Parsable {
    /**
     * Data used for charting etc
     */
    datasets?: UntypedNode;
    /**
     * Percentage of all APIs where auto fixes have been applied
     */
    fixedPct?: number;
    /**
     * Total number of fixes applied across all APIs
     */
    fixes?: number;
    /**
     * Number of newly invalid APIs
     */
    invalid?: number;
    /**
     * Open GitHub issues on our main repo
     */
    issues?: number;
    /**
     * Number of unique APIs
     */
    numAPIs?: number;
    /**
     * Number of methods of API retrieval
     */
    numDrivers?: number;
    /**
     * Total number of endpoints inside all definitions
     */
    numEndpoints?: number;
    /**
     * Number of API providers in directory
     */
    numProviders?: number;
    /**
     * Number of API definitions including different versions of the same API
     */
    numSpecs?: number;
    /**
     * GitHub stars for our main repo
     */
    stars?: number;
    /**
     * Summary totals for the last 7 days
     */
    thisWeek?: Metrics_thisWeek;
    /**
     * Number of unofficial APIs
     */
    unofficial?: number;
    /**
     * Number of unreachable (4XX,5XX status) APIs
     */
    unreachable?: number;
}
/**
 * Summary totals for the last 7 days
 */
export interface Metrics_thisWeek extends AdditionalDataHolder, Parsable {
    /**
     * APIs added in the last week
     */
    added?: number;
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * APIs updated in the last week
     */
    updated?: number;
}
/**
 * Serializes information the current object
 * @param writer Serialization writer to use to serialize this model
 */
export function serializeAPI(writer: SerializationWriter, aPI: Partial<API> | undefined = {}) : void {
    writer.writeDateValue("added", aPI.added);
    writer.writeStringValue("preferred", aPI.preferred);
    writer.writeObjectValue<API_versions>("versions", aPI.versions, serializeAPI_versions);
}
/**
 * Serializes information the current object
 * @param writer Serialization writer to use to serialize this model
 */
export function serializeAPI_versions(writer: SerializationWriter, aPI_versions: Partial<API_versions> | undefined = {}) : void {
    writer.writeAdditionalData(aPI_versions.additionalData);
}
/**
 * Serializes information the current object
 * @param writer Serialization writer to use to serialize this model
 */
export function serializeAPIs(writer: SerializationWriter, aPIs: Partial<APIs> | undefined = {}) : void {
    writer.writeAdditionalData(aPIs.additionalData);
}
/**
 * Serializes information the current object
 * @param writer Serialization writer to use to serialize this model
 */
export function serializeMetrics(writer: SerializationWriter, metrics: Partial<Metrics> | undefined = {}) : void {
    writer.writeObjectValue("datasets", metrics.datasets);
    writer.writeNumberValue("fixedPct", metrics.fixedPct);
    writer.writeNumberValue("fixes", metrics.fixes);
    writer.writeNumberValue("invalid", metrics.invalid);
    writer.writeNumberValue("issues", metrics.issues);
    writer.writeNumberValue("numAPIs", metrics.numAPIs);
    writer.writeNumberValue("numDrivers", metrics.numDrivers);
    writer.writeNumberValue("numEndpoints", metrics.numEndpoints);
    writer.writeNumberValue("numProviders", metrics.numProviders);
    writer.writeNumberValue("numSpecs", metrics.numSpecs);
    writer.writeNumberValue("stars", metrics.stars);
    writer.writeObjectValue<Metrics_thisWeek>("thisWeek", metrics.thisWeek, serializeMetrics_thisWeek);
    writer.writeNumberValue("unofficial", metrics.unofficial);
    writer.writeNumberValue("unreachable", metrics.unreachable);
}
/**
 * Serializes information the current object
 * @param writer Serialization writer to use to serialize this model
 */
export function serializeMetrics_thisWeek(writer: SerializationWriter, metrics_thisWeek: Partial<Metrics_thisWeek> | undefined = {}) : void {
    writer.writeNumberValue("added", metrics_thisWeek.added);
    writer.writeNumberValue("updated", metrics_thisWeek.updated);
    writer.writeAdditionalData(metrics_thisWeek.additionalData);
}
/* tslint:enable */
/* eslint-enable */
