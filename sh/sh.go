package sh

// Checker interface
type Checker interface {
	Check(file string) error
	Install() error
}

// All checkers
func All() []Checker {
	return []Checker{
		&shellcheck{},
		&shfmt{},
	}
}
