package providersjson

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ProvidersGetResponseable instead.
type ProvidersResponse struct {
    ProvidersGetResponse
}
// NewProvidersResponse instantiates a new ProvidersResponse and sets the default values.
func NewProvidersResponse()(*ProvidersResponse) {
    m := &ProvidersResponse{
        ProvidersGetResponse: *NewProvidersGetResponse(),
    }
    return m
}
// CreateProvidersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvidersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvidersResponse(), nil
}
// Deprecated: This class is obsolete. Use ProvidersGetResponseable instead.
type ProvidersResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvidersGetResponseable
}