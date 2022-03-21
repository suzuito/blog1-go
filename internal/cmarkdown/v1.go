package cmarkdown

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"golang.org/x/xerrors"
)

type astTransformerAddLinkBlank struct {
}

func (a *astTransformerAddLinkBlank) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if n.Kind() != ast.KindLink {
			return ast.WalkContinue, nil
		}
		n.SetAttributeString("target", []byte("blank"))
		return ast.WalkContinue, nil
	})
}

type astTransformerImage struct {
}

func (a *astTransformerImage) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if n.Kind() != ast.KindImage {
			return ast.WalkContinue, nil
		}
		n.SetAttributeString("class", []byte("md-image"))
		return ast.WalkContinue, nil
	})
}

type astTransformerHeading struct {
}

func (a *astTransformerHeading) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if n.Kind() != ast.KindHeading {
			return ast.WalkContinue, nil
		}
		n.SetAttributeString("class", []byte("md-heading"))
		return ast.WalkContinue, nil
	})
}

type V1 struct {
	md goldmark.Markdown
}

func NewV1() *V1 {
	r := V1{}
	r.md = goldmark.New(
		goldmark.WithExtensions(mathjax.MathJax),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
			parser.WithASTTransformers(
				util.Prioritized(&astTransformerAddLinkBlank{}, 1),
				util.Prioritized(&astTransformerImage{}, 1),
				util.Prioritized(&astTransformerHeading{}, 1),
			),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
	return &r
}

func (g *V1) Convert(
	ctx context.Context,
	source []byte,
	w io.Writer,
	meta *CMMeta,
	tocs *[]CMTOC,
	images *[]CMImage,
) error {
	sourceWithoutMeta := []byte{}
	if err := parseMeta(source, meta, &sourceWithoutMeta); err != nil {
		return xerrors.Errorf("parseMeta : %w", err)
	}
	tempHTML := bytes.NewBufferString("")
	if err := g.md.Convert(sourceWithoutMeta, tempHTML); err != nil {
		return xerrors.Errorf("Cannot convert : %w", err)
	}
	tempDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(tempHTML.Bytes()))
	if err != nil {
		return xerrors.Errorf("Cannot convert to html : %w", err)
	}
	tempDoc.Find("pre").Each(func(i int, s *goquery.Selection) {
		s.SetAttr("class", "code-block")
		s.SetAttr("style", "width: 100%; overflow: scroll;")
	})
	tempDoc.Find("img.md-image").Each(func(i int, s *goquery.Selection) {
		s.SetAttr("style", "width: 100%;")
		*images = append(*images, CMImage{
			URL: s.AttrOr("src", ""),
		})
	})
	tempDoc.Find(".md-heading").Each(func(i int, s *goquery.Selection) {
		toc := CMTOC{
			Name:  s.Text(),
			ID:    s.AttrOr("id", ""),
			Level: NewTOCLevel(goquery.NodeName(s)),
		}
		*tocs = append(*tocs, toc)
	})
	returned, err := tempDoc.Html()
	if err != nil {
		return xerrors.Errorf("Cannot HTML : %w", err)
	}
	returned = strings.Replace(returned, "<html><head></head><body>", "", 1)
	returned = strings.Replace(returned, "</body></html>", "", 1)
	fmt.Fprint(w, returned)
	return nil
}
