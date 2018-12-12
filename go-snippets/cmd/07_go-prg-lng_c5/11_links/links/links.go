package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// Extract makes http Get request and parses the response body
// returns slice of links in the response
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("link Extract: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("link Extract : response code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("link Extract: %s", err)
	}

	var links []string

	vistNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(attr.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, vistNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
