package repository

type Interface interface {
	WriteDevice(*Device) error
	WriteSend(*Send) error
	GetSendByFileID(string, string) (*Send, error)
	GetDevice(string, string) (*Device, error)
	UpdateSendState(userID string, deviceID string, syncID string, state State) error
	ReadDevices(string) ([]*DeviceQueryResult, error)
	DeleteDevice(*Device) error
}
