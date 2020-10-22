package base

const (
	NumberPoint byte = 46
	NegativeFlag byte = 45
)

type BigData interface {
	String() string
	IsNegative() bool
}
