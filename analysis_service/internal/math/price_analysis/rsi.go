package price_analysis

import (
	"fmt"
	"math"
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

	rsiValues := make([]float64, len(prices))
	gains := make([]float64, len(prices)-1)
	losses := make([]float64, len(prices)-1)

	// Шаг 1: Расчет изменений цен
	for i := 1; i < len(prices); i++ {
		change := prices[i] - prices[i-1]
		if change > 0 {
			gains[i-1] = change
			losses[i-1] = 0
		} else {
			gains[i-1] = 0
			losses[i-1] = math.Abs(change)
		}
	}

	// Шаг 2: Расчет начальных средних значений приростов и падений
	sumGains := 0.0
	sumLosses := 0.0
	for i := 0; i < period; i++ {
		sumGains += gains[i]
		sumLosses += losses[i]
	}
	avgGain := sumGains / float64(period)
	avgLoss := sumLosses / float64(period)

	// Шаг 3: Расчет RSI для каждого элемента
	for i := period; i < len(prices); i++ {
		if avgLoss == 0 {
			rsiValues[i] = 100 // Если нет потерь, RSI = 100
		} else {
			rs := avgGain / avgLoss
			rsiValues[i] = 100 - (100 / (1 + rs))
		}

		// Обновление средних значений приростов и падений
		if i < len(prices)-1 {
			avgGain = (avgGain*(float64(period)-1) + gains[i]) / float64(period)
			avgLoss = (avgLoss*(float64(period)-1) + losses[i]) / float64(period)
		}
	}

	// Установка NaN для первых period-1 значений
	for i := 0; i < period; i++ {
		rsiValues[i] = math.NaN()
	}

	return rsiValues, nil
}

func (r *Rsi) GetLastRSI(prices []float64, period int) (float64, error) {
	rsiValues, err := r.CalculateRSI(prices, period)
	if err != nil {
		return 0, err
	}

	lastIndex := len(rsiValues) - 1
	if lastIndex >= 0 && !math.IsNaN(rsiValues[lastIndex]) {
		return rsiValues[lastIndex], nil
	}
	return 0, fmt.Errorf("couldn't calculate the last value of the RSI")
}
