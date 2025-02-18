package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	taxIncludedPrices map[string]float64
}

func (job TaxIncludedPriceJob) Process() {
	results := make(map[string]float64)
	for _, price := range job.InputPrices {
		results[fmt.Sprintf("%.1f", price)] = price * (1 + job.TaxRate)
	}
	fmt.Println(results)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
