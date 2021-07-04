package depot

type DepotApiResponse struct {
	Paging Paging  `json:"paging"`
	Depots []Depot `json:"values"`
}

type DepotPortfolioApiResponse struct {
	Paging     Paging     `json:"paging"`
	Aggregated Aggregated `json:"aggregated"`
	Positions  []Position `json:"values"`
}

type Depot struct {
	DepotID                    string   `json:"depotId"`
	DepotDisplayID             string   `json:"depotDisplayId"`
	ClientID                   string   `json:"clientId"`
	DefaultSettlementAccountId string   `json:"defaultSettlementAccountId"`
	SettlementAccountIds       []string `json:"settlementAccountIds"`
}

type Paging struct {
	Index   int `json:"index"`
	Matches int `json:"matches"`
}

type Aggregated struct {
	Depot                 Depot                 `json:"depot"`
	Prevdayvalue          Prevdayvalue          `json:"prevDayValue"`
	Currentvalue          Currentvalue          `json:"currentValue"`
	Purchasevalue         Purchasevalue         `json:"purchaseValue"`
	Profitlosspurchaseabs Profitlosspurchaseabs `json:"profitLossPurchaseAbs"`
	Profitlosspurchaserel string                `json:"profitLossPurchaseRel"`
	Profitlossprevdayabs  Profitlossprevdayabs  `json:"profitLossPrevDayAbs"`
	Profitlossprevdayrel  string                `json:"profitLossPrevDayRel"`
}

type Prevdayvalue struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type Currentvalue struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type Purchasevalue struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type Profitlosspurchaseabs struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type Profitlossprevdayabs struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type Quantity struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type Availablequantity struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type Price struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type Currentprice struct {
	Price         Price  `json:"price"`
	Pricedatetime string `json:"priceDateTime"`
}
type Purchaseprice struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type Prevdayprice struct {
	Price         Price  `json:"price"`
	Pricedatetime string `json:"priceDateTime"`
}
type Availablequantitytohedge struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type Position struct {
	Depotid                  string                   `json:"depotId"`
	Positionid               string                   `json:"positionId"`
	Wkn                      string                   `json:"wkn"`
	Custodytype              string                   `json:"custodyType"`
	Quantity                 Quantity                 `json:"quantity"`
	Availablequantity        Availablequantity        `json:"availableQuantity"`
	Currentprice             Currentprice             `json:"currentPrice"`
	Purchaseprice            Purchaseprice            `json:"purchasePrice"`
	Prevdayprice             Prevdayprice             `json:"prevDayPrice"`
	Currentvalue             Currentvalue             `json:"currentValue"`
	Purchasevalue            Purchasevalue            `json:"purchaseValue"`
	Profitlosspurchaseabs    Profitlosspurchaseabs    `json:"profitLossPurchaseAbs"`
	Profitlosspurchaserel    string                   `json:"profitLossPurchaseRel"`
	Profitlossprevdayabs     Profitlossprevdayabs     `json:"profitLossPrevDayAbs"`
	Profitlossprevdayrel     string                   `json:"profitLossPrevDayRel"`
	Version                  interface{}              `json:"version"`
	Hedgeability             string                   `json:"hedgeability"`
	Availablequantitytohedge Availablequantitytohedge `json:"availableQuantityToHedge"`
	Currentpricedeterminable bool                     `json:"currentPriceDeterminable"`
}
