package shellexecutors

// Shells
type shellexecutors interface {
	exec(name string, arg ...string) (string, string, error)
}
