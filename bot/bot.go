package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/afrizal423/sopibot/checkout"
)

type Itemnya struct {
	Error    interface{} `json:"error"`
	ErrorMsg interface{} `json:"error_msg"`
	Data     struct {
		Itemid                  int64       `json:"itemid"`
		Shopid                  int         `json:"shopid"`
		Userid                  int         `json:"userid"`
		PriceMaxBeforeDiscount  int64       `json:"price_max_before_discount"`
		HasLowestPriceGuarantee bool        `json:"has_lowest_price_guarantee"`
		PriceBeforeDiscount     int64       `json:"price_before_discount"`
		PriceMinBeforeDiscount  int64       `json:"price_min_before_discount"`
		ExclusivePriceInfo      interface{} `json:"exclusive_price_info"`
		HiddenPriceDisplay      interface{} `json:"hidden_price_display"`
		PriceMin                int64       `json:"price_min"`
		PriceMax                int64       `json:"price_max"`
		Price                   int64       `json:"price"`
		Stock                   int         `json:"stock"`
		Discount                string      `json:"discount"`
		HistoricalSold          int         `json:"historical_sold"`
		Sold                    int         `json:"sold"`
		ShowDiscount            int         `json:"show_discount"`
		RawDiscount             int         `json:"raw_discount"`
		MinPurchaseLimit        int         `json:"min_purchase_limit"`
		OverallPurchaseLimit    struct {
			OrderMaxPurchaseLimit int         `json:"order_max_purchase_limit"`
			OverallPurchaseLimit  interface{} `json:"overall_purchase_limit"`
			ItemOverallQuota      interface{} `json:"item_overall_quota"`
			StartDate             interface{} `json:"start_date"`
			EndDate               interface{} `json:"end_date"`
		} `json:"overall_purchase_limit"`
		PackSize                interface{} `json:"pack_size"`
		IsLiveStreamingPrice    interface{} `json:"is_live_streaming_price"`
		Name                    string      `json:"name"`
		Ctime                   int         `json:"ctime"`
		ItemStatus              string      `json:"item_status"`
		Status                  int         `json:"status"`
		Condition               int         `json:"condition"`
		Catid                   int         `json:"catid"`
		Description             string      `json:"description"`
		IsMart                  bool        `json:"is_mart"`
		RichTextDescription     interface{} `json:"rich_text_description"`
		ShowShopeeVerifiedLabel bool        `json:"show_shopee_verified_label"`
		SizeChart               interface{} `json:"size_chart"`
		ReferenceItemID         string      `json:"reference_item_id"`
		Brand                   string      `json:"brand"`
		ItemRating              struct {
			RatingStar  float64 `json:"rating_star"`
			RatingCount []int   `json:"rating_count"`
		} `json:"item_rating"`
		LabelIds   []int `json:"label_ids"`
		Attributes []struct {
			Name        string      `json:"name"`
			Value       string      `json:"value"`
			ID          int         `json:"id"`
			IsTimestamp bool        `json:"is_timestamp"`
			BrandOption interface{} `json:"brand_option"`
			ValID       interface{} `json:"val_id"`
		} `json:"attributes"`
		Liked                 bool `json:"liked"`
		LikedCount            int  `json:"liked_count"`
		CmtCount              int  `json:"cmt_count"`
		Flag                  int  `json:"flag"`
		ShopeeVerified        bool `json:"shopee_verified"`
		IsAdult               bool `json:"is_adult"`
		IsPreferredPlusSeller bool `json:"is_preferred_plus_seller"`
		TierVariations        []struct {
			Name       string      `json:"name"`
			Options    []string    `json:"options"`
			Images     interface{} `json:"images"`
			Properties interface{} `json:"properties"`
			Type       int         `json:"type"`
		} `json:"tier_variations"`
		BundleDealID     int  `json:"bundle_deal_id"`
		CanUseBundleDeal bool `json:"can_use_bundle_deal"`
		AddOnDealInfo    struct {
			AddOnDealID    int    `json:"add_on_deal_id"`
			AddOnDealLabel string `json:"add_on_deal_label"`
			SubType        int    `json:"sub_type"`
			Status         int    `json:"status"`
		} `json:"add_on_deal_info"`
		BundleDealInfo               interface{}   `json:"bundle_deal_info"`
		CanUseWholesale              bool          `json:"can_use_wholesale"`
		WholesaleTierList            []interface{} `json:"wholesale_tier_list"`
		IsGroupBuyItem               interface{}   `json:"is_group_buy_item"`
		GroupBuyInfo                 interface{}   `json:"group_buy_info"`
		WelcomePackageType           int           `json:"welcome_package_type"`
		WelcomePackageInfo           interface{}   `json:"welcome_package_info"`
		TaxCode                      interface{}   `json:"tax_code"`
		InvoiceOption                interface{}   `json:"invoice_option"`
		ComplaintPolicy              interface{}   `json:"complaint_policy"`
		Images                       []string      `json:"images"`
		Image                        string        `json:"image"`
		VideoInfoList                interface{}   `json:"video_info_list"`
		ItemType                     int           `json:"item_type"`
		IsOfficialShop               bool          `json:"is_official_shop"`
		ShowOfficialShopLabelInTitle bool          `json:"show_official_shop_label_in_title"`
		ShopLocation                 string        `json:"shop_location"`
		CoinEarnLabel                interface{}   `json:"coin_earn_label"`
		CbOption                     int           `json:"cb_option"`
		IsPreOrder                   bool          `json:"is_pre_order"`
		EstimatedDays                int           `json:"estimated_days"`
		BadgeIconType                int           `json:"badge_icon_type"`
		ShowFreeShipping             bool          `json:"show_free_shipping"`
		ShippingIconType             int           `json:"shipping_icon_type"`
		CodFlag                      int           `json:"cod_flag"`
		IsServiceByShopee            bool          `json:"is_service_by_shopee"`
		ShowOriginalGuarantee        bool          `json:"show_original_guarantee"`
		Categories                   []struct {
			Catid           int    `json:"catid"`
			DisplayName     string `json:"display_name"`
			NoSub           bool   `json:"no_sub"`
			IsDefaultSubcat bool   `json:"is_default_subcat"`
		} `json:"categories"`
		OtherStock                      int  `json:"other_stock"`
		ItemHasPost                     bool `json:"item_has_post"`
		DiscountStock                   int  `json:"discount_stock"`
		CurrentPromotionHasReserveStock bool `json:"current_promotion_has_reserve_stock"`
		CurrentPromotionReservedStock   int  `json:"current_promotion_reserved_stock"`
		NormalStock                     int  `json:"normal_stock"`
		BrandID                         int  `json:"brand_id"`
		IsAlcoholProduct                bool `json:"is_alcohol_product"`
		ShowRecyclingInfo               bool `json:"show_recycling_info"`
		CoinInfo                        struct {
			SpendCashUnit int `json:"spend_cash_unit"`
			CoinEarnItems []struct {
				CoinEarn int `json:"coin_earn"`
			} `json:"coin_earn_items"`
		} `json:"coin_info"`
		Models []struct {
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
		} `json:"models"`
		SplInfo struct {
			InstallmentInfo interface{} `json:"installment_info"`
			UserCreditInfo  interface{} `json:"user_credit_info"`
			ChannelID       interface{} `json:"channel_id"`
			ShowSpl         bool        `json:"show_spl"`
			ShowSplLite     bool        `json:"show_spl_lite"`
		} `json:"spl_info"`
		PreviewInfo  interface{} `json:"preview_info"`
		FeCategories []struct {
			Catid           int    `json:"catid"`
			DisplayName     string `json:"display_name"`
			NoSub           bool   `json:"no_sub"`
			IsDefaultSubcat bool   `json:"is_default_subcat"`
		} `json:"fe_categories"`
		PresaleInfo                       interface{} `json:"presale_info"`
		IsCcInstallmentPaymentEligible    bool        `json:"is_cc_installment_payment_eligible"`
		IsNonCcInstallmentPaymentEligible bool        `json:"is_non_cc_installment_payment_eligible"`
		FlashSale                         interface{} `json:"flash_sale"`
		UpcomingFlashSale                 struct {
			FlashSaleType       int         `json:"flash_sale_type"`
			ExtraDiscountInfo   interface{} `json:"extra_discount_info"`
			Promotionid         int         `json:"promotionid"`
			StartTime           int         `json:"start_time"`
			EndTime             int         `json:"end_time"`
			PromoImages         []string    `json:"promo_images"`
			Price               int64       `json:"price"`
			FlashSaleStock      interface{} `json:"flash_sale_stock"`
			Stock               int         `json:"stock"`
			HiddenPriceDisplay  interface{} `json:"hidden_price_display"`
			PromoOverlayImage   interface{} `json:"promo_overlay_image"`
			PriceBeforeDiscount int64       `json:"price_before_discount"`
		} `json:"upcoming_flash_sale"`
		DeepDiscount               interface{} `json:"deep_discount"`
		HasLowFulfillmentRate      bool        `json:"has_low_fulfillment_rate"`
		IsPartialFulfilled         bool        `json:"is_partial_fulfilled"`
		Makeups                    interface{} `json:"makeups"`
		ShopVouchers               interface{} `json:"shop_vouchers"`
		GlobalSold                 interface{} `json:"global_sold"`
		IsInfantMilkFormulaProduct bool        `json:"is_infant_milk_formula_product"`
		ShouldShowAmpTag           bool        `json:"should_show_amp_tag"`
	} `json:"data"`
	IsIndexable bool `json:"is_indexable"`
}

