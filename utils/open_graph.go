package utils

import (
	"net/http"
	"souflair/errors"
	"souflair/initialize"
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
	if meta.Image != "" {
		meta.Image, err = o.downloadImage(meta.Image)
		if err != nil {
			return meta, err
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

func (o *OpenGraph) downloadImage(url string) (string, error) {
	logger := initialize.NewLogger()
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	ext, err := o.getExt(resp.Header.Get("content-type"))
	if err == errors.ErrInvalidContentType {
		return "", nil
	} else if err != nil {
		return "", err
	}
	uploader := NewUploader()
	fileName, err := uploader.Upload(resp.Body, ext)
	defer resp.Body.Close()
	if err != nil {
		logger.Error(err)
		return "", err
	}
	return fileName, nil
}

func (o *OpenGraph) getExt(contentType string) (string, error) {
	strArr := strings.Split(contentType, "/")
	if len(strArr) == 2 {
		if strArr[0] != "image" {
			return "", errors.ErrInvalidContentType
		}
		return strArr[1], nil
	}
	return "jpeg", nil
}
