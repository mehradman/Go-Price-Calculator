package prices

import (
	"bufio"
	"fmt"
	"os"

	"example.com/price-calculator/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	taxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {

	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)

	var lines []string

	// scanner reads line by line and returns boolean
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading the file content failed.")
		fmt.Println(err)
		file.Close()
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	job.InputPrices = prices
	file.Close()
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	results := make(map[string]string)
	for _, price := range job.InputPrices {
		taxedPrice := price * (1 + job.TaxRate)
		results[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxedPrice)
	}
	fmt.Println(results)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}
