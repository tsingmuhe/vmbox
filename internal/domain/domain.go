package domain

type Domain struct {
	Type      string    `yaml:"type"` //kvm、hvf、tcg
	OS        *DomainOS `yaml:"os"`
	VCPU      uint
	IOThreads uint
	Memory    *DomainMemory
	Devices   *DomainDeviceList `xml:"devices"`
}
