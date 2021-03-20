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

type StockInstitutionalInvestors struct {
	Code                      string  `json:"code"`                         // 股票代碼
	ForeignDealersExcludeBuy  float64 `json:"foreign_dealers_exclude_buy"`  // 外資買進(不含自營商)
	ForeignDealersExcludeSell float64 `json:"foreign_dealers_exclude_sell"` // 外資賣出(不含自營商)
	ForeignDealersBuy         float64 `json:"foreign_buy"`                  // 外資買進(含自營商)
	ForeignDealersSell        float64 `json:"foreign_sell"`                 // 外資賣出(含自營商)
	TrustBuy                  float64 `json:"trust_buy"`                    // 投信買進
	TrustSell                 float64 `json:"trust_sell"`                   // 投信賣出
	DealersProprietaryBuy     float64 `json:"dealers_proprietary_buy"`      // 自營商買進
	DealersProprietarySell    float64 `json:"dealers_proprietary_sell"`     // 自營商賣出
	DealersHedgeBuy           float64 `json:"dealers_hedge_buy"`            // 自營商買進(避險)
	DealersHedgeSell          float64 `json:"dealers_hedge_sell"`           // 自營商賣出(避險)
	Total                     float64 `json:"total"`                        // 總計
}

type StockMarginTrading struct {
	Code                       string  `json:"code"`                          // 股票代碼
	MarginPurchase             float64 `json:"margin_purchase"`               // 融資買進
	MarginSale                 float64 `json:"margin_sale"`                   // 融資賣出
	CashRedemption             float64 `json:"cash_redemption"`               // 現金償還
	MarginBalanceOfPreviousDay float64 `json:"balance_of_previous_day"`       // 融資昨日償還
	MarginBalanceOfTheDay      float64 `json:"margin_balance_of_theDay"`      // 融資今日償還
	MarginQuota                float64 `json:"margin_quota"`                  // 融資限額
	ShortCovering              float64 `json:"short_covering"`                // 融券買進
	ShortSale                  float64 `json:"short_sale"`                    // 融券賣出
	StockRedemption            float64 `json:"stock_redemption"`              // 現券償還
	ShortBalanceOfPreviousDay  float64 `json:"short_balance_of_previous_day"` // 融券昨日償還
	ShortBalanceOfTheDay       float64 `json:"short_balance_of_theDay"`       // 融券今日償還
	ShortQuota                 float64 `json:"short_quota"`                   // 融券限額
	OffSetting                 float64 `json:"off_setting"`                   // 資券互抵
	Note                       string  `json:"note"`                          // 備註
}
