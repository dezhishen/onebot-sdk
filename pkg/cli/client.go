package cli

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/dezhishen/onebot-sdk/pkg/config"
)

func PostWithRequsetForResult(url string, req, result interface{}) error {
	requestBody, _ := json.Marshal(req)
	resp, err := http.Post(
		config.GetHttpUrl()+url,
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

func PostForResult(url string, result interface{}) error {
	resp, err := http.Post(
		config.GetHttpUrl()+url,
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

func PostWithRequest(url string, req interface{}) error {
	requestBody, _ := json.Marshal(req)
	_, err := http.Post(
		config.GetHttpUrl()+url,
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	return err
}

func Post(url string) error {
	_, err := http.Post(
		config.GetHttpUrl()+url,
		"application/json",
		bytes.NewBuffer([]byte("{}")),
	)
	return err
}
