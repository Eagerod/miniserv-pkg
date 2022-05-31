package html_xpath

import (
	"fmt"
	"net/url"
)

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func TextAttributeFromQueryUrl(aUrl, xpath, attribute string) ([]string, error) {
	rv := []string{}

	theUrl, err := url.Parse(aUrl)
	if err != nil {
		return rv, err
	}

	var rootNode *html.Node
	if theUrl.Scheme == "file" || theUrl.Scheme == "" {
		rootNode, err = htmlquery.LoadDoc(theUrl.Path)
	} else {
		rootNode, err = htmlquery.LoadURL(aUrl)
	}

	if err != nil {
		return rv, fmt.Errorf("failed to load url: %s - %s", aUrl, err.Error())
	}

	elems, err := htmlquery.QueryAll(rootNode, xpath)
	if err != nil {
		return rv, fmt.Errorf("failed to query xpath: %s - %s", xpath, err.Error())
	}

	for _, elem := range elems {
		for _, attr := range elem.Attr {
			if attr.Key == attribute {
				rv = append(rv, htmlquery.SelectAttr(elem, attribute))
				break
			}
		}
	}

	return rv, nil
}
