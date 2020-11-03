package linkfinder

import (
	"fmt"
	"io"

	"github.com/antchfx/htmlquery"
)

func FindLinks(html io.Reader, xpath string) ([]string, error) {

	node, err := htmlquery.Parse(html)

	if err != nil {
		return nil, err
	}

	nodes, err := htmlquery.QueryAll(node, xpath)

	if err != nil {
		return nil, err
	}

	for _, node := range nodes {
		// fmt.Printf("%T\n %v\n %#v\n %q\n", node, node, node, node)
		for _, attribute := range node.Attr {
			fmt.Printf("%s=%s\n", attribute.Key, attribute.Val)
		}
	}

	_, _ = node, err

	//go htmlquery.QuerySelectorAll(top *html.Node, selector *xpath.Expr) []*html.Node

	return nil, nil
}

//*[@id="items"]/li[1]/article/div/h1[2]/a/span
//*[@id="items"]/*/article/div/h1[2]/a/span
//a/span
