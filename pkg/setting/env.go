package setting

import (
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/xerrors"
)

// Environment ...
type Environment struct {
	Env              string   `envconfig:"ENV"`
	GCPProjectID     string   `envconfig:"GCP_PROJECT_ID"`
	GCPBucketArticle string   `envconfig:"GCP_BUCKET_ARTICLE"`
	AllowedOrigins   []string `envconfig:"ALLOWED_ORIGINS"`
	AllowedMethods   []string `envconfig:"ALLOWED_METHODS"`
	DirPathTemplate  string   `envconfig:"DIR_PATH_TEMPLATE"`
	DirPathCSS       string   `envconfig:"DIR_PATH_CSS"`
}

func NewEnvironment() (*Environment, error) {
	r := Environment{}
	if err := envconfig.Process("", &r); err != nil {
		return nil, xerrors.Errorf("Cannot envconfig.Process : %w", err)
	}
	return &r, nil
}
