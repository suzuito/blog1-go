package setting

import (
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/xerrors"
)

// Environment ...
type Environment struct {
	Env            string   `envconfig:"ENV"`
	GCPProjectID   string   `envconfig:"GCP_PROJECT_ID"`
	AllowedOrigins []string `envconfig:"ALLOWED_ORIGINS"`
	AllowedMethods []string `envconfig:"ALLOWED_METHODS"`
	GHSHA          string   `envconfig:"GH_SHA"`
}

func NewEnvironment() (*Environment, error) {
	r := Environment{}
	if err := envconfig.Process("", &r); err != nil {
		return nil, xerrors.Errorf("Cannot envconfig.Process : %w", err)
	}
	return &r, nil
}
