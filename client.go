package main

import (
	"fmt"
	"io"
	"net/http"
)

var remoteUrl = "https://api.github.com/repos/tldr-pages/tldr/contents/pages"

func getTldr(platform string, cmd string) (io.Reader, error) {
	url := fmt.Sprintf("%s/%s/%s.md", remoteUrl, platform, cmd)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.raw")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	//content, err := ioutil.ReadAll(resp.Body)

	return resp.Body, err
}
