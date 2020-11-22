package application

import env "github.com/suzuito/common-env"

// Application ...
type Application struct {
	DirData      string
	GCPProjectID string
}

// NewApplication ...
func NewApplication() *Application {
	a := Application{}
	a.DirData = env.GetenvAsString("DIR_DATA", "data")
	a.GCPProjectID = env.GetenvAsString("GCP_PROJECT_ID", "")
	return &a
}
