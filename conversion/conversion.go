package conversion

import (
	"errors"
	"strconv"
)

func StringtoaFloat(strings []string) ([]float64, error) {
	var floats []float64
	for _, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("Failed to covert")
		}
		floats = append(floats, floatVal)
	}
	return floats, nil
}
