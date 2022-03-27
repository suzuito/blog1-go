package cmarkdown

import (
	"bufio"
	"strings"

	"github.com/suzuito/blog1-go/pkg/usecase"
	"golang.org/x/xerrors"
	"gopkg.in/yaml.v2"
)

func parseMeta(source string, embedMeta *usecase.CMMeta, sourceWithOutMeta *[]byte) error {
	s := bufio.NewScanner(strings.NewReader(source))
	isMetaBlock := false
	isMetaBlockDone := false
	metaBlock := ""
	notMetaBlock := ""
	for s.Scan() {
		l := s.Text()
		if strings.HasPrefix(l, "---") && !isMetaBlockDone {
			if !isMetaBlock {
				isMetaBlock = true
				continue
			}
			isMetaBlock = false
			isMetaBlockDone = true
			continue
		}
		if isMetaBlock {
			metaBlock += l + "\n"
		} else {
			notMetaBlock += l + "\n"
		}
	}
	if !isMetaBlockDone {
		return xerrors.Errorf("Meta data is not found : %w", usecase.ErrMetaNotFound)
	}
	if err := yaml.Unmarshal([]byte(metaBlock), &embedMeta); err != nil {
		return xerrors.Errorf("Cannot parse yaml block '%s' : %w", metaBlock, err)
	}
	*sourceWithOutMeta = []byte(notMetaBlock)
	return nil
}