type AddtoCardStruct struct {
	Checkout             bool   `json:"checkout"`
	ClientSource         int    `json:"client_source"`
	Donot_add_quantit    bool   `json:"donot_add_quantity"`
	Itemid               int64  `json:"itemid"`
	Modelid              int64  `json:"modelid"`
	Quantity             int    `json:"quantity"`
	Shopid               int    `json:"shopid"`
	Source               string `json:"source"`
	Update_checkout_only bool   `json:"update_checkout_only"`
}

type statusCard struct {
	Error    int    `json:"error"`
	ErrorMsg string `json:"error_msg"`
	Data     struct {
		ProblematicItems interface{} `json:"problematic_items"`
		CartItem         struct {
			Itemid      int         `json:"itemid"`
			Modelid     int64       `json:"modelid"`
			ItemGroupID interface{} `json:"item_group_id"`
			Currency    string      `json:"currency"`
			Price       int64       `json:"price"`
			Quantity    int         `json:"quantity"`
		} `json:"cart_item"`
	} `json:"data"`
}

var (
	Itemid int64
	Shopid int
)

/*

param url: the item url
return: Item object
the url will definitely be one of these:
- https://shopee.co.id/product/xxxx/xxxx
- https://shopee.co.id/Item-Name.xxxx.xxxx /\./g
*/
func Fetch_item_from_url(url string, csrf string) Item {
	//https://shopee.co.id/produk/Item-Name.xxxx.xxxx
	var str2 = getParams(`/(?P<shopid>\d+)/(?P<itemid>\d+).`, url)
	if str2["itemid"] != "" {
		return fetch_item(str2["itemid"], str2["shopid"], csrf)
	}

	//https://shopee.co.id/Item-Name.xxxx.xxxx
	var str1 = getParams(`\.(?P<shopid>\d+)\.(?P<itemid>\d+)`, url)
	if str1["itemid"] != "" {
		return fetch_item(str1["itemid"], str1["shopid"], csrf)
	}
	// fmt.Println(str1["itemid"])
	return fetch_item("0", "0", csrf)
}

