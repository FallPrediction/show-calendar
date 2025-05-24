package utils

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type OpenGraph struct{}

type OpenGraphMeta struct {
	Title       string
	Type        string
	Image       string
	Url         string
	Description string
}

func (o *OpenGraph) Fetch(url string) (OpenGraphMeta, error) {
	meta := OpenGraphMeta{}
	doc, err := o.getDoc(url)
	if err != nil {
		return meta, err
	}
	for node := range doc.Descendants() {
		if node.Type == html.ElementNode && node.Data == "meta" {
			property, value := o.getOgContent(node.Attr)
			o.setField(&meta, property, value)
		}
	}
	return meta, nil
}

func (o *OpenGraph) getDoc(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (o *OpenGraph) getOgContent(attrs []html.Attribute) (string, string) {
	isOg := false
	property := ""
	for _, attr := range attrs {
		if attr.Key == "property" && strings.HasPrefix(attr.Val, "og:") {
			isOg = true
			property = attr.Val
			break
		}
	}
	if isOg {
		for _, attr := range attrs {
			if attr.Key == "content" {
				return property, attr.Val
			}
		}
	}
	return property, ""
}

func (o *OpenGraph) setField(meta *OpenGraphMeta, property, value string) {
	switch property {
	case "og:title":
		meta.Title = value
	case "og:type":
		meta.Type = value
	case "og:image":
		meta.Image = value
	case "og:url":
		meta.Url = value
	case "og:description":
		meta.Description = value
	}
}
