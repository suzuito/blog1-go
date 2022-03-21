package usecase

import (
	"bytes"
	"image"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	_ "image/gif"
	_ "image/jpeg"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/pkg/entity"
)

func (u *Impl) attacheArticleImages(article *entity.Article, htmlBody []byte) error {
	if err := extractImageURLs(htmlBody, &article.Images); err != nil {
		return errors.Wrapf(err, "Cannot extract image urls")
	}
	cli1 := retryablehttp.NewClient()
	cli := cli1.StandardClient()
	for i := range article.Images {
		img := article.Images[i]
		if err := u.refineArticleImage(cli, &img); err != nil {
			return errors.Wrapf(err, "Cannot refineArticleImage")
		}
		article.Images[i] = img
	}
	return nil
}

func (u *Impl) refineArticleImage(cli *http.Client, articleImage *entity.ArticleImage) error {
	uri, err := url.Parse(articleImage.URL)
	if err != nil {
		return errors.Wrapf(err, "invalid url '%s'", articleImage.URL)
	}
	if uri.Scheme == "" {
		uri.Scheme = "https"
	}
	r, err := cli.Get(uri.String())
	if err != nil {
		return errors.Wrapf(err, "cannot get url '%s'", articleImage.URL)
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(r.Body)
		return errors.Wrapf(err, "cannot get url '%s' body %s", articleImage.URL, string(body))
	}
	img, _, err := image.Decode(r.Body)
	if err != nil {
		return errors.Wrapf(err, "cannot decode url '%s'", articleImage.URL)
	}
	rect := img.Bounds()
	size := rect.Size()
	articleImage.RealHeight = size.Y
	articleImage.RealWidth = size.X
	articleImage.URL = r.Request.URL.String()
	return nil
}

func extractImageURLs(body []byte, articleImages *[]entity.ArticleImage) error {
	d, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return errors.Wrapf(err, "cannot new goquery")
	}
	d.Find("img").Each(func(i int, s *goquery.Selection) {
		articleImage := entity.ArticleImage{}
		imageURL, exists := s.Attr("src")
		if exists {
			uri, err := url.Parse(imageURL)
			if err == nil {
				if uri.Scheme == "" {
					uri.Scheme = "https"
				}
				if uri.Host == "ir-jp.amazon-adsystem.com" {
					return
				}
				articleImage.URL = uri.String()
			} else {
				articleImage.URL = imageURL
			}
		}
		width, exists := s.Attr("width")
		if exists {
			i, err := strconv.Atoi(width)
			if err == nil {
				articleImage.Width = i
			}
		}
		height, exists := s.Attr("height")
		if exists {
			i, err := strconv.Atoi(height)
			if err == nil {
				articleImage.Height = i
			}
		}
		*articleImages = append(*articleImages, articleImage)
	})
	return nil
}
