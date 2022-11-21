package functional_options

type FuncOption interface {
	apply(*option)
}

type funcOption struct {
	f func(*option)
}

func (fdo *funcOption) apply(do *option) {
	fdo.f(do)
}

func newFunOption(f func(*option)) *funcOption {
	return &funcOption{f: f}
}

var defaultFuncOption = &option{
	parameterC: 1,
	parameterD: 1,
}

type option struct {
	parameterC int
	parameterD int
}

func WithParameterC(c int) FuncOption {
	return newFunOption(func(o *option) {
		o.parameterC = c
	})
}

func WithParameterD(d int) FuncOption {
	return newFunOption(func(o *option) {
		o.parameterD = d
	})
}

func evaluateFuncOptions(funcOptions ...FuncOption) *option {
	optionCopy := *defaultFuncOption

	for _, funcOption := range funcOptions {
		funcOption.apply(&optionCopy)
	}

	return &optionCopy
}

func NewFunc(funcOptions ...FuncOption) {
	_ = evaluateFuncOptions(funcOptions...)
}