func fetch_item(item_id string, shop_id string, csrf string) Item {
	params := url.Values{}
	params.Add("itemid", item_id)
	params.Add("shopid", shop_id)
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://shopee.co.id/api/v4/item/get?"+params.Encode(), nil)
	if err != nil {
		fmt.Println("No response from request")
	}
	req.Header = http.Header{
		"Accept":        []string{"application/json"},
		"Content-Type":  []string{"application/json"},
		"Referer":       []string{"https://shopee.co.id/"},
		"User-Agent":    []string{os.Getenv("USERAGENT")},
		"Cookie":        []string{os.Getenv("COOKIE")},
		"X-Csrftoken":   []string{csrf},
		"if-none-match": []string{"*"},
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("No response from request")
	}
	body, _ := ioutil.ReadAll(res.Body)
	var result Itemnya
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	// fmt.Println("itemid =", result.Item.Models[0].Itemid)
	// fmt.Println("modelid =", result.Item.Models[0].Modelid)
	dt := Item{
		Item_id: result.Data.Itemid,
		Shop_id: result.Data.Shopid,
		Models: func() []Model {
			mySlice := []Model{}
			for _, name := range result.Data.Models {
				mySlice = append(
					mySlice,
					Model{
						Itemid:      name.Itemid,
						Modelid:     name.Modelid,
						Promotionid: name.Promotionid,
						Name:        name.Name,
						Price:       name.Price,
						Stock:       name.Stock,
					})
			}
			return mySlice
		}(),
		Name:                  result.Data.Name,
		Price:                 float64(result.Data.Price),
		Price_before_discount: int(result.Data.PriceBeforeDiscount),
		Brand:                 result.Data.Brand,
		Shop_location:         result.Data.ShopLocation,
		Price_min:             int(result.Data.PriceMin),
		Price_max:             int(result.Data.PriceMax),
		Stock:                 result.Data.Stock,
		Is_flash_sale: func() bool {
			if result.Data.FlashSale != nil {
				return true
			}
			return false
		}(),
		Upcoming_flash_sale: func() UpcomingFlashSaleInfo {
			dt := UpcomingFlashSaleInfo{}
			if result.Data.UpcomingFlashSale.EndTime != 0 {
				dt.End_time = int(result.Data.UpcomingFlashSale.EndTime)
				dt.Price = result.Data.UpcomingFlashSale.Price
				dt.Price_before_discount = int(result.Data.UpcomingFlashSale.PriceBeforeDiscount)
				dt.Promotion_id = result.Data.UpcomingFlashSale.Promotionid
				dt.Start_time = result.Data.UpcomingFlashSale.StartTime
				dt.Stock = result.Data.UpcomingFlashSale.Stock

				return dt
			}
			return dt
		}(),
	}
	// fmt.Println(dt)
	return dt
}

