package domain

type DomainDeviceList struct {
	Disks      []DomainDisk
	Interfaces []DomainInterface
}

type DomainDeviceBoot struct {
	Order uint `xml:"order,attr"`
}

type DomainDisk struct {
	Device string //disk cdrom

	Boot *DomainDeviceBoot
}

type DomainInterface struct {
	Boot *DomainDeviceBoot
}
