package models

import "time"

type HistoricCandle struct {
	InstrumentId string    `json:"instrumentID"`
	Time         time.Time `json:"time"`
	Volume       string    `json:"volume"`
	High         High      `json:"high"`
	Low          Low       `json:"low"`
	Close        Close     `json:"close"`
	Open         Open      `json:"open"`
}

type High struct {
	InstrumentId string    `json:"-"`
	Time         time.Time `json:"-"`
	Units        string    `json:"units"`
	Nano         int       `json:"nano"`
}

type Low struct {
	InstrumentId string    `json:"-"`
	Time         time.Time `json:"-"`
	Units        string    `json:"units"`
	Nano         int       `json:"nano"`
}

type Close struct {
	InstrumentId string    `json:"-"`
	Time         time.Time `json:"-"`
	Units        string    `json:"units"`
	Nano         int       `json:"nano"`
}

type Open struct {
	InstrumentId string    `json:"-"`
	Time         time.Time `json:"-"`
	Units        string    `json:"units"`
	Nano         int       `json:"nano"`
}

type HistoricCandles struct {
	Candles []HistoricCandle `json:"candles"`
}
