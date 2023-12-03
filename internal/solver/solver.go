package solver

type Solver interface {
	Solve() (string, error)
	Name() string
}
