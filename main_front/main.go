package main

import (
	"log"
	"net/http"

	gstorage "cloud.google.com/go/storage"
	"github.com/suzuito/blog1-go/bgcp/storage"
	"github.com/suzuito/blog1-go/usecase"
	env "github.com/suzuito/common-env"
	"github.com/suzuito/common-go/clogger"
)

func handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var scli *gstorage.Client
		ctx := r.Context()
		scli, err = gstorage.NewClient(ctx)
		if err != nil {
			w.Write([]byte("Cannot init"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer scli.Close()
		u := usecase.NewImpl(&clogger.LoggerPrint{}, nil, storage.New(scli, env.GetenvAsString("GCP_BUCKET", "")), nil)
		u.ServeFront(
			ctx,
			w,
			r,
		)
	}
}

func main() {
	http.HandleFunc("/", handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
