package zakat

type Zakat interface {
	GetTotalZakat() float64
	Calc() Result
}
type Zakat2 interface {
	GetZakatPayable() float64
}
