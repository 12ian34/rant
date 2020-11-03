/*
 check why defer is failing
 test with different methods
 try using timeout, inside the client
*/

package main

import (
	"bytes"
	"crawler/httpc"
	"crawler/linkfinder"
	"flag"
	"fmt"
	"net/http"
	packageUrl "net/url"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var url, method string

func init() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	flag.StringVar(&url, "u", "", "enter the full URL that you'd like to scrape")
	flag.StringVar(&method, "m", http.MethodGet, "enter the desired HTTP method")
	flag.Parse()

}

func main() {

	if len(url) == 0 {
		flag.PrintDefaults()
		return
	}

	body, err := httpc.FetchUrl(method, url, httpc.Options{})

	if err != nil {
		log.Error().Err(err).Str("url", url).Msg("error!")
		return
	}

	// _, err = linkfinder.FindLinks(bytes.NewReader(body), "//a[starts-with(@href, 'https://')]")

	parsedUrl, err := packageUrl.Parse(url)

	if err != nil {
		log.Error().Err(err).Str("url", url).Msg("error!")
		return
	}

	rootUrl := fmt.Sprintf("%s://%s", parsedUrl.Scheme, parsedUrl.Host)

	// fmt.Printf("%s=%s\n", attribute.Key, attribute.Val)

	linkfinder.NewResidentAdvisorFinder(rootUrl).FindLinks(bytes.NewReader(body))

	// find the proper xpath query
	// fmt.Println(string(body))

	// url.Scheme+"://"+url.Host

}
