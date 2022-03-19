package bgin

import (
	"github.com/suzuito/blog1-go/pkg/setting"
)

func htmlGlobal(env *setting.Environment) map[string]string {
	return map[string]string{
		"Title": "otiuzu pages",
		"GA":    "",
	}
}
