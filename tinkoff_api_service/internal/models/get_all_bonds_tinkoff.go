package models

import "time"

// InstrumentStatus: "INSTRUMENT_STATUS_UNSPECIFIED" - Значение не определено.
//
// InstrumentStatus: "INSTRUMENT_STATUS_BASE" - Базовый список инструментов (по умолчанию).
// Инструменты доступные для торговли через TINKOFF INVEST API. Cейчас списки бумаг,
// доступных из api и других интерфейсах совпадают (за исключением внебиржевых бумаг)
// , но в будущем возможны ситуации, когда списки инструментов будут отличаться.
//
// InstrumentStatus: "INSTRUMENT_STATUS_ALL" - Список всех инструментов.
type BondsRequest struct {
	InstrumentStatus string `json:"instrumentStatus"`
}

type GetBondsResponse struct {
	Instruments []PlacementPrice `json:"instruments"`
}

type PlacementPrice struct {
	Figi                  string `json:"figi"`
	Ticker                string `json:"ticker"`
	ClassCode             string `json:"classCode"`
	Isin                  string `json:"isin"`
	Lot                   int    `json:"lot"`
	Currency              string `json:"currency"`
	ShortEnabledFlag      bool
	Name                  string            `json:"name"`
	Exchange              string            `json:"exchange"`
	IssueSize             string            `json:"issueSize"`
	CountryOfRisk         string            `json:"countryOfRisk"`
	CountryOfRiskName     string            `json:"countryOfRiskName"`
	Sector                string            `json:"sector"`
	IssueSizePlan         string            `json:"issueSizePlan"`
	Nominal               Nominal           `json:"nominal"`
	TradingStatus         string            `json:"tradingStatus"`
	OtcFlag               bool              `json:"octFlag"`
	BuyAvailableFlag      bool              `json:"buyAvailableFlag"`
	SellAvailableFlag     bool              `json:"sellAvailableFlag"`
	DivYieldFlag          bool              `json:"divYieldFlag"`
	ShareType             string            `json:"shareType"`
	MinPriceIncrement     MinPriceIncrement `json:"minPriceIncrement"`
	ApiTradeAvailableFlag bool              `json:"apiTradeAvailableFlag"`
	Uid                   string            `json:"uid"`
	RealExchange          string            `json:"realExchange"`
	PositionUid           string            `json:"positionUid"`
	AssetUid              string            `json:"assetUid"`
	InstrumentExchange    string            `json:"instrumentExchange"`
	ForIisFlag            bool              `json:"forIisFlag"`
	ForQualInvestorFlag   bool              `json:"forQualInvestorFlag"`
	WeekendFlag           bool              `json:"weekendFlag"`
	BlockedTcaFlag        bool              `json:"blockedTcaFlag"`
	LiquidityFlag         bool              `json:"liquidityFlag"`
	First1minCandleDate   time.Time         `json:"first1minCandleDate"`
	First1dayCandleDate   time.Time         `json:"first1dayCandleDate"`
	brand                 brand             
}

type Nominal struct {
	Currency string `json:"currency"`
	Units    string `json:"units"`
	Nano     int    `json:"nano"`
}

type MinPriceIncrement struct {
	Units string `json:"units"`
	Nano  int    `json:"nano"`
}

type brand struct {
	logoName      string
	logoBaseColor string
	textColor     string
}
