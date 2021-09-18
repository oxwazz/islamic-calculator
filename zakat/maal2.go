package zakat

type Gold struct {
	TotalInGram  int
	OneGramPrice float64
}

func (g *Gold) ZakatPayable() float64 {
	nishab := g.OneGramPrice * 85
	wajibZakat := float64(g.TotalInGram) >= nishab
	var zakat float64 = 0
	if wajibZakat {
		zakat = float64(g.TotalInGram) * 2.5 / 100
	}
	return zakat
}
