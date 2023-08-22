package tasmotamanager

type TasmotaDevice interface {
	SendCommand(string) (map[string]string, error)
}
