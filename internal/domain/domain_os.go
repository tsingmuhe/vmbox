package domain

type DomainOS struct {
	Loader *DomainLoader
	NVRam  *DomainNVRam
	DomainVCPU
}

type DomainLoader struct {
	Path      string // /path/to/OVMF_CODE.fd
	Readonly  string // true
	Stateless string // yes/no
	Type      string // pflash
	Format    string // raw
}

type DomainNVRam struct {
	Path   string // /path/to/OVMF_VARS.fd
	Type   string // pflash
	Format string // raw
}
