package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/errorreporting"
	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/internal/bgcp/cloudlogging"
)

func f1() error {
	return errors.New("hoge1")
}

func f2() error {
	err := f1()
	return errors.Wrapf(err, "hoge2")
}

func f3() error {
	err := f2()
	return errors.Wrapf(err, "hoge3")
}

func main() {
	ctx := context.Background()
	cli, err := errorreporting.NewClient(ctx, "suzuito-minilla", errorreporting.Config{
		ServiceName: "blog",
		OnError: func(err error) {
			fmt.Println(err)
		},
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	if err := f3(); err != nil {
		// cloudlogging.Error(err)
		v := cloudlogging.NewMessageInPayloadFromError(err)
		cli.Report(errorreporting.Entry{
			Error: err,
			Stack: []byte(v),
		})
	}
}
