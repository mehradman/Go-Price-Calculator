package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
	IOManager         filemanager.FileManager `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	results := make(map[string]string)

	for _, price := range job.InputPrices {
		taxedPrice := price * (1 + job.TaxRate)
		results[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxedPrice)
	}

	job.TaxIncludedPrices = results

	err := job.IOManager.WriteResult(job)

	if err != nil {
		println(err)
	}
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		IOManager: fm,
	}
}
