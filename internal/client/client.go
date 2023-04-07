package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	URI string
}

func New(serverURI string) Client {
	return Client{URI: serverURI}
}

func (c Client) Generate(length uint) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/generate?n=%d", c.URI, length))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	strBody := string(body)
	strings.TrimRight(strBody, "\n")
	return string(body), err
}
