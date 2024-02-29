package item

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ServicesJsonServicesGetResponseable instead.
type ServicesJsonServicesResponse struct {
    ServicesJsonServicesGetResponse
}
// NewServicesJsonServicesResponse instantiates a new ServicesJsonServicesResponse and sets the default values.
func NewServicesJsonServicesResponse()(*ServicesJsonServicesResponse) {
    m := &ServicesJsonServicesResponse{
        ServicesJsonServicesGetResponse: *NewServicesJsonServicesGetResponse(),
    }
    return m
}
// CreateServicesJsonServicesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServicesJsonServicesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicesJsonServicesResponse(), nil
}
// Deprecated: This class is obsolete. Use ServicesJsonServicesGetResponseable instead.
type ServicesJsonServicesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServicesJsonServicesGetResponseable
}
