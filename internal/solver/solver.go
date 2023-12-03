package solver

type Solver interface {
	Solve() error
	Name() string
}
