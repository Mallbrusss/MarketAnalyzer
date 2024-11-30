package models

import "time"

type PlacementPrice struct {
	Figi                  string            `json:"figi"`
	Ticker                string            `json:"ticker"`
	ClassCode             string            `json:"classCode"`
	Isin                  string            `json:"isin"`
	Lot                   int               `json:"lot"`
	Currency              string            `json:"currency"`
	ShortEnabledFlag      bool              `json:"shortEnabledFlag"`
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

type Instrument struct {
	InstrumentID      string `gorm:"primaryKey;size:255"`
	Figi              string `gorm:"index;size:255;not null"`
	Ticker            string `gorm:"index;size:255;not null"`
	ClassCode         string `gorm:"index;size:255;not null"`
	Name              string `gorm:"size:255;not null"`
	Type              string `gorm:"size:50;not null"`
	Currency          string `gorm:"size:10;not null"`
	CountryOfRisk     string `gorm:"size:255;not null"`
	CountryOfRiskName string `gorm:"size:255;not null"`
	Sector            string `gorm:"size:50;not null"`
	CreatedAt         time.Time
}

type Instruments struct{
	Instruments []PlacementPrice `json:"instruments"`
}

type InstrumentPart struct {
	MessageID string `json:"message_id"`
	Part      int    `json:"part"`
	Total     int    `json:"total"`
	Data      []byte `json:"data"`
}