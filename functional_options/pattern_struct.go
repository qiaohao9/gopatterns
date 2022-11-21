package functional_options

type OptionalFunc func(*Option)

type Option struct {
	ParametersA int
	ParametersB int
}

var defaultOption = &Option{
	ParametersA: 1,
	ParametersB: 1,
}

func WithParametersA(a int) OptionalFunc {
	return func(o *Option) {
		o.ParametersA = a
	}
}

func WithParametersB(b int) OptionalFunc {
	return func(o *Option) {
		o.ParametersB = b
	}
}

func evaluate(optionsFunc ...OptionalFunc) Option {
	optionCopy := *defaultOption

	for _, o := range optionsFunc {
		o(&optionCopy)
	}

	return optionCopy
}

func NewSomething(optionsFunc ...OptionalFunc) {
	_ = evaluate(optionsFunc...)
}
