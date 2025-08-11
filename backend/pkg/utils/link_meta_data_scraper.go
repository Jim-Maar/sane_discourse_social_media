package utils

import (
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type LinkMetadata struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	SiteName    string `json:"site_name"`
	Author      string `json:"author"`
	URL         string `json:"url"`
	Type        string `json:"type"`
}

func ScrapeMetadata(url string) (*LinkMetadata, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; SaneDiscourse/1.0)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	// Initialize metadata with empty strings for all fields
	metadata := &LinkMetadata{
		Title:       "",
		Description: "",
		ImageURL:    "",
		SiteName:    "",
		Author:      "",
		URL:         url,
		Type:        "",
	}

	var parseNode func(*html.Node)
	parseNode = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "meta":
				extractMetaTag(n, metadata)
			case "title":
				if n.FirstChild != nil && metadata.Title == "" {
					metadata.Title = strings.TrimSpace(n.FirstChild.Data)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseNode(c)
		}
	}

	parseNode(doc)

	// Ensure all fields have empty strings if not found (already done in initialization)
	return metadata, nil
}

func extractMetaTag(n *html.Node, metadata *LinkMetadata) {
	var property, name, content string

	for _, attr := range n.Attr {
		switch attr.Key {
		case "property":
			property = attr.Val
		case "name":
			name = attr.Val
		case "content":
			content = attr.Val
		}
	}

	switch property {
	case "og:title":
		if metadata.Title == "" {
			metadata.Title = content
		}
	case "og:description":
		if metadata.Description == "" {
			metadata.Description = content
		}
	case "og:image":
		if metadata.ImageURL == "" {
			metadata.ImageURL = content
		}
	case "og:site_name":
		if metadata.SiteName == "" {
			metadata.SiteName = content
		}
	case "og:type":
		if metadata.Type == "" {
			metadata.Type = content
		}
	case "article:author":
		if metadata.Author == "" {
			metadata.Author = content
		}
	}

	switch name {
	case "twitter:title":
		if metadata.Title == "" {
			metadata.Title = content
		}
	case "twitter:description":
		if metadata.Description == "" {
			metadata.Description = content
		}
	case "twitter:image":
		if metadata.ImageURL == "" {
			metadata.ImageURL = content
		}
	case "twitter:creator":
		if metadata.Author == "" {
			metadata.Author = strings.TrimPrefix(content, "@")
		}
	case "description":
		if metadata.Description == "" {
			metadata.Description = content
		}
	case "author":
		if metadata.Author == "" {
			metadata.Author = content
		}
	}
}
