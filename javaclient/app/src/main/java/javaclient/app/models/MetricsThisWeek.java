package javaclient.client.models;

import com.microsoft.kiota.serialization.AdditionalDataHolder;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
/**
 * Summary totals for the last 7 days
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public class MetricsThisWeek implements AdditionalDataHolder, Parsable {
    /**
     * APIs added in the last week
     */
    private Integer added;
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    private Map<String, Object> additionalData;
    /**
     * APIs updated in the last week
     */
    private Integer updated;
    /**
     * Instantiates a new {@link MetricsThisWeek} and sets the default values.
     */
    public MetricsThisWeek() {
        this.setAdditionalData(new HashMap<>());
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a {@link MetricsThisWeek}
     */
    @jakarta.annotation.Nonnull
    public static MetricsThisWeek createFromDiscriminatorValue(@jakarta.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new MetricsThisWeek();
    }
    /**
     * Gets the added property value. APIs added in the last week
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getAdded() {
        return this.added;
    }
    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a {@link Map<String, Object>}
     */
    @jakarta.annotation.Nonnull
    public Map<String, Object> getAdditionalData() {
        return this.additionalData;
    }
    /**
     * The deserialization information for the current model
     * @return a {@link Map<String, java.util.function.Consumer<ParseNode>>}
     */
    @jakarta.annotation.Nonnull
    public Map<String, java.util.function.Consumer<ParseNode>> getFieldDeserializers() {
        final HashMap<String, java.util.function.Consumer<ParseNode>> deserializerMap = new HashMap<String, java.util.function.Consumer<ParseNode>>(2);
        deserializerMap.put("added", (n) -> { this.setAdded(n.getIntegerValue()); });
        deserializerMap.put("updated", (n) -> { this.setUpdated(n.getIntegerValue()); });
        return deserializerMap;
    }
    /**
     * Gets the updated property value. APIs updated in the last week
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getUpdated() {
        return this.updated;
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public void serialize(@jakarta.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeIntegerValue("added", this.getAdded());
        writer.writeIntegerValue("updated", this.getUpdated());
        writer.writeAdditionalData(this.getAdditionalData());
    }
    /**
     * Sets the added property value. APIs added in the last week
     * @param value Value to set for the added property.
     */
    public void setAdded(@jakarta.annotation.Nullable final Integer value) {
        this.added = value;
    }
    /**
     * Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     */
    public void setAdditionalData(@jakarta.annotation.Nullable final Map<String, Object> value) {
        this.additionalData = value;
    }
    /**
     * Sets the updated property value. APIs updated in the last week
     * @param value Value to set for the updated property.
     */
    public void setUpdated(@jakarta.annotation.Nullable final Integer value) {
        this.updated = value;
    }
}
