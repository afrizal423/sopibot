package bot

import "math"

type Model struct {
	Itemid                        int64  `json:"itemid"`
	Status                        int    `json:"status"`
	CurrentPromotionReservedStock int    `json:"current_promotion_reserved_stock"`
	Name                          string `json:"name"`
	Promotionid                   int    `json:"promotionid"`
	Price                         int64  `json:"price"`
	PriceStocks                   []struct {
		AllocatedStock           int         `json:"allocated_stock"`
		StockBreakdownByLocation interface{} `json:"stock_breakdown_by_location"`
	} `json:"price_stocks"`
	CurrentPromotionHasReserveStock bool `json:"current_promotion_has_reserve_stock"`
	NormalStock                     int  `json:"normal_stock"`
	Extinfo                         struct {
		TierIndex     []int       `json:"tier_index"`
		GroupBuyInfo  interface{} `json:"group_buy_info"`
		IsPreOrder    interface{} `json:"is_pre_order"`
		EstimatedDays int         `json:"estimated_days"`
	} `json:"extinfo"`
	PriceBeforeDiscount int64 `json:"price_before_discount"`
	Modelid             int64 `json:"modelid"`
	Stock               int   `json:"stock"`
	HasGimmickTag       bool  `json:"has_gimmick_tag"`
}

func (self *Model) is_available() bool {
	return self.Stock != 0
}

func (price *Model) Get_price() int64 {
	return int64(math.Floor(float64(price.Price) / float64(99999)))
}

type UpcomingFlashSaleInfo struct {
	End_time              int
	Item_id               int
	Name                  string
	Price                 int64
	Price_before_discount int
	Promotion_id          int
	Shop_id               int
	Start_time            int
	Stock                 int
}

type AddOnDealInfo struct {
	Add_on_deal_id    int
	Add_on_deal_label string
	Sub_type          int
}

type Item struct {
	Item_id               int64
	Shop_id               int
	Models                []Model
	Name                  string
	Price                 float64
	Price_before_discount int
	Brand                 string
	Shop_location         string
	Upcoming_flash_sale   UpcomingFlashSaleInfo
	Add_on_deal_info      AddOnDealInfo
	Price_min             int
	Price_max             int
	Stock                 int
	Is_flash_sale         bool
}

func (price *Item) Get_price() float64 {
	return math.Floor(price.Price / float64(99999))
}

type CartItem struct {
	Add_on_deal_id int
	Item_group_id  string
	Item_id        int
	Model_id       int
	Price          int64
	Shop_id        int
}
