package zakat

type Mal struct {
	TotalAssets        float64
	OneGramGoldInPrice float64
}

func (v *Mal) GetTotalZakat() float64 {
	nishab := v.OneGramGoldInPrice * 85
	wajibZakat := v.TotalAssets >= nishab
	var zakat float64 = 0
	if wajibZakat {
		zakat = v.TotalAssets * 2.5 / 100
	}
	return zakat
}

func (value *Mal) Calc() Result {
	nishab := value.OneGramGoldInPrice * 85
	wajibZakat := value.TotalAssets >= nishab
	var zakat float64 = 0
	if wajibZakat {
		zakat = value.TotalAssets * 2.5 / 100
	}

	return Result{
		Nishab:            nishab,
		IsHaveToPay:       wajibZakat,
		TotalZakatPayable: zakat,
	}
}

type Rikaz struct {
	TotalAssets float64
}

func (v *Rikaz) GetTotalZakat() float64 {
	return v.TotalAssets * 20 / 100
}

func (val *Rikaz) Calc() Result {
	return Result{
		TotalZakatPayable: val.TotalAssets * 20 / 100,
		IsHaveToPay:       true,
		Nishab:            0,
	}
}

type Agriculture struct {
	TotalAssets float64
}
