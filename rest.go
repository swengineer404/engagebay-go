package engagebay

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
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

func (c *restClient) do(path, method string, payload, out any) error {
	var r io.Reader
	if payload != nil {
		b, _ := json.Marshal(payload)
		r = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, c.baseURL+path, r)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 && res.StatusCode <= 599 {
		var apierr APIError

		if err := json.NewDecoder(res.Body).Decode(&apierr); err != nil {
			// Assume raw text error, because the developers at engagebay are fucking idiots.
			var idiotErr *json.SyntaxError
			if !errors.As(err, &idiotErr) {
				return err
			}

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return err
			}

			apierr.Message = string(b)
		}

		apierr.Code = res.StatusCode

		return &apierr
	}

	return json.NewDecoder(res.Body).Decode(out)
}
