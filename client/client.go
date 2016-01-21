package client

import (
	"fmt"
	"io/ioutil"
	// "log"
	"net/http"
	"time"
)

var remoteUrl = "https://raw.githubusercontent.com/tldr-pages/tldr/master/pages"

type Client struct {
	client *http.Client
}

func New() *Client {
	return &Client{client: &http.Client{Timeout: 60 * time.Second}}
}

func (c *Client) do(req *http.Request) ([]byte, string, error) {
	req.Header.Set("Accept", "application/vnd.github.v3.raw")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return nil, "", fmt.Errorf("Bad response status %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)

	return data, resp.Header.Get("ETag"), err
}

func (c *Client) GetTldr(platform string, cmd string) (string, string, error) {
	url := fmt.Sprintf("%s/%s/%s.md", remoteUrl, platform, cmd)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", "", err
	}

	data, etag, err := c.do(req)

	return string(data), etag, err
}

func (c *Client) IsExpired(platform string, cmd string, etag string) (bool, error) {
	url := fmt.Sprintf("%s/%s/%s.md", remoteUrl, platform, cmd)
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Add("If-None-Match", etag)

	resp, err := c.client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 304 {
		return false, nil
	}

	return true, nil
}

func (c *Client) GetIndex() (string, error) {
	url := fmt.Sprintf("http://tldr-pages.github.io/assets/index.json")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	data, _, err := c.do(req)

	return string(data), err
}
