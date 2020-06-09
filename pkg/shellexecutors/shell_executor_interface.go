package shellexecutors

// Shells
type ShellExecutor interface {
	Exec(name string, arg ...string) (string, string, error)
}
