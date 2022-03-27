package setting

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
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
	DirPathAsset     string   `envconfig:"DIR_PATH_ASSET"`
	GA               string   `envconfig:"GA"`
	SiteOrigin       string   `envconfig:"SITE_ORIGIN"`
}

var E *Environment

func init() {
	E = &Environment{}
	if err := envconfig.Process("", E); err != nil {
		panic(errors.Wrapf(err, "cannot envconfig.Process"))
	}
}
