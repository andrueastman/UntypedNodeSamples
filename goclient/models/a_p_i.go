package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// API meta information about API
type API struct {
    // Timestamp when the API was first added to the directory
    added *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Recommended version
    preferred *string
    // List of supported versions of the API
    versions API_versionsable
}
// NewAPI instantiates a new API and sets the default values.
func NewAPI()(*API) {
    m := &API{
    }
    return m
}
// CreateAPIFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAPIFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAPI(), nil
}
// GetAdded gets the added property value. Timestamp when the API was first added to the directory
// returns a *Time when successful
func (m *API) GetAdded()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.added
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *API) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["added"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdded(val)
        }
        return nil
    }
    res["preferred"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferred(val)
        }
        return nil
    }
    res["versions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAPI_versionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersions(val.(API_versionsable))
        }
        return nil
    }
    return res
}
// GetPreferred gets the preferred property value. Recommended version
// returns a *string when successful
func (m *API) GetPreferred()(*string) {
    return m.preferred
}
// GetVersions gets the versions property value. List of supported versions of the API
// returns a API_versionsable when successful
func (m *API) GetVersions()(API_versionsable) {
    return m.versions
}
// Serialize serializes information the current object
func (m *API) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("added", m.GetAdded())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("preferred", m.GetPreferred())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("versions", m.GetVersions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdded sets the added property value. Timestamp when the API was first added to the directory
func (m *API) SetAdded(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.added = value
}
// SetPreferred sets the preferred property value. Recommended version
func (m *API) SetPreferred(value *string)() {
    m.preferred = value
}
// SetVersions sets the versions property value. List of supported versions of the API
func (m *API) SetVersions(value API_versionsable)() {
    m.versions = value
}
type APIable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdded()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPreferred()(*string)
    GetVersions()(API_versionsable)
    SetAdded(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPreferred(value *string)()
    SetVersions(value API_versionsable)()
}
