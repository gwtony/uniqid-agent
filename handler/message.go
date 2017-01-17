package handler

type UDPdata struct {
	Uid []byte
	Puid []byte
	Pip []byte
	Pport uint16
	Lip []byte
	Lport uint16
	Dlen uint16
	Data []byte
}
type RouterData struct {
	Magic byte
	Len   uint16
	Data  []byte
}
