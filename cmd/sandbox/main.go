package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/internal/bgcp/cloudlogging"
)

func f1() error {
	return errors.New("hoge")
}

func f2() error {
	err := f1()
	return errors.Wrapf(err, "fuga")
}

func main() {
	if err := f2(); err != nil {
		cloudlogging.Error(err)
		cloudlogging.Error(fmt.Errorf("hoge"))
	}
}
