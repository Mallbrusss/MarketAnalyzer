package price_analysis

import (
	"fmt"
)

type Rsi struct{}

func NewRsi() *Rsi {
	return &Rsi{}
}

func (r *Rsi) CalculateRSI(prices []float64, period int) ([]float64, error) {
	if len(prices) == 0 {
		return nil, fmt.Errorf("price data is empty")
	}
	if len(prices) < period+1 {
		return nil, fmt.Errorf("not enough data to calculate RSI with period %d", period)
	}

	return nil, nil
}
