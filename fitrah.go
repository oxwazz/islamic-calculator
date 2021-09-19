package islamic_calculator

// islamiccalc

type Fitrah struct {
	OneShaInPrice float32 // 1 sha' = 4 muds = 2.6 ~ 3.0kg of staple food
	TotalPerson   int
}

func (val *Fitrah) GetTotalZakat() float64 {
	return float64(val.OneShaInPrice) * float64(val.TotalPerson)
}
