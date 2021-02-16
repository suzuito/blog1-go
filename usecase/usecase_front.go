package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/xerrors"
)

func (i *Impl) ServeFront(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var err error
	body := []byte{}
	headers := make(map[string]string)
	status := 0
	defer (func() {
		m := struct {
			Path   string
			Status int
		}{
			Path: r.URL.Path,
		}
		if err != nil {
			if xerrors.Is(err, ErrNotFound) {
				status = http.StatusNotFound
				body = []byte("Not found")
			} else {
				status = http.StatusInternalServerError
				body = []byte("Internal server error")
				i.logger.Infof("%+v", err)
			}
		} else {
			status = http.StatusOK
		}
		for k, v := range headers {
			w.Header().Add(k, v)
		}
		w.WriteHeader(status)
		w.Write(body)
		m.Status = status
		b, _ := json.Marshal(m)
		i.logger.Infof("%s", string(b))
	})()
	p := r.URL.Path
	if strings.Contains(r.UserAgent(), "Googlebot") {
		err = i.storage.GetFileAsHTTPResponse(
			ctx,
			fmt.Sprintf("rendering%s", p),
			&body,
			headers,
		)
		return err
	}
	if p == "/" {
		p = "/index.html"
	}
	err = i.storage.GetFileAsHTTPResponse(
		ctx,
		fmt.Sprintf("app%s", p),
		&body,
		headers,
	)
	if err != nil {
		if xerrors.Is(err, ErrNotFound) {
			p = "/index.html"
			err = i.storage.GetFileAsHTTPResponse(
				ctx,
				fmt.Sprintf("app%s", p),
				&body,
				headers,
			)
		}
	}
	return err
}