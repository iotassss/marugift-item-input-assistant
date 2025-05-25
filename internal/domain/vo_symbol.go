package domain

type Symbol struct {
	value string
}

func NewSymbol(value string) (Symbol, error) {
	return Symbol{value: value}, nil
}

func (s Symbol) Value() string {
	return s.value
}
