package device

import (
	"fmt"

	"github.com/kevinpita/tasmotamanager"
)

type LightBulbRgb struct {
	tasmotamanager.TasmotaDevice
}

func NewLightBulbRgb(d tasmotamanager.TasmotaDevice) *LightBulbRgb {
	return &LightBulbRgb{d}
}

// On TODO: check response && log errors
func (lb *LightBulbRgb) On() bool {
	_, errCommand := lb.SendCommand("Power on")
	if errCommand != nil {
		return false
	}

	return true
}

// Off TODO: check response && log errors
func (lb *LightBulbRgb) Off() bool {
	_, errCommand := lb.SendCommand("Power off")
	if errCommand != nil {
		return false
	}

	return true
}

// Bright TODO: check response && log errors
func (lb *LightBulbRgb) Bright(v int) bool {
	if v > 100 || v < 0 {
		return false
	}
	cmd := fmt.Sprintf("Dimmer %d", v)

	_, errCommand := lb.SendCommand(cmd)
	if errCommand != nil {
		return false
	}

	return true
}

// ColorTemperature TODO: check response && log errors
func (lb *LightBulbRgb) ColorTemperature(v int) bool {
	if v > 500 || v < 153 {
		return false
	}
	cmd := fmt.Sprintf("CT %d", v)

	_, errCommand := lb.SendCommand(cmd)
	if errCommand != nil {
		return false
	}

	return true
}

// Saturation TODO: check response && log errors
func (lb *LightBulbRgb) Saturation(v int) bool {
	if v > 100 || v < 0 {
		return false
	}
	cmd := fmt.Sprintf("HSBColor2 %d", v)

	_, errCommand := lb.SendCommand(cmd)
	if errCommand != nil {
		return false
	}

	return true
}

// Hue TODO: check response && log errors
func (lb *LightBulbRgb) Hue(v int) bool {
	if v > 360 || v < 0 {
		return false
	}
	cmd := fmt.Sprintf("HSBColor1 %d", v)

	_, errCommand := lb.SendCommand(cmd)
	if errCommand != nil {
		return false
	}

	return true
}
