package main

import (
	"fmt"

	"github.com/mahdibouaziz/price-calculator/filemanager"
	"github.com/mahdibouaziz/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// cmd := cmdmanager.New()
		// priceJob := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", (taxRate*100)))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println(err)
		}

	}

}
