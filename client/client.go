package client

import (
	"fmt"
	"io/ioutil"
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

func (c *Client) do(req *http.Request) ([]byte, error) {
	req.Header.Set("Accept", "application/vnd.github.v3.raw")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return nil, fmt.Errorf("Bad response status %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

func (c *Client) GetTldr(platform string, cmd string) (string, error) {
	url := fmt.Sprintf("%s/%s/%s.md", remoteUrl, platform, cmd)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	data, err := c.do(req)

	return string(data), err
}

func (c *Client) GetIndex() (string, error) {
	url := fmt.Sprintf("%s/index.json", remoteUrl)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	data, err := c.do(req)

	return string(data), err
}
