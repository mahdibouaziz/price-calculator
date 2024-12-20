package prices

import (
	"fmt"

	"github.com/mahdibouaziz/price-calculator/conversion"
	"github.com/mahdibouaziz/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOmanager         iomanager.IOManager `json:"-"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{},
		IOmanager:   iom,
	}
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOmanager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) error {
	err := job.LoadData()
	if err != nil {
		errorChan <- err
		return err
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	err = job.IOmanager.WriteResult(job)
	if err != nil {
		errorChan <- err
		return err
	}

	doneChan <- true
	return nil
}
