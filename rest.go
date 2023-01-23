package engagebay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	ContentTypeJson = "application/json"
	ContentTypeForm = "application/x-www-form-urlencoded"
)

type restClient struct {
	c       *http.Client
	baseURL string
	key     string
}

func newAPIClient(key string) *restClient {
	return &restClient{
		c:       &http.Client{Timeout: time.Second * 30},
		baseURL: "https://app.engagebay.com/dev/api/panel",
		key:     key,
	}
}

func (c *restClient) do(path, method, contentType string, payload, out any) error {
	var r io.Reader
	if payload != nil {
		switch contentType {
		case ContentTypeJson:
			b, _ := json.Marshal(payload)
			r = bytes.NewReader(b)
		case ContentTypeForm:
			r = strings.NewReader(payload.(url.Values).Encode())
		default:
			return fmt.Errorf("invalid content type: %s", contentType)
		}

	}

	req, err := http.NewRequest(method, c.baseURL+path, r)
	if err != nil {
		return err
	}

	if r != nil {
		req.Header.Set("Content-Type", contentType)
	}

	req.Header.Set("Authorization", c.key)
	req.Header.Set("Accept", "application/json")

	res, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 && res.StatusCode <= 599 {
		var apierr APIError
		if err := json.Unmarshal(b, out); err != nil {
			apierr.Message = string(b)
		}

		apierr.Code = res.StatusCode

		return &apierr
	}

	if out == nil {
		return nil
	}

	if err := json.Unmarshal(b, out); err != nil {
		return &APIError{
			Message: string(b),
			Code:    res.StatusCode,
		}
	}

	return nil
}
