package semantic

// Version struct holds semantic versioning information
type Version struct {
	Major int
	Minor int
	Patch int
}

// VersionFromString creates new semantic Version by parsing version string
func VersionFromString(version string) (Version, error) {
	_ = "STUB: not implemented"
	return *new(Version), nil
}

// Lt returns true if this version is less than version parameter
func (v Version) Lt(version Version) bool { _ = "STUB: not implemented"; return false }
