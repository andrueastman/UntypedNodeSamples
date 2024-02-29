package javaclient.client.models;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
/**
 * Meta information about API
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public class API implements Parsable {
    /**
     * Timestamp when the API was first added to the directory
     */
    private OffsetDateTime added;
    /**
     * Recommended version
     */
    private String preferred;
    /**
     * List of supported versions of the API
     */
    private APIVersions versions;
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a {@link API}
     */
    @jakarta.annotation.Nonnull
    public static API createFromDiscriminatorValue(@jakarta.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new API();
    }
    /**
     * Gets the added property value. Timestamp when the API was first added to the directory
     * @return a {@link OffsetDateTime}
     */
    @jakarta.annotation.Nullable
    public OffsetDateTime getAdded() {
        return this.added;
    }
    /**
     * The deserialization information for the current model
     * @return a {@link Map<String, java.util.function.Consumer<ParseNode>>}
     */
    @jakarta.annotation.Nonnull
    public Map<String, java.util.function.Consumer<ParseNode>> getFieldDeserializers() {
        final HashMap<String, java.util.function.Consumer<ParseNode>> deserializerMap = new HashMap<String, java.util.function.Consumer<ParseNode>>(3);
        deserializerMap.put("added", (n) -> { this.setAdded(n.getOffsetDateTimeValue()); });
        deserializerMap.put("preferred", (n) -> { this.setPreferred(n.getStringValue()); });
        deserializerMap.put("versions", (n) -> { this.setVersions(n.getObjectValue(APIVersions::createFromDiscriminatorValue)); });
        return deserializerMap;
    }
    /**
     * Gets the preferred property value. Recommended version
     * @return a {@link String}
     */
    @jakarta.annotation.Nullable
    public String getPreferred() {
        return this.preferred;
    }
    /**
     * Gets the versions property value. List of supported versions of the API
     * @return a {@link APIVersions}
     */
    @jakarta.annotation.Nullable
    public APIVersions getVersions() {
        return this.versions;
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public void serialize(@jakarta.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeOffsetDateTimeValue("added", this.getAdded());
        writer.writeStringValue("preferred", this.getPreferred());
        writer.writeObjectValue("versions", this.getVersions());
    }
    /**
     * Sets the added property value. Timestamp when the API was first added to the directory
     * @param value Value to set for the added property.
     */
    public void setAdded(@jakarta.annotation.Nullable final OffsetDateTime value) {
        this.added = value;
    }
    /**
     * Sets the preferred property value. Recommended version
     * @param value Value to set for the preferred property.
     */
    public void setPreferred(@jakarta.annotation.Nullable final String value) {
        this.preferred = value;
    }
    /**
     * Sets the versions property value. List of supported versions of the API
     * @param value Value to set for the versions property.
     */
    public void setVersions(@jakarta.annotation.Nullable final APIVersions value) {
        this.versions = value;
    }
}
