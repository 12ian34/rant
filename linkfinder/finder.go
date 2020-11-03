package linkfinder

import "io"

type Linkfinder interface {
	FindLinks(html io.Reader) ([]string, error)
}
