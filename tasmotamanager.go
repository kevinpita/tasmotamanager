package tasmotamanager

type TasmotaDevice interface {
	SendCommand(string) error
}
