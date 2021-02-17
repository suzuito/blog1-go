package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"golang.org/x/xerrors"
)

func (i *Impl) ServeFront(
	ctx context.Context,
	urlPrerendering string,
	originFront string,
	w http.ResponseWriter,
	r *http.Request,
) error {
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
		u := fmt.Sprintf("%s%s", originFront, r.RequestURI)
		bodyString := ""
		err := GeneratePrerenderingPage(
			ctx,
			urlPrerendering,
			u,
			&bodyString,
		)
		if err == nil {
			body = []byte(bodyString)
		}
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

func GeneratePrerenderingPage(
	ctx context.Context,
	urlPrerendering string,
	urlTarget string,
	returned *string,
) error {
	url := fmt.Sprintf("%s/render/%s", urlPrerendering, urlTarget)
	res, err := retryablehttp.Get(url)
	if err != nil {
		return xerrors.Errorf("Cannot get '%s': %w", url, err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return xerrors.Errorf("Cannot read all '%s': %w", url, err)
	}
	if res.StatusCode != http.StatusOK {
		return xerrors.Errorf("Cannot get '%s' with status %d, with body '%s': %w", url, res.StatusCode, string(body), err)
	}
	*returned = string(body)
	return nil
}

func (i *Impl) SetPrerenderingPages(
	ctx context.Context,
	urlPrerendering string,
	urlsTarget []string,
) error {
	for _, urlTarget := range urlsTarget {
		u, err := url.Parse(urlTarget)
		if err != nil {
			return xerrors.Errorf("Cannot parser url '%s' : %w", urlTarget, err)
		}
		resource := ""
		if err := GeneratePrerenderingPage(ctx, urlPrerendering, urlTarget, &resource); err != nil {
			return xerrors.Errorf("Cannot generate prerendering page '%s': %w", urlTarget, err)
		}
		p := fmt.Sprintf("rendering%s", u.Path)
		i.logger.Infof("Upload '%s'", p)
		if err := i.storage.UploadHTML(ctx, p, resource); err != nil {
			return xerrors.Errorf("Cannot upload html '%s': %w", p, err)
		}
	}
	return nil
}
