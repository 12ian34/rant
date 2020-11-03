package httpc

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Options struct {
	Timeout int
}

func FetchUrl(method, url string, options Options) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code = %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
