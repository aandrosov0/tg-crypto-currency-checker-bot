package coinranking

const (
	BTC = "Qwsogvtv82FCd"
	ETH = "razxDUgYGNAdQ"
	TON = "67YlI0K1b"
)

type Coin struct {
	UUID                  string   `json:"uuid"`
	Symbol                string   `json:"symbol"`
	Name                  string   `json:"name"`
	Description           string   `json:"description"`
	Color                 string   `json:"color"`
	IconURL               string   `json:"iconUrl"`
	WebsiteURL            string   `json:"websiteUrl"`
	NumberOfMarkets       int      `json:"numberOfMarkets"`
	NumberOfExchanges     int      `json:"numberOfExchanges"`
	Two4HVolume           string   `json:"24hVolume"`
	MarketCap             string   `json:"marketCap"`
	FullyDilutedMarketCap string   `json:"fullyDilutedMarketCap"`
	Price                 string   `json:"price"`
	BtcPrice              string   `json:"btcPrice"`
	PriceAt               int      `json:"priceAt"`
	Change                string   `json:"change"`
	Rank                  int      `json:"rank"`
	Sparkline             []string `json:"sparkline"`
	CoinrankingURL        string   `json:"coinrankingUrl"`
	Tier                  int      `json:"tier"`
	LowVolume             bool     `json:"lowVolume"`
	ListedAt              int      `json:"listedAt"`
}
