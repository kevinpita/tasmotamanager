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

// On TODO: check response && log errors.
func (lb *LightBulbRgb) On() bool {
	_, errCommand := lb.SendCommand("Power on")
	return errCommand == nil
}

// Off TODO: check response && log errors.
func (lb *LightBulbRgb) Off() bool {
	_, errCommand := lb.SendCommand("Power off")
	return errCommand == nil
}

// Toggle TODO: check response && log errors.
func (lb *LightBulbRgb) Toggle() bool {
	_, errCommand := lb.SendCommand("Power toggle")
	return errCommand == nil
}

// Blink TODO: check response && log errors.
func (lb *LightBulbRgb) Blink() bool {
	_, errCommand := lb.SendCommand("Power blink")
	return errCommand == nil
}

// StopBlink TODO: check response && log errors.
func (lb *LightBulbRgb) StopBlink() bool {
	_, errCommand := lb.SendCommand("Power blinkoff")
	return errCommand == nil
}

// Bright TODO: check response && log errors.
func (lb *LightBulbRgb) Bright(v int) bool {
	if v > 100 || v < 0 {
		return false
	}
	cmd := fmt.Sprintf("Dimmer %d", v)

	_, errCommand := lb.SendCommand(cmd)
	return errCommand == nil
}

// ColorTemperature TODO: check response && log errors.
func (lb *LightBulbRgb) ColorTemperature(v int) bool {
	if v > 500 || v < 153 {
		return false
	}
	cmd := fmt.Sprintf("CT %d", v)

	_, errCommand := lb.SendCommand(cmd)
	return errCommand == nil
}

// Saturation TODO: check response && log errors.
func (lb *LightBulbRgb) Saturation(v int) bool {
	if v > 100 || v < 0 {
		return false
	}
	cmd := fmt.Sprintf("HSBColor2 %d", v)

	_, errCommand := lb.SendCommand(cmd)
	return errCommand == nil
}

// Hue TODO: check response && log errors.
func (lb *LightBulbRgb) Hue(v int) bool {
	if v > 360 || v < 0 {
		return false
	}
	cmd := fmt.Sprintf("HSBColor1 %d", v)

	_, errCommand := lb.SendCommand(cmd)
	return errCommand == nil
}
