package cmarkdown

import (
	"bytes"
	"context"

	"github.com/PuerkitoBio/goquery"
	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/pkg/usecase"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
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
		if n.Kind() != ast.KindLink && n.Kind() != ast.KindAutoLink {
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
		goldmark.WithExtensions(
			mathjax.MathJax,
			extension.Linkify,
		),
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
	src string,
	dst *string,
	meta *usecase.CMMeta,
) error {
	sourceWithoutMeta := []byte{}
	if err := parseMeta(src, meta, &sourceWithoutMeta); err != nil {
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
	})
	returned, err := tempDoc.Html()
	if err != nil {
		return errors.Errorf("cannot HTML : %+v", err)
	}
	*dst = returned
	return nil
}
