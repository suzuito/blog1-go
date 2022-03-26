package bhtml

import (
	"context"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/pkg/entity"
)

type MediaFetcher struct {
	Cli *http.Client
}

func (m *MediaFetcher) Fetch(
	ctx context.Context,
	src string,
	images *[]entity.ArticleImage,
) error {
	d, err := goquery.NewDocumentFromReader(strings.NewReader(src))
	if err != nil {
		return errors.Wrapf(err, "cannot convert to html")
	}

	srcURLs := []string{}
	srcURLMap := map[string]*goquery.Selection{}
	d.Find("img").Each(func(i int, s *goquery.Selection) {
		srcURL := s.AttrOr("src", "")
		if srcURL == "" {
			return
		}
		srcURLs = append(srcURLs, srcURL)
		srcURLMap[srcURL] = s
	})

	for _, srcURL := range srcURLs {
		img, err := fetchSrcURL(ctx, m.Cli, srcURL)
		if err != nil {
			continue
		}
		rect := img.Bounds()
		point := rect.Size()
		aimg := entity.ArticleImage{
			RealHeight: point.Y,
			RealWidth:  point.X,
			URL:        srcURL,
		}
		if s, exists := srcURLMap[srcURL]; exists {
			if width, exists := s.Attr("width"); exists {
				if iwidth, err := strconv.Atoi(width); err == nil {
					aimg.Width = iwidth
				}
			}
			if height, exists := s.Attr("height"); exists {
				if iheight, err := strconv.Atoi(height); err == nil {
					aimg.Height = iheight
				}
			}
		}
		*images = append(*images, aimg)
	}

	return nil
}

type imageProp struct {
	RealHeight int
	RealWidth  int
	URL        string
	Err        error
}

var (
	invaliURLErr    = fmt.Errorf("invalid url")
	requestErr      = fmt.Errorf("cannot request to url")
	httpErrorErr    = fmt.Errorf("http error")
	imageDeocodeErr = fmt.Errorf("image decode error")
)

func fetchSrcURL(
	ctx context.Context,
	cli *http.Client,
	srcURL string,
) (image.Image, error) {
	uri, err := url.Parse(srcURL)
	if err != nil {
		return nil, errors.Wrapf(invaliURLErr, "%s : %+v", srcURL, err)
	}
	if uri.Scheme == "" {
		uri.Scheme = "https"
	}
	r, err := cli.Get(uri.String())
	if err != nil {
		return nil, errors.Wrapf(requestErr, "%s : %+v", srcURL, err)
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(r.Body)
		return nil, errors.Wrapf(httpErrorErr, "%s status=%d body=%s", srcURL, string(body))
	}
	img, _, err := image.Decode(r.Body)
	if err != nil {
		return nil, errors.Wrapf(imageDeocodeErr, "%s : %+v", srcURL, err)
	}
	return img, nil
}
