package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		t, err := title(url)
		if err != nil {
			fmt.Println("main:", err)
		}
		fmt.Println("title:", t)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
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

func title(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("title: http get error: %v", err)
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return "", fmt.Errorf("title: %s has type %s", url, ct)
	}
	fmt.Println(ct)
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("title: html parse error: %v", err)
	}

	t, err := soleTitle(doc)

	if err != nil {
		return "", fmt.Errorf("title: error: %v", err)
	}

	return t, nil
}

func soleTitle(n *html.Node) (string, error) {
	t := ""
	var err error
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			err = fmt.Errorf("multiple title elements")
			fmt.Println(err)
		default:
			panic(p)
		}
	}()

	forEachNode(n, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if t != "" {
				// multiple title elements
				panic(bailout{})
			}
			t = n.FirstChild.Data
		}
	}, nil)
	return t, err
}
