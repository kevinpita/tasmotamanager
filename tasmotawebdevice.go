package tasmotamanager

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type TasmotaWebDevice struct {
	Url *url.URL
}

func NewWebDevice(deviceUrl string, username string, password string) (*TasmotaWebDevice, error) {
	url, errParse := url.Parse(deviceUrl)
	if errParse != nil {
		return nil, fmt.Errorf("newwebdevice parse: %w", errParse)

	}
	if !strings.HasSuffix(deviceUrl, "/") {
		url.Path += "/"
	}
	url.Path += "cm"

	if len(username) != 0 && len(password) != 0 {
		q := url.Query()
		q.Set("user", username)
		q.Set("password", password)
		url.RawQuery = q.Encode()

	}

	return &TasmotaWebDevice{url}, nil
}

func (t *TasmotaWebDevice) SendCommand(c string) (map[string]string, error) {
	q := t.Url.Query()
	q.Set("cmnd", c)
	t.Url.RawQuery = q.Encode()

	resp, errReq := http.Get(t.Url.String())
	if errReq != nil {
		return nil, fmt.Errorf("sendcommand get: %w", errReq)
	}
	defer resp.Body.Close()

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return nil, fmt.Errorf("sendcommand readall: %w", errRead)
	}

	var result map[string]string
	errJson := json.Unmarshal(body, &result)
	if errJson != nil {
		return nil, fmt.Errorf("sendcommand unmarshal: %w", errJson)
	}

	return result, nil
}
