package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Metrics list of basic metrics
type Metrics struct {
    // Data used for charting etc
    datasets i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
    // Percentage of all APIs where auto fixes have been applied
    fixedPct *int32
    // Total number of fixes applied across all APIs
    fixes *int32
    // Number of newly invalid APIs
    invalid *int32
    // Open GitHub issues on our main repo
    issues *int32
    // Number of unique APIs
    numAPIs *int32
    // Number of methods of API retrieval
    numDrivers *int32
    // Total number of endpoints inside all definitions
    numEndpoints *int32
    // Number of API providers in directory
    numProviders *int32
    // Number of API definitions including different versions of the same API
    numSpecs *int32
    // GitHub stars for our main repo
    stars *int32
    // Summary totals for the last 7 days
    thisWeek Metrics_thisWeekable
    // Number of unofficial APIs
    unofficial *int32
    // Number of unreachable (4XX,5XX status) APIs
    unreachable *int32
}
// NewMetrics instantiates a new Metrics and sets the default values.
func NewMetrics()(*Metrics) {
    m := &Metrics{
    }
    return m
}
// CreateMetricsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMetricsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMetrics(), nil
}
// GetDatasets gets the datasets property value. Data used for charting etc
// returns a UntypedNodeable when successful
func (m *Metrics) GetDatasets()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.datasets
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Metrics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["datasets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatasets(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    res["fixedPct"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFixedPct(val)
        }
        return nil
    }
    res["fixes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFixes(val)
        }
        return nil
    }
    res["invalid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvalid(val)
        }
        return nil
    }
    res["issues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssues(val)
        }
        return nil
    }
    res["numAPIs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumAPIs(val)
        }
        return nil
    }
    res["numDrivers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumDrivers(val)
        }
        return nil
    }
    res["numEndpoints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumEndpoints(val)
        }
        return nil
    }
    res["numProviders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumProviders(val)
        }
        return nil
    }
    res["numSpecs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumSpecs(val)
        }
        return nil
    }
    res["stars"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStars(val)
        }
        return nil
    }
    res["thisWeek"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMetrics_thisWeekFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThisWeek(val.(Metrics_thisWeekable))
        }
        return nil
    }
    res["unofficial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnofficial(val)
        }
        return nil
    }
    res["unreachable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnreachable(val)
        }
        return nil
    }
    return res
}
// GetFixedPct gets the fixedPct property value. Percentage of all APIs where auto fixes have been applied
// returns a *int32 when successful
func (m *Metrics) GetFixedPct()(*int32) {
    return m.fixedPct
}
// GetFixes gets the fixes property value. Total number of fixes applied across all APIs
// returns a *int32 when successful
func (m *Metrics) GetFixes()(*int32) {
    return m.fixes
}
// GetInvalid gets the invalid property value. Number of newly invalid APIs
// returns a *int32 when successful
func (m *Metrics) GetInvalid()(*int32) {
    return m.invalid
}
// GetIssues gets the issues property value. Open GitHub issues on our main repo
// returns a *int32 when successful
func (m *Metrics) GetIssues()(*int32) {
    return m.issues
}
// GetNumAPIs gets the numAPIs property value. Number of unique APIs
// returns a *int32 when successful
func (m *Metrics) GetNumAPIs()(*int32) {
    return m.numAPIs
}
// GetNumDrivers gets the numDrivers property value. Number of methods of API retrieval
// returns a *int32 when successful
func (m *Metrics) GetNumDrivers()(*int32) {
    return m.numDrivers
}
// GetNumEndpoints gets the numEndpoints property value. Total number of endpoints inside all definitions
// returns a *int32 when successful
func (m *Metrics) GetNumEndpoints()(*int32) {
    return m.numEndpoints
}
// GetNumProviders gets the numProviders property value. Number of API providers in directory
// returns a *int32 when successful
func (m *Metrics) GetNumProviders()(*int32) {
    return m.numProviders
}
// GetNumSpecs gets the numSpecs property value. Number of API definitions including different versions of the same API
// returns a *int32 when successful
func (m *Metrics) GetNumSpecs()(*int32) {
    return m.numSpecs
}
// GetStars gets the stars property value. GitHub stars for our main repo
// returns a *int32 when successful
func (m *Metrics) GetStars()(*int32) {
    return m.stars
}
// GetThisWeek gets the thisWeek property value. Summary totals for the last 7 days
// returns a Metrics_thisWeekable when successful
func (m *Metrics) GetThisWeek()(Metrics_thisWeekable) {
    return m.thisWeek
}
// GetUnofficial gets the unofficial property value. Number of unofficial APIs
// returns a *int32 when successful
func (m *Metrics) GetUnofficial()(*int32) {
    return m.unofficial
}
// GetUnreachable gets the unreachable property value. Number of unreachable (4XX,5XX status) APIs
// returns a *int32 when successful
func (m *Metrics) GetUnreachable()(*int32) {
    return m.unreachable
}
// Serialize serializes information the current object
func (m *Metrics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("datasets", m.GetDatasets())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("fixedPct", m.GetFixedPct())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("fixes", m.GetFixes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("invalid", m.GetInvalid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("issues", m.GetIssues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("numAPIs", m.GetNumAPIs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("numDrivers", m.GetNumDrivers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("numEndpoints", m.GetNumEndpoints())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("numProviders", m.GetNumProviders())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("numSpecs", m.GetNumSpecs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("stars", m.GetStars())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("thisWeek", m.GetThisWeek())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unofficial", m.GetUnofficial())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unreachable", m.GetUnreachable())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDatasets sets the datasets property value. Data used for charting etc
func (m *Metrics) SetDatasets(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.datasets = value
}
// SetFixedPct sets the fixedPct property value. Percentage of all APIs where auto fixes have been applied
func (m *Metrics) SetFixedPct(value *int32)() {
    m.fixedPct = value
}
// SetFixes sets the fixes property value. Total number of fixes applied across all APIs
func (m *Metrics) SetFixes(value *int32)() {
    m.fixes = value
}
// SetInvalid sets the invalid property value. Number of newly invalid APIs
func (m *Metrics) SetInvalid(value *int32)() {
    m.invalid = value
}
// SetIssues sets the issues property value. Open GitHub issues on our main repo
func (m *Metrics) SetIssues(value *int32)() {
    m.issues = value
}
// SetNumAPIs sets the numAPIs property value. Number of unique APIs
func (m *Metrics) SetNumAPIs(value *int32)() {
    m.numAPIs = value
}
// SetNumDrivers sets the numDrivers property value. Number of methods of API retrieval
func (m *Metrics) SetNumDrivers(value *int32)() {
    m.numDrivers = value
}
// SetNumEndpoints sets the numEndpoints property value. Total number of endpoints inside all definitions
func (m *Metrics) SetNumEndpoints(value *int32)() {
    m.numEndpoints = value
}
// SetNumProviders sets the numProviders property value. Number of API providers in directory
func (m *Metrics) SetNumProviders(value *int32)() {
    m.numProviders = value
}
// SetNumSpecs sets the numSpecs property value. Number of API definitions including different versions of the same API
func (m *Metrics) SetNumSpecs(value *int32)() {
    m.numSpecs = value
}
// SetStars sets the stars property value. GitHub stars for our main repo
func (m *Metrics) SetStars(value *int32)() {
    m.stars = value
}
// SetThisWeek sets the thisWeek property value. Summary totals for the last 7 days
func (m *Metrics) SetThisWeek(value Metrics_thisWeekable)() {
    m.thisWeek = value
}
// SetUnofficial sets the unofficial property value. Number of unofficial APIs
func (m *Metrics) SetUnofficial(value *int32)() {
    m.unofficial = value
}
// SetUnreachable sets the unreachable property value. Number of unreachable (4XX,5XX status) APIs
func (m *Metrics) SetUnreachable(value *int32)() {
    m.unreachable = value
}
type Metricsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDatasets()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    GetFixedPct()(*int32)
    GetFixes()(*int32)
    GetInvalid()(*int32)
    GetIssues()(*int32)
    GetNumAPIs()(*int32)
    GetNumDrivers()(*int32)
    GetNumEndpoints()(*int32)
    GetNumProviders()(*int32)
    GetNumSpecs()(*int32)
    GetStars()(*int32)
    GetThisWeek()(Metrics_thisWeekable)
    GetUnofficial()(*int32)
    GetUnreachable()(*int32)
    SetDatasets(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
    SetFixedPct(value *int32)()
    SetFixes(value *int32)()
    SetInvalid(value *int32)()
    SetIssues(value *int32)()
    SetNumAPIs(value *int32)()
    SetNumDrivers(value *int32)()
    SetNumEndpoints(value *int32)()
    SetNumProviders(value *int32)()
    SetNumSpecs(value *int32)()
    SetStars(value *int32)()
    SetThisWeek(value Metrics_thisWeekable)()
    SetUnofficial(value *int32)()
    SetUnreachable(value *int32)()
}
