package linkfinder

import (
	"fmt"
	"io"

	"github.com/antchfx/htmlquery"
)

var (
	xpathRules = []string{
		`//*[@id="Form1"]/main/ul/li/section/div/div/span/a[@href]`,
	}
)

type ResidentAdvisorFinder struct {
	rootUrl string
}

func (x *ResidentAdvisorFinder) FindLinks(html io.Reader) ([]string, error) {

	node, err := htmlquery.Parse(html)

	if err != nil {
		return nil, err
	}

	for _, rule := range xpathRules {

		nodes, err := htmlquery.QueryAll(node, rule)

		if err != nil {
			return nil, err
		}

		for _, node := range nodes {
			// fmt.Printf("%T\n %v\n %#v\n %q\n", node, node, node, node)
			for _, attribute := range node.Attr {
				fmt.Printf("%s%s\n", x.rootUrl, attribute.Val)
			}
		}

		_, _ = node, err

	}

	return nil, nil
}

func NewResidentAdvisorFinder(rootUrl string) *ResidentAdvisorFinder {

	return &ResidentAdvisorFinder{
		rootUrl: rootUrl,
	}
}

// type Linkfinder interface {
// 	FindLinks(html io.Reader) ([]string, error)
// }
