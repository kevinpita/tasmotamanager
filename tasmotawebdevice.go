package tasmotamanager

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type TasmotaWebDevice struct {
	URL *url.URL
}

func NewWebDevice(deviceURL string, username string, password string) (*TasmotaWebDevice, error) {
	u, errParse := url.Parse(deviceURL)
	if errParse != nil {
		return nil, fmt.Errorf("newwebdevice parse: %w", errParse)
	}

	if !strings.HasSuffix(deviceURL, "/") {
		u.Path += "/"
	}
	u.Path += "cm"

	if len(username) != 0 && len(password) != 0 {
		q := u.Query()
		q.Set("user", username)
		q.Set("password", password)
		u.RawQuery = q.Encode()
	}

	return &TasmotaWebDevice{u}, nil
}

func (t *TasmotaWebDevice) SendCommand(c string) (map[string]string, error) {
	const timeout = 5
	t.PrepareCommandURL(c)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*timeout)
	defer cancel()

	req, errCtx := http.NewRequestWithContext(ctx, http.MethodGet, t.URL.String(), nil)
	if errCtx != nil {
		// Handle error
		return nil, errCtx
	}

	resp, errReq := http.DefaultClient.Do(req)
	if errReq != nil {
		// Handle error
		return nil, errReq
	}
	defer resp.Body.Close()

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return nil, fmt.Errorf("sendcommand readall: %w", errRead)
	}

	var result map[string]string
	errJSON := json.Unmarshal(body, &result)
	if errJSON != nil {
		return nil, fmt.Errorf("sendcommand unmarshal: %w", errJSON)
	}

	return result, nil
}

func (t *TasmotaWebDevice) PrepareCommandURL(c string) {
	q := t.URL.Query()
	q.Set("cmnd", c)
	t.URL.RawQuery = q.Encode()
}
