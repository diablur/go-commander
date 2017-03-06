package commander

type Version struct {
	version string
}

func (v *Version) Set(ver string) *Version {
	v.version = ver
	return v
}

func (v Version) Get() string {
	return v.version
}

func (v Version) Valid() bool {
	return len(v.version) != 0
}