func Add_to_card(i Item, model_index int, csrf string) Model {
	// fmt.Println("Kamu memilih produk", i.Name)

	if !i.Models[model_index].is_available() {
		panic(fmt.Errorf("Exception: %v", "out of stock"))
	}
	// fmt.Println(model_index)
	fmt.Println("Dengan model", i.Models[model_index].Name)
	client := http.Client{}
	dt := &AddtoCardStruct{
		Checkout:             true,
		ClientSource:         1,
		Donot_add_quantit:    false,
		Itemid:               i.Item_id,
		Modelid:              i.Models[model_index].Modelid,
		Quantity:             1,
		Shopid:               i.Shop_id,
		Source:               "",
		Update_checkout_only: false,
	}
	res1B, _ := json.Marshal(dt)
	req, err := http.NewRequest("POST", "https://shopee.co.id/api/v4/cart/add_to_cart", bytes.NewBuffer(res1B))
	if err != nil {
		fmt.Println("No response from request")
	}
	req.Header = http.Header{
		"Accept":        []string{"application/json"},
		"Content-Type":  []string{"application/json"},
		"Referer":       []string{"https://shopee.co.id/"},
		"User-Agent":    []string{os.Getenv("USERAGENT")},
		"Cookie":        []string{os.Getenv("COOKIE")},
		"X-Csrftoken":   []string{csrf},
		"if-none-match": []string{"*"},
	}
	// fmt.Println(req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("No response from request")
	}

	body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))
	var h statusCard
	if err := json.Unmarshal(body, &h); err != nil {
		fmt.Println(err)
	}
	if h.Error == 0 {
		fmt.Println()
		fmt.Println("✅Barang sudah ada di keranjang✅")
	} else {
		fmt.Println()
		fmt.Printf("Error: %s", h.ErrorMsg)
	}
	// fmt.Println(string(body))
	return i.Models[model_index]

}

/**
 * Parses url with the given regular expression and returns the
 * group values defined in the expression.
 *
 */
func getParams(regEx, url string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}

func Checkout(itemid Model, item Item, no_item int, csrf string) {
	co := checkout.Checkout{
		Csrf:   csrf,
		Itemid: itemid.Itemid,
		Shopid: item.Shop_id,
	}
	co.GetUserAddressList()
	co.GoCheckOut(item.Models[no_item].Modelid)
	// usernya := checkout.GetUserAddressList(csrf)

}
