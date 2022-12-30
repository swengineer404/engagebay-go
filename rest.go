package engagebay

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
		default:
			return fmt.Errorf("invalid content type: %s", contentType)
		}

	}

	req, err := http.NewRequest(method, c.baseURL+path, r)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", c.key)

	res, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 && res.StatusCode <= 599 {
		var apierr APIError

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(b, out); err != nil {
			var fallbackErr *json.SyntaxError
			if !errors.As(err, &fallbackErr) {
				return err
			}

			apierr.Message = string(b)
		}

		apierr.Code = res.StatusCode

		return &apierr
	}

	return json.NewDecoder(res.Body).Decode(out)
}
