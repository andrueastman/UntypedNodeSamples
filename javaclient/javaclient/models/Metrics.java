package javaclient.models;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.UntypedNode;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
/**
 * List of basic metrics
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public class Metrics implements Parsable {
    /**
     * Data used for charting etc
     */
    private UntypedNode datasets;
    /**
     * Percentage of all APIs where auto fixes have been applied
     */
    private Integer fixedPct;
    /**
     * Total number of fixes applied across all APIs
     */
    private Integer fixes;
    /**
     * Number of newly invalid APIs
     */
    private Integer invalid;
    /**
     * Open GitHub issues on our main repo
     */
    private Integer issues;
    /**
     * Number of unique APIs
     */
    private Integer numAPIs;
    /**
     * Number of methods of API retrieval
     */
    private Integer numDrivers;
    /**
     * Total number of endpoints inside all definitions
     */
    private Integer numEndpoints;
    /**
     * Number of API providers in directory
     */
    private Integer numProviders;
    /**
     * Number of API definitions including different versions of the same API
     */
    private Integer numSpecs;
    /**
     * GitHub stars for our main repo
     */
    private Integer stars;
    /**
     * Summary totals for the last 7 days
     */
    private MetricsThisWeek thisWeek;
    /**
     * Number of unofficial APIs
     */
    private Integer unofficial;
    /**
     * Number of unreachable (4XX,5XX status) APIs
     */
    private Integer unreachable;
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a {@link Metrics}
     */
    @jakarta.annotation.Nonnull
    public static Metrics createFromDiscriminatorValue(@jakarta.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new Metrics();
    }
    /**
     * Gets the datasets property value. Data used for charting etc
     * @return a {@link UntypedNode}
     */
    @jakarta.annotation.Nullable
    public UntypedNode getDatasets() {
        return this.datasets;
    }
    /**
     * The deserialization information for the current model
     * @return a {@link Map<String, java.util.function.Consumer<ParseNode>>}
     */
    @jakarta.annotation.Nonnull
    public Map<String, java.util.function.Consumer<ParseNode>> getFieldDeserializers() {
        final HashMap<String, java.util.function.Consumer<ParseNode>> deserializerMap = new HashMap<String, java.util.function.Consumer<ParseNode>>(14);
        deserializerMap.put("datasets", (n) -> { this.setDatasets(n.getObjectValue(UntypedNode::createFromDiscriminatorValue)); });
        deserializerMap.put("fixedPct", (n) -> { this.setFixedPct(n.getIntegerValue()); });
        deserializerMap.put("fixes", (n) -> { this.setFixes(n.getIntegerValue()); });
        deserializerMap.put("invalid", (n) -> { this.setInvalid(n.getIntegerValue()); });
        deserializerMap.put("issues", (n) -> { this.setIssues(n.getIntegerValue()); });
        deserializerMap.put("numAPIs", (n) -> { this.setNumAPIs(n.getIntegerValue()); });
        deserializerMap.put("numDrivers", (n) -> { this.setNumDrivers(n.getIntegerValue()); });
        deserializerMap.put("numEndpoints", (n) -> { this.setNumEndpoints(n.getIntegerValue()); });
        deserializerMap.put("numProviders", (n) -> { this.setNumProviders(n.getIntegerValue()); });
        deserializerMap.put("numSpecs", (n) -> { this.setNumSpecs(n.getIntegerValue()); });
        deserializerMap.put("stars", (n) -> { this.setStars(n.getIntegerValue()); });
        deserializerMap.put("thisWeek", (n) -> { this.setThisWeek(n.getObjectValue(MetricsThisWeek::createFromDiscriminatorValue)); });
        deserializerMap.put("unofficial", (n) -> { this.setUnofficial(n.getIntegerValue()); });
        deserializerMap.put("unreachable", (n) -> { this.setUnreachable(n.getIntegerValue()); });
        return deserializerMap;
    }
    /**
     * Gets the fixedPct property value. Percentage of all APIs where auto fixes have been applied
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getFixedPct() {
        return this.fixedPct;
    }
    /**
     * Gets the fixes property value. Total number of fixes applied across all APIs
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getFixes() {
        return this.fixes;
    }
    /**
     * Gets the invalid property value. Number of newly invalid APIs
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getInvalid() {
        return this.invalid;
    }
    /**
     * Gets the issues property value. Open GitHub issues on our main repo
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getIssues() {
        return this.issues;
    }
    /**
     * Gets the numAPIs property value. Number of unique APIs
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getNumAPIs() {
        return this.numAPIs;
    }
    /**
     * Gets the numDrivers property value. Number of methods of API retrieval
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getNumDrivers() {
        return this.numDrivers;
    }
    /**
     * Gets the numEndpoints property value. Total number of endpoints inside all definitions
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getNumEndpoints() {
        return this.numEndpoints;
    }
    /**
     * Gets the numProviders property value. Number of API providers in directory
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getNumProviders() {
        return this.numProviders;
    }
    /**
     * Gets the numSpecs property value. Number of API definitions including different versions of the same API
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getNumSpecs() {
        return this.numSpecs;
    }
    /**
     * Gets the stars property value. GitHub stars for our main repo
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getStars() {
        return this.stars;
    }
    /**
     * Gets the thisWeek property value. Summary totals for the last 7 days
     * @return a {@link MetricsThisWeek}
     */
    @jakarta.annotation.Nullable
    public MetricsThisWeek getThisWeek() {
        return this.thisWeek;
    }
    /**
     * Gets the unofficial property value. Number of unofficial APIs
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getUnofficial() {
        return this.unofficial;
    }
    /**
     * Gets the unreachable property value. Number of unreachable (4XX,5XX status) APIs
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getUnreachable() {
        return this.unreachable;
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public void serialize(@jakarta.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeObjectValue("datasets", this.getDatasets());
        writer.writeIntegerValue("fixedPct", this.getFixedPct());
        writer.writeIntegerValue("fixes", this.getFixes());
        writer.writeIntegerValue("invalid", this.getInvalid());
        writer.writeIntegerValue("issues", this.getIssues());
        writer.writeIntegerValue("numAPIs", this.getNumAPIs());
        writer.writeIntegerValue("numDrivers", this.getNumDrivers());
        writer.writeIntegerValue("numEndpoints", this.getNumEndpoints());
        writer.writeIntegerValue("numProviders", this.getNumProviders());
        writer.writeIntegerValue("numSpecs", this.getNumSpecs());
        writer.writeIntegerValue("stars", this.getStars());
        writer.writeObjectValue("thisWeek", this.getThisWeek());
        writer.writeIntegerValue("unofficial", this.getUnofficial());
        writer.writeIntegerValue("unreachable", this.getUnreachable());
    }
    /**
     * Sets the datasets property value. Data used for charting etc
     * @param value Value to set for the datasets property.
     */
    public void setDatasets(@jakarta.annotation.Nullable final UntypedNode value) {
        this.datasets = value;
    }
    /**
     * Sets the fixedPct property value. Percentage of all APIs where auto fixes have been applied
     * @param value Value to set for the fixedPct property.
     */
    public void setFixedPct(@jakarta.annotation.Nullable final Integer value) {
        this.fixedPct = value;
    }
    /**
     * Sets the fixes property value. Total number of fixes applied across all APIs
     * @param value Value to set for the fixes property.
     */
    public void setFixes(@jakarta.annotation.Nullable final Integer value) {
        this.fixes = value;
    }
    /**
     * Sets the invalid property value. Number of newly invalid APIs
     * @param value Value to set for the invalid property.
     */
    public void setInvalid(@jakarta.annotation.Nullable final Integer value) {
        this.invalid = value;
    }
    /**
     * Sets the issues property value. Open GitHub issues on our main repo
     * @param value Value to set for the issues property.
     */
    public void setIssues(@jakarta.annotation.Nullable final Integer value) {
        this.issues = value;
    }
    /**
     * Sets the numAPIs property value. Number of unique APIs
     * @param value Value to set for the numAPIs property.
     */
    public void setNumAPIs(@jakarta.annotation.Nullable final Integer value) {
        this.numAPIs = value;
    }
    /**
     * Sets the numDrivers property value. Number of methods of API retrieval
     * @param value Value to set for the numDrivers property.
     */
    public void setNumDrivers(@jakarta.annotation.Nullable final Integer value) {
        this.numDrivers = value;
    }
    /**
     * Sets the numEndpoints property value. Total number of endpoints inside all definitions
     * @param value Value to set for the numEndpoints property.
     */
    public void setNumEndpoints(@jakarta.annotation.Nullable final Integer value) {
        this.numEndpoints = value;
    }
    /**
     * Sets the numProviders property value. Number of API providers in directory
     * @param value Value to set for the numProviders property.
     */
    public void setNumProviders(@jakarta.annotation.Nullable final Integer value) {
        this.numProviders = value;
    }
    /**
     * Sets the numSpecs property value. Number of API definitions including different versions of the same API
     * @param value Value to set for the numSpecs property.
     */
    public void setNumSpecs(@jakarta.annotation.Nullable final Integer value) {
        this.numSpecs = value;
    }
    /**
     * Sets the stars property value. GitHub stars for our main repo
     * @param value Value to set for the stars property.
     */
    public void setStars(@jakarta.annotation.Nullable final Integer value) {
        this.stars = value;
    }
    /**
     * Sets the thisWeek property value. Summary totals for the last 7 days
     * @param value Value to set for the thisWeek property.
     */
    public void setThisWeek(@jakarta.annotation.Nullable final MetricsThisWeek value) {
        this.thisWeek = value;
    }
    /**
     * Sets the unofficial property value. Number of unofficial APIs
     * @param value Value to set for the unofficial property.
     */
    public void setUnofficial(@jakarta.annotation.Nullable final Integer value) {
        this.unofficial = value;
    }
    /**
     * Sets the unreachable property value. Number of unreachable (4XX,5XX status) APIs
     * @param value Value to set for the unreachable property.
     */
    public void setUnreachable(@jakarta.annotation.Nullable final Integer value) {
        this.unreachable = value;
    }
}
