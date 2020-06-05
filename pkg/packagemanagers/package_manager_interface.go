package packagemanagers

// Package Managers
type packageManager interface {
	install(string) error
}
