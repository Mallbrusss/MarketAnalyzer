package models

import (
	"time"
)

type PlacementPrice struct {
	// gorm.Model

	Figi                  string            `json:"figi" gorm:"unique;type:VARCHAR(255)"`        // Уникальное строковое поле
	Ticker                string            `json:"ticker" gorm:"type:VARCHAR(255)"`             // Обычная строка
	ClassCode             string            `json:"classCode" gorm:"type:VARCHAR(255)"`          // Обычная строка
	Isin                  string            `json:"isin" gorm:"type:VARCHAR(255)"`               // Обычная строка
	Lot                   int               `json:"lot" gorm:"type:INT"`                         // Число
	Currency              string            `json:"currency" gorm:"type:VARCHAR(50)"`            // Строка с ограниченной длиной
	ShortEnabledFlag      bool              `json:"shortEnabledFlag" gorm:"type:BOOLEAN"`        // Логическое значение
	Name                  string            `json:"name" gorm:"type:VARCHAR(255)"`               // Обычная строка
	Exchange              string            `json:"exchange" gorm:"type:VARCHAR(100)"`           // Строка с ограниченной длиной
	IssueSize             string            `json:"issueSize" gorm:"type:VARCHAR(100)"`          // Строка для размеров эмиссии
	CountryOfRisk         string            `json:"countryOfRisk" gorm:"type:VARCHAR(100)"`      // Страна риска
	CountryOfRiskName     string            `json:"countryOfRiskName" gorm:"type:VARCHAR(255)"`  // Название страны риска
	Sector                string            `json:"sector" gorm:"type:VARCHAR(255)"`             // Сектор
	IssueSizePlan         string            `json:"issueSizePlan" gorm:"type:VARCHAR(100)"`      // Планируемый объем эмиссии
	Nominal               Nominal           `json:"nominal" gorm:"-"`                            // Внешний ключ
	TradingStatus         string            `json:"tradingStatus" gorm:"type:VARCHAR(100)"`      // Торговый статус
	OtcFlag               bool              `json:"octFlag" gorm:"type:BOOLEAN"`                 // Внебиржевой флаг
	BuyAvailableFlag      bool              `json:"buyAvailableFlag" gorm:"type:BOOLEAN"`        // Флаг доступности покупки
	SellAvailableFlag     bool              `json:"sellAvailableFlag" gorm:"type:BOOLEAN"`       // Флаг доступности продажи
	DivYieldFlag          bool              `json:"divYieldFlag" gorm:"type:BOOLEAN"`            // Флаг дивидендной доходности
	ShareType             string            `json:"shareType" gorm:"type:VARCHAR(100)"`          // Тип акции
	MinPriceIncrement     MinPriceIncrement `json:"minPriceIncrement" gorm:"-"`                  // Внешний ключ
	ApiTradeAvailableFlag bool              `json:"apiTradeAvailableFlag" gorm:"type:BOOLEAN"`   // API-доступность
	Uid                    string            `json:"uid" gorm:"primaryKey;type:VARCHAR(255)"`     // Первичный ключ
	RealExchange          string            `json:"realExchange" gorm:"type:VARCHAR(100)"`       // Реальная биржа
	PositionUid           string            `json:"positionUid" gorm:"type:VARCHAR(255)"`        // UID позиции
	AssetUid              string            `json:"assetUid" gorm:"type:VARCHAR(255)"`           // UID актива
	InstrumentExchange    string            `json:"instrumentExchange" gorm:"type:VARCHAR(100)"` // Биржа инструмента
	ForIisFlag            bool              `json:"forIisFlag" gorm:"type:BOOLEAN"`              // Доступность для ИИС
	ForQualInvestorFlag   bool              `json:"forQualInvestorFlag" gorm:"type:BOOLEAN"`     // Только для квалифицированных инвесторов
	WeekendFlag           bool              `json:"weekendFlag" gorm:"type:BOOLEAN"`             // Торги в выходные
	BlockedTcaFlag        bool              `json:"blockedTcaFlag" gorm:"type:BOOLEAN"`          // Флаг блокировки TCA
	LiquidityFlag         bool              `json:"liquidityFlag" gorm:"type:BOOLEAN"`           // Флаг ликвидности
	First1minCandleDate   time.Time         `json:"first1minCandleDate" gorm:"index"`            // Индексируемое время
	First1dayCandleDate   time.Time         `json:"first1dayCandleDate" gorm:"index"`            // Индексируемое время
	Brand                 brand             `json:"brand" gorm:"-"`                              // Игнорируемое поле

	PlacementPriceID uint `json:"-" gorm:"not null;index;foreignKey:PlacementPriceID;references:uid"`
}

type Nominal struct {
	Currency         string `json:"currency" gorm:"type:VARCHAR(50)"` // Валюта
	Units            string `json:"units" gorm:"type:VARCHAR(100)"`   // Единицы
	Nano             int    `json:"nano" gorm:"type:INT"`             // Нано-части
	PlacementPriceID uint   `json:"-" gorm:"index"`                   // Внешний ключ
}

type MinPriceIncrement struct {
	Units            string `json:"units" gorm:"type:VARCHAR(100)"` // Единицы
	Nano             int    `json:"nano" gorm:"type:INT"`           // Нано-части
	PlacementPriceID uint   `json:"-" gorm:"index"`                 // Внешний ключ
}

type brand struct {
	logoName      string
	logoBaseColor string
	textColor     string
}

type Instruments struct {
	Instruments []PlacementPrice `json:"instruments"`
}

type InstrumentPart struct {
	MessageID string `json:"message_id"`
	Part      int    `json:"part"`
	Total     int    `json:"total"`
	Data      []byte `json:"data"`
}
