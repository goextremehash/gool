package common

import (
	"golang.org/x/mod/semver"
)

const version = "v0.0.5"

func GetVersion() string {
	return semver.Canonical(version)
}
