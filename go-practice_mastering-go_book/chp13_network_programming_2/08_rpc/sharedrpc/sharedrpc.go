package sharedrpc

// CalcInput definescalc inputs
type CalcInput struct {
	A1, A2 float64
}

// MathExecer is math operations interface
type MathExecer interface {
	Multiply(args *CalcInput, result *float64) error
	Power(args *CalcInput, result *float64) error
}
