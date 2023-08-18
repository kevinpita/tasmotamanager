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
	Url string
}

func NewWebDevice(deviceUrl string, username string, password string) *TasmotaWebDevice {
	var baseUrl string
	if strings.HasSuffix(deviceUrl, "/") {
		baseUrl = fmt.Sprintf("%scm?", deviceUrl)
	} else {
		baseUrl = fmt.Sprintf("%s/cm?", deviceUrl)
	}

	if len(username) != 0 && len(password) != 0 {
		auth := fmt.Sprintf("user=%s&password=%s", username, password)
		baseUrl = fmt.Sprintf("%s%s&", baseUrl, url.QueryEscape(auth))
	}

	return &TasmotaWebDevice{baseUrl}
}

func (t *TasmotaWebDevice) SendCommand(c string) (map[string]string, error) {
	urlRequest := fmt.Sprintf("%scmnd=%s", t.Url, url.QueryEscape(c))
	resp, errReq := http.Get(urlRequest)
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
