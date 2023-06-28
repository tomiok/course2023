package add

func Adding(a, b int) int {
	return a + b
}

type Multiplier interface {
	Multiply(a, b int) int
}

type Calculator struct{}

func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

type Service struct {
	M Multiplier
}

func (s *Service) Foo(a, b int) int {
	result := s.M.Multiply(a, b)

	return 10 + result
	// if result = 10
	// this func returns 20
}
