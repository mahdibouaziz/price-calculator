package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloat(data []string) ([]float64, error) {
	result := []float64{}
	for _, value := range data {
		floatVal, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println("failed to convert string to float")
			return nil, err
		}
		result = append(result, floatVal)
	}
	return result, nil
}
