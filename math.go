package math

// Calc is a function that return an operation function.
// The available arguments are "+", "-", "*", "/".
// If an unknown operation, then the addition operation.
func Calc(operation string) func(a, b float64) float64 {
	switch operation {
	case "+":
		return func(a, b float64) float64 {
			return a + b
		}
	case "-":
		return func(a, b float64) float64 {
			return a - b
		}
	case "*":
		return func(a, b float64) float64 {
			return a * b
		}
	case "/":
		return func(a, b float64) float64 {
			return a / b
		}
	default:
		return func(a, b float64) float64 {
			return a + b
		}
	}
}
