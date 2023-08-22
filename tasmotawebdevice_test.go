package tasmotamanager_test

import (
	"testing"

	. "github.com/kevinpita/tasmotamanager"
)

// TODO: Tests should have a mock server that could give a json response so it can be compared too

func TestNoLoginDeviceBackslash(t *testing.T) {
	const expected = "http://192.168.1.2/cm?cmnd=Power+off"
	device, errDevice := NewWebDevice("http://192.168.1.2/", "", "")
	if errDevice != nil {
		t.Errorf("no error was expected when creating the device -> %v", errDevice)
	}

	device.PrepareCommandURL("Power off")
	actual := device.URL.String()

	if actual != expected {
		t.Errorf("got %s expected %s", actual, expected)
	}
}

func TestNoLoginDeviceNoBackslash(t *testing.T) {
	const expected = "http://192.168.1.2/cm?cmnd=Power+off"
	device, errDevice := NewWebDevice("http://192.168.1.2", "", "")
	if errDevice != nil {
		t.Errorf("no error was expected when creating the device -> %v", errDevice)
	}

	device.PrepareCommandURL("Power off")
	actual := device.URL.String()

	if actual != expected {
		t.Errorf("got %s expected %s", actual, expected)
	}
}

func TestLoginDeviceBackslash(t *testing.T) {
	const expected = "http://192.168.1.2/cm?cmnd=Power+off&password=pass&user=admin"
	device, errDevice := NewWebDevice("http://192.168.1.2/", "admin", "pass")
	if errDevice != nil {
		t.Errorf("no error was expected when creating the device -> %v", errDevice)
	}

	device.PrepareCommandURL("Power off")
	actual := device.URL.String()

	if actual != expected {
		t.Errorf("got %s expected %s", actual, expected)
	}
}

func TestLoginDeviceNoBackslash(t *testing.T) {
	const expected = "http://192.168.1.2/cm?cmnd=Power+off&password=pass&user=admin"
	device, errDevice := NewWebDevice("http://192.168.1.2", "admin", "pass")
	if errDevice != nil {
		t.Errorf("no error was expected when creating the device -> %v", errDevice)
	}

	device.PrepareCommandURL("Power off")
	actual := device.URL.String()

	if actual != expected {
		t.Errorf("got %s expected %s", actual, expected)
	}
}
