package main

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
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
	err := sentry.Init(sentry.ClientOptions{
		Environment: "minilla",
		Release:     os.Getenv("COMMIT_SHA"),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)
	if err := f3(); err != nil {
		sentry.CaptureException(err)
	}
}
