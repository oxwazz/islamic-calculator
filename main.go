package main

import (
	"fmt"

	"github.com/oxwazz/islamic-calculator/zakat"
)


func HitungZakat(z zakat.Zakat)  {
	fmt.Println(z.GetTotalZakat())
}
func HitungZakats(z zakat.Zakat)  {
	fmt.Println(z.GetTotalZakat())
}
func main() {
	mal := &zakat.Mal{
		TotalAssets     :1000,
		OneGramGoldInPrice :234,
	}

	rikaz := &zakat.Rikaz{
		TotalAssets     :1000,
	}

	// _ := &zakat.Fitrah{
	// 	OneShaInPrice: 1000,
	// 	TotalPerson: 3,
	// }

	fmt.Println("halo")
	HitungZakat(mal)
	HitungZakats(rikaz)

	// mal := &IslamicCalc.zakat.Mal{
	// 	TotalAssets     :1000,
	// 	OneGramGoldInPrice :234,
	// }

	// IslamicCalc.Zakat(mal)
	// IslamicCalc.Faraid(mal)
}