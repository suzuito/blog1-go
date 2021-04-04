package application

import (
	"context"

	env "github.com/suzuito/common-env"
	"github.com/suzuito/common-go/cgin"
)

// Application ...
type Application struct {
	*cgin.ApplicationGinImpl
	DirData      string
	GCPProjectID string
	GCPBucket    string
}

// NewApplication ...
func NewApplication(ctx context.Context) (*Application, error) {
	parent, err := cgin.NewApplicationGinImpl(ctx)
	if err != nil {
		return nil, err
	}
	a := Application{
		ApplicationGinImpl: parent,
	}
	a.DirData = env.GetenvAsString("DIR_DATA", "data")
	a.GCPProjectID = env.GetenvAsString("GCP_PROJECT_ID", "")
	a.GCPBucket = env.GetenvAsString("GCP_BUCKET", "")
	return &a, nil
}
