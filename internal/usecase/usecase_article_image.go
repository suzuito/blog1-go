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
	"github.com/suzuito/blog1-go/internal/entity/model"
	"golang.org/x/xerrors"
)

func (u *Impl) attacheArticleImages(article *model.Article, htmlBody []byte) error {
	if err := extractImageURLs(htmlBody, &article.Images); err != nil {
		return xerrors.Errorf("Cannot extract image urls : %w", err)
	}
	cli1 := retryablehttp.NewClient()
	cli := cli1.StandardClient()
	for i := range article.Images {
		img := article.Images[i]
		if err := u.refineArticleImage(cli, &img); err != nil {
			return xerrors.Errorf("Cannot refineArticleImage : %w", err)
		}
		article.Images[i] = img
	}
	return nil
}

func (u *Impl) refineArticleImage(cli *http.Client, articleImage *model.ArticleImage) error {
	uri, err := url.Parse(articleImage.URL)
	if err != nil {
		return xerrors.Errorf("Invalid url '%s' : %w", articleImage.URL, err)
	}
	if uri.Scheme == "" {
		uri.Scheme = "https"
	}
	r, err := cli.Get(uri.String())
	if err != nil {
		return xerrors.Errorf("Cannot get url '%s' : %w", articleImage.URL, err)
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(r.Body)
		return xerrors.Errorf("Cannot get url '%s' body %s : %w", articleImage.URL, string(body), err)
	}
	img, _, err := image.Decode(r.Body)
	if err != nil {
		return xerrors.Errorf("Cannot decode url '%s' : %w", articleImage.URL, err)
	}
	rect := img.Bounds()
	size := rect.Size()
	articleImage.RealHeight = size.Y
	articleImage.RealWidth = size.X
	articleImage.URL = r.Request.URL.String()
	return nil
}

func extractImageURLs(body []byte, articleImages *[]model.ArticleImage) error {
	d, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return xerrors.Errorf("Cannot new goquery : %w", err)
	}
	d.Find("img").Each(func(i int, s *goquery.Selection) {
		articleImage := model.ArticleImage{}
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
