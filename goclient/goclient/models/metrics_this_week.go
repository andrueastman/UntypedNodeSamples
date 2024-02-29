package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Metrics_thisWeek summary totals for the last 7 days
type Metrics_thisWeek struct {
    // APIs added in the last week
    added *int32
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // APIs updated in the last week
    updated *int32
}
// NewMetrics_thisWeek instantiates a new Metrics_thisWeek and sets the default values.
func NewMetrics_thisWeek()(*Metrics_thisWeek) {
    m := &Metrics_thisWeek{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMetrics_thisWeekFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMetrics_thisWeekFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMetrics_thisWeek(), nil
}
// GetAdded gets the added property value. APIs added in the last week
// returns a *int32 when successful
func (m *Metrics_thisWeek) GetAdded()(*int32) {
    return m.added
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Metrics_thisWeek) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Metrics_thisWeek) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["added"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdded(val)
        }
        return nil
    }
    res["updated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdated(val)
        }
        return nil
    }
    return res
}
// GetUpdated gets the updated property value. APIs updated in the last week
// returns a *int32 when successful
func (m *Metrics_thisWeek) GetUpdated()(*int32) {
    return m.updated
}
// Serialize serializes information the current object
func (m *Metrics_thisWeek) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("added", m.GetAdded())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("updated", m.GetUpdated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdded sets the added property value. APIs added in the last week
func (m *Metrics_thisWeek) SetAdded(value *int32)() {
    m.added = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Metrics_thisWeek) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetUpdated sets the updated property value. APIs updated in the last week
func (m *Metrics_thisWeek) SetUpdated(value *int32)() {
    m.updated = value
}
type Metrics_thisWeekable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdded()(*int32)
    GetUpdated()(*int32)
    SetAdded(value *int32)()
    SetUpdated(value *int32)()
}
