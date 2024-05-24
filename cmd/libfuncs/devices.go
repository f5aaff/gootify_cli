package libfuncs

type Devices struct {
	*BasicEndPoint
	EndPointStr string
}

type Device struct {
	*BasicResource
}

type DeviceContent struct {
	ID               string `json:"id"`
	IsActive         bool   `json:"is_active"`
	IsPrivateSession bool   `json:"is_private_session"`
	IsRestricted     bool   `json:"is_restricted"`
	Name             string `json:"name"`
	SupportsVolume   bool   `json:"supports_volume"`
	Type             string `json:"type"`
	VolumePercent    int    `json:"volume_percent"`
}

func (d Devices) GetAll() []Resource {
	return nil
}

func (d Devices) GetItem() Device {

	device := Device{}
	device.Content = DeviceContent{}
	device.Typ = "device"
	GetItem(device, d.GetEndPointStr())
	return device
}
