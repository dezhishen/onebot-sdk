package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/dezhishen/onebot-sdk/pkg/config"
)

type HttpCli struct {
	cli         http.Client
	endpoint    string
	accessToken string
}

type accessTokenTransport struct {
	http.RoundTripper
	accessToken string
}

func (ct *accessTokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ct.accessToken))
	return ct.RoundTripper.RoundTrip(req)
}

func NewHttpCli(conf *config.OnebotConfig) *HttpCli {
	if strings.HasSuffix("/", conf.Endpoint) {
		conf.Endpoint = strings.TrimSuffix(conf.Endpoint, "/")
	}
	return &HttpCli{
		cli: http.Client{
			Transport: &accessTokenTransport{http.DefaultTransport, conf.AccessToken},
		},
		endpoint:    conf.Endpoint,
		accessToken: conf.AccessToken,
	}
}

func concatUrl(basePath, url string) string {
	return basePath + "/" + url
}

func (c *HttpCli) PostWithRequsetForResult(url string, req, result interface{}) error {
	requestBody, _ := json.Marshal(req)
	resp, err := c.cli.Post(
		concatUrl(c.endpoint, url),
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	json.Unmarshal(respBodyContent, &result)
	return nil
}

func (c *HttpCli) PostForResult(url string, result interface{}) error {
	resp, err := c.cli.Post(
		concatUrl(c.endpoint, url),
		"application/json",
		bytes.NewBuffer([]byte("{}")),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBodyContent, _ := io.ReadAll(resp.Body)
	json.Unmarshal(respBodyContent, &result)
	return nil
}

func (c *HttpCli) PostWithRequest(url string, req interface{}) error {
	requestBody, _ := json.Marshal(req)
	_, err := c.cli.Post(
		concatUrl(c.endpoint, url),
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func (c *HttpCli) Post(url string) error {
	_, err := c.cli.Post(
		concatUrl(c.endpoint, url),
		"application/json",
		bytes.NewBuffer([]byte("{}")),
	)
	return err
}
