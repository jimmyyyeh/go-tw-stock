package schema

import "time"

type StockInfo struct {
	Code               string    `json:"code"`                 // 股票代碼
	Name               string    `json:"name"`                 // 公司名稱
	ShortName          string    `json:"short_name"`           // 公司簡稱
	Type               string    `json:"type"`                 // 上市/上櫃
	Category           string    `json:"category"`             // 產業別
	Address            string    `json:"address"`              // 地址
	UnifiedNo          string    `json:"unified_no"`           // 統一編號
	IncorporationDay   time.Time `json:"incorporation_day"`    // 成立日期
	ListingDate        time.Time `json:"listing_date"`         // 上市/櫃日期
	Value              float64   `json:"value"`                // 普通股每股面額
	Capital            float64   `json:"capital"`              // 資本額
	PrivateStockShares float64   `json:"private_stock_shares"` // 私募股數
	PreferredShares    float64   `json:"preferred_shares"`     // 特別股
}

type StockTrading struct {
	Code        string  `json:"code"`         // 股票代碼
	TradeVolume float64 `json:"trade_volume"` // 成交股數
	Transaction float64 `json:"transaction"`  // 成交筆數
	TradeValue  float64 `json:"trade_value"`  // 成交金額
	Open        float64 `json:"open"`         // 開盤價
	High        float64 `json:"high"`         // 最高價
	Low         float64 `json:"low"`          // 最低價
	Close       float64 `json:"close"`        // 收盤價
	Change      float64 `json:"change"`       // 漲跌價差
}
