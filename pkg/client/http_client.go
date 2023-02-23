package client

import (
	"bytes"
	"encoding/json"
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

func NewHttpCli(server *config.OnebotConfig) *HttpCli {
	return &HttpCli{
		cli:         http.Client{},
		endpoint:    server.Endpoint,
		accessToken: server.AccessToken,
	}
}

func concatUrl(basePath, url string) string {
	basePath = strings.TrimSuffix(basePath, "/")
	if !strings.HasPrefix(url, "/") {
		url = "/" + url
	}
	return basePath + url
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
