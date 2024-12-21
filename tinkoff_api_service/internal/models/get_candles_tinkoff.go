package models

import "time"

type GetCandlesRequest struct {
	Figi string `json:"figi"`
	From string `json:"from"`
	To   string `json:"to"`
	// CANDLE_INTERVAL_UNSPECIFIED Интервал не определён.
	//
	// CANDLE_INTERVAL_1_MIN от 1 минуты до 1 дня.
	//
	// CANDLE_INTERVAL_5_MIN от 5 минут до 1 дня.
	//
	// CANDLE_INTERVAL_15_MIN от 15 минут до 1 дня.
	//
	// CANDLE_INTERVAL_HOUR от 1 часа до 1 недели.
	//
	// CANDLE_INTERVAL_DAY от 1 дня до 1 года.
	//
	// CANDLE_INTERVAL_2_MIN от 2 минут до 1 дня.
	//
	// CANDLE_INTERVAL_3_MIN от 3 минут до 1 дня.
	//
	// CANDLE_INTERVAL_10_MIN от 10 минут до 1 дня.
	//
	// CANDLE_INTERVAL_30_MIN от 30 минут до 2 дней.
	//
	// CANDLE_INTERVAL_2_HOUR от 2 часов до 1 месяца.
	//
	// CANDLE_INTERVAL_4_HOUR от 4 часов до 1 месяца.
	//
	// CANDLE_INTERVAL_WEEK от 1 недели до 2 лет.
	//
	// CANDLE_INTERVAL_MONTH 	от 1 месяца до 10 лет.
	Interval     string `json:"interval"`
	InstrumentId string `json:"instrumentId"`
}

type HistoricCandle struct {
	InstrumentId string    `json:"instrumentID"`
	Volume       string    `json:"volume"`
	High         High      `json:"high"`
	Low          Low       `json:"low"`
	Time         time.Time `json:"time"`
	Close        Close     `json:"close"`
	Open         Open      `json:"open"`
}

type High struct {
	Units string `json:"units"`
	Nano  int    `json:"nano"`
}

type Low struct {
	Units string `json:"units"`
	Nano  int    `json:"nano"`
}

type Close struct {
	Units string `json:"units"`
	Nano  int    `json:"nano"`
}

type Open struct {
	Units string `json:"units"`
	Nano  int    `json:"nano"`
}

type GetCandlesResponse struct {
	Candles      []HistoricCandle `json:"candles"`
}
