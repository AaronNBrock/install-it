package packagemanagers

// Package Managers
type PackageManager interface {
	Install(string) error
}
