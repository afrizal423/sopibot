package checkout

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/afrizal423/sopibot/pengguna"
)

const (
	ALFAMART       = 8003200
	INDOMART_ISAKU = 8003001
	AKULAKU        = 8000700
	// Akulaku
	AKULAKU_CICILAN_1X  = "8000700-25"
	AKULAKU_CICILAN_2X  = "8000700-26"
	AKULAKU_CICILAN_3X  = "8000700-27"
	AKULAKU_CICILAN_6X  = "8000700-28"
	AKULAKU_CICILAN_9X  = "8000700-29"
	AKULAKU_CICILAN_12X = "8000700-30"
	TRANSFER_BANK       = 8005200
	// Transfer Bank
	TRANSFER_BANK_BCA_AUTO     = "89052001"
	TRANSFER_BANK_MANDIRI_AUTO = "89052002"
	TRANSFER_BANK_BNI_AUTO     = "89052003"
	TRANSFER_BANK_BRI_AUTO     = "89052004"
	TRANSFER_BANK_SYARIAH_AUTO = "89052005"
	TRANSFER_BANK_PERMATA_AUTO = "89052006"
	COD_BAYAR_DI_TEMPAT        = 89000
)

type Checkout struct {
	Csrf   string
	Itemid int64
	Shopid int
}

type GetQuickPayload struct {
	Cft        []int `json:"_cft"`
	Shoporders []struct {
		Shop struct {
			Shopid int `json:"shopid"`
		} `json:"shop"`
		Items []struct {
			Itemid         int64         `json:"itemid"`
			Modelid        int64         `json:"modelid"`
			Quantity       int           `json:"quantity"`
			AddOnDealID    int           `json:"add_on_deal_id"`
			IsAddOnSubItem bool          `json:"is_add_on_sub_item"`
			ItemGroupID    string        `json:"item_group_id"`
			Insurances     []interface{} `json:"insurances"`
		} `json:"items"`
	} `json:"shoporders"`
	SelectedPaymentChannelData struct {
	} `json:"selected_payment_channel_data"`
	PromotionData struct {
		UseCoins                bool `json:"use_coins"`
		FreeShippingVoucherInfo struct {
			FreeShippingVoucherID int    `json:"free_shipping_voucher_id"`
			DisabledReason        string `json:"disabled_reason"`
			Description           string `json:"description"`
		} `json:"free_shipping_voucher_info"`
		PlatformVouchers          []interface{} `json:"platform_vouchers"`
		ShopVouchers              []interface{} `json:"shop_vouchers"`
		CheckShopVoucherEntrances bool          `json:"check_shop_voucher_entrances"`
		AutoApplyShopVoucher      bool          `json:"auto_apply_shop_voucher"`
	} `json:"promotion_data"`
	DeviceInfo struct {
		DeviceID          string `json:"device_id"`
		DeviceFingerprint string `json:"device_fingerprint"`
		TongdunBlackbox   string `json:"tongdun_blackbox"`
		BuyerPaymentInfo  struct {
		} `json:"buyer_payment_info"`
	} `json:"device_info"`
	TaxInfo struct {
		TaxID string `json:"tax_id"`
	} `json:"tax_info"`
}

type GetQuickResponse struct {
	ClientID          int `json:"client_id"`
	CartType          int `json:"cart_type"`
	Timestamp         int `json:"timestamp"`
	CheckoutPriceData struct {
		MerchandiseSubtotal             int         `json:"merchandise_subtotal"`
		ShippingSubtotalBeforeDiscount  int         `json:"shipping_subtotal_before_discount"`
		ShippingDiscountSubtotal        int         `json:"shipping_discount_subtotal"`
		ShippingSubtotal                int         `json:"shipping_subtotal"`
		TaxPayable                      int         `json:"tax_payable"`
		TaxExemption                    int         `json:"tax_exemption"`
		CustomTaxSubtotal               int         `json:"custom_tax_subtotal"`
		PromocodeApplied                interface{} `json:"promocode_applied"`
		CreditCardPromotion             interface{} `json:"credit_card_promotion"`
		ShopeeCoinsRedeemed             interface{} `json:"shopee_coins_redeemed"`
		GroupBuyDiscount                int         `json:"group_buy_discount"`
		BundleDealsDiscount             interface{} `json:"bundle_deals_discount"`
		BuyerTxnFee                     int         `json:"buyer_txn_fee"`
		InsuranceSubtotal               int         `json:"insurance_subtotal"`
		InsuranceBeforeDiscountSubtotal int         `json:"insurance_before_discount_subtotal"`
		InsuranceDiscountSubtotal       int         `json:"insurance_discount_subtotal"`
		VatSubtotal                     int         `json:"vat_subtotal"`
		TotalPayable                    int         `json:"total_payable"`
	} `json:"checkout_price_data"`
	OrderUpdateInfo struct {
	} `json:"order_update_info"`
	Shoporders []struct {
		Shop struct {
			Shopid          int    `json:"shopid"`
			ShopName        string `json:"shop_name"`
			CbOption        bool   `json:"cb_option"`
			IsOfficialShop  bool   `json:"is_official_shop"`
			RemarkType      int    `json:"remark_type"`
			SupportEreceipt bool   `json:"support_ereceipt"`
			SellerUserID    int    `json:"seller_user_id"`
			ShopTag         int    `json:"shop_tag"`
		} `json:"shop"`
		Items []struct {
			Itemid                  int64         `json:"itemid"`
			Modelid                 int64         `json:"modelid"`
			Quantity                int           `json:"quantity"`
			ItemGroupID             interface{}   `json:"item_group_id"`
			Insurances              []interface{} `json:"insurances"`
			Shopid                  int           `json:"shopid"`
			Shippable               bool          `json:"shippable"`
			NonShippableErr         string        `json:"non_shippable_err"`
			NoneShippableReason     string        `json:"none_shippable_reason"`
			NoneShippableFullReason string        `json:"none_shippable_full_reason"`
			Price                   int           `json:"price"`
			Name                    string        `json:"name"`
			ModelName               string        `json:"model_name"`
			AddOnDealID             int           `json:"add_on_deal_id"`
			IsAddOnSubItem          bool          `json:"is_add_on_sub_item"`
			IsPreOrder              bool          `json:"is_pre_order"`
			IsStreamingPrice        bool          `json:"is_streaming_price"`
			Image                   string        `json:"image"`
			Checkout                bool          `json:"checkout"`
			Categories              []struct {
				Catids []int `json:"catids"`
			} `json:"categories"`
			IsSplZeroInterest bool `json:"is_spl_zero_interest"`
		} `json:"items"`
		TaxInfo struct {
			UseNewCustomTaxMsg  bool   `json:"use_new_custom_tax_msg"`
			CustomTaxMsg        string `json:"custom_tax_msg"`
			CustomTaxMsgShort   string `json:"custom_tax_msg_short"`
			RemoveCustomTaxHint bool   `json:"remove_custom_tax_hint"`
		} `json:"tax_info"`
		TaxPayable                int         `json:"tax_payable"`
		ShippingID                int         `json:"shipping_id"`
		ShippingFeeDiscount       int         `json:"shipping_fee_discount"`
		ShippingFee               int         `json:"shipping_fee"`
		OrderTotalWithoutShipping int         `json:"order_total_without_shipping"`
		OrderTotal                int         `json:"order_total"`
		BuyerRemark               interface{} `json:"buyer_remark"`
	} `json:"shoporders"`
	CanCheckout bool `json:"can_checkout"`
}

func (c *Checkout) GetUserAddressList() pengguna.AlamatCheckout {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://shopee.co.id/api/v1/account_info", nil)
	if err != nil {
		fmt.Println("No response from request")
	}
	req.Header = http.Header{
		"Accept":        []string{"application/json"},
		"Content-Type":  []string{"application/json"},
		"Referer":       []string{"https://shopee.co.id/"},
		"User-Agent":    []string{os.Getenv("USERAGENT")},
		"Cookie":        []string{os.Getenv("COOKIE")},
		"X-Csrftoken":   []string{c.Csrf},
		"if-none-match": []string{"*"},
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("No response from request")
	}
	body, _ := ioutil.ReadAll(res.Body) // response body is []byte
	var result pengguna.AlamatCheckout
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return result
}

func (c *Checkout) GetCheckout(modelid int64, cardget ResponseCardGet) ResponseCheckOutGetTanpaAlfa {
	client := http.Client{}
	dt := &GetQuickPayload{}
	dt.TaxInfo.TaxID = ""
	dt.Cft = append(dt.Cft, 107)
	dt.PromotionData.AutoApplyShopVoucher = false
	dt.PromotionData.CheckShopVoucherEntrances = true
	dt.PromotionData.PlatformVouchers = append(dt.PromotionData.PlatformVouchers)
	dt.PromotionData.ShopVouchers = append(dt.PromotionData.ShopVouchers)
	dt.PromotionData.UseCoins = false
	dt.PromotionData.FreeShippingVoucherInfo.Description = ""
	dt.PromotionData.FreeShippingVoucherInfo.DisabledReason = ""
	dt.PromotionData.FreeShippingVoucherInfo.FreeShippingVoucherID = 0
	dt.DeviceInfo.DeviceFingerprint = ""
	dt.DeviceInfo.DeviceID = ""
	dt.DeviceInfo.TongdunBlackbox = ""
	dt.DeviceInfo.BuyerPaymentInfo = struct{}{}
	dt.SelectedPaymentChannelData = struct{}{}
	dt.Shoporders = []struct {
		Shop struct {
			Shopid int "json:\"shopid\""
		} "json:\"shop\""
		Items []struct {
			Itemid         int64         "json:\"itemid\""
			Modelid        int64         "json:\"modelid\""
			Quantity       int           "json:\"quantity\""
			AddOnDealID    int           "json:\"add_on_deal_id\""
			IsAddOnSubItem bool          "json:\"is_add_on_sub_item\""
			ItemGroupID    string        "json:\"item_group_id\""
			Insurances     []interface{} "json:\"insurances\""
		} "json:\"items\""
	}{
		{
			Items: []struct {
				Itemid         int64         "json:\"itemid\""
				Modelid        int64         "json:\"modelid\""
				Quantity       int           "json:\"quantity\""
				AddOnDealID    int           "json:\"add_on_deal_id\""
				IsAddOnSubItem bool          "json:\"is_add_on_sub_item\""
				ItemGroupID    string        "json:\"item_group_id\""
				Insurances     []interface{} "json:\"insurances\""
			}{
				{
					Itemid:         c.Itemid,
					Modelid:        modelid,
					Quantity:       1,
					AddOnDealID:    cardget.Data.ShopOrders[0].Items[0].AddOnDealID,
					Insurances:     nil,
					IsAddOnSubItem: cardget.Data.ShopOrders[0].Items[0].IsAddOnSubItem,
					ItemGroupID:    cardget.Data.ShopOrders[0].Items[0].ItemGroupID,
				},
			},
			Shop: struct {
				Shopid int "json:\"shopid\""
			}{
				Shopid: c.Shopid,
			},
		},
	}

	res1B, _ := json.Marshal(dt)
	// fmt.Println(string(res1B))
	req, err := http.NewRequest("POST", "https://shopee.co.id/api/v4/checkout/get", bytes.NewBuffer(res1B))
	if err != nil {
		fmt.Println("No response from request")
	}
	req.Header = http.Header{
		"Accept":        []string{"application/json"},
		"Content-Type":  []string{"application/json"},
		"Referer":       []string{"https://shopee.co.id/checkout"},
		"User-Agent":    []string{os.Getenv("USERAGENT")},
		"Cookie":        []string{os.Getenv("COOKIE")},
		"X-Csrftoken":   []string{c.Csrf},
		"X-Api-Source":  []string{"pc"},
		"if-none-match": []string{"*"},
	}
	// fmt.Println(req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("No response from request")
	}
	body, _ := ioutil.ReadAll(res.Body) // response body is []byte
	// fmt.Println()
	// fmt.Println("get checkout pertama")
	// fmt.Println()
	// fmt.Println(string(body))
	var h ResponseCheckOutGetTanpaAlfa
	if err := json.Unmarshal(body, &h); err != nil { // Parse []byte to the go struct pointer
		fmt.Println(err)
	}
	// fmt.Println(h)
	// var data1 map[string]interface{}
	// json.Unmarshal(body, &data1)
	// fmt.Println(data1["checkout_price_data"])

	return h
}

func (c *Checkout) getCheckoutFinal(dt_old ResponseCheckOutGetTanpaAlfa, modelid int64) ResponseCheckOutGetFinal {
	client := http.Client{}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	timestamp := time.Now().In(loc).Unix()
	dt := &PayloadGetCheckOut{}
	dt.TaxInfo.TaxID = ""
	dt.Cft = append(dt.Cft, 107)
	dt.Timestamp = int(timestamp)
	dt.Shoporders = []struct {
		Shop struct {
			Shopid int "json:\"shopid\""
		} "json:\"shop\""
		Items []struct {
			Itemid         int64       "json:\"itemid\""
			Modelid        int64       "json:\"modelid\""
			Quantity       int         "json:\"quantity\""
			AddOnDealID    int         "json:\"add_on_deal_id\""
			IsAddOnSubItem bool        "json:\"is_add_on_sub_item\""
			ItemGroupID    interface{} "json:\"item_group_id\""
			Insurances     []struct {
				InsuranceProductID    string "json:\"insurance_product_id\""
				Name                  string "json:\"name\""
				Description           string "json:\"description\""
				ProductDetailPageURL  string "json:\"product_detail_page_url\""
				Premium               int    "json:\"premium\""
				PremiumBeforeDiscount int    "json:\"premium_before_discount\""
				InsuranceQuantity     int    "json:\"insurance_quantity\""
				Selected              bool   "json:\"selected\""
			} "json:\"insurances\""
		} "json:\"items\""
		ShippingID int "json:\"shipping_id\""
	}{
		{
			Shop: struct {
				Shopid int "json:\"shopid\""
			}{
				Shopid: c.Shopid,
			},
			Items: []struct {
				Itemid         int64       "json:\"itemid\""
				Modelid        int64       "json:\"modelid\""
				Quantity       int         "json:\"quantity\""
				AddOnDealID    int         "json:\"add_on_deal_id\""
				IsAddOnSubItem bool        "json:\"is_add_on_sub_item\""
				ItemGroupID    interface{} "json:\"item_group_id\""
				Insurances     []struct {
					InsuranceProductID    string "json:\"insurance_product_id\""
					Name                  string "json:\"name\""
					Description           string "json:\"description\""
					ProductDetailPageURL  string "json:\"product_detail_page_url\""
					Premium               int    "json:\"premium\""
					PremiumBeforeDiscount int    "json:\"premium_before_discount\""
					InsuranceQuantity     int    "json:\"insurance_quantity\""
					Selected              bool   "json:\"selected\""
				} "json:\"insurances\""
			}{
				{
					Itemid:         c.Itemid,
					Modelid:        modelid,
					Quantity:       1,
					AddOnDealID:    0,
					IsAddOnSubItem: dt_old.Shoporders[0].Items[0].IsAddOnSubItem,
					ItemGroupID:    dt_old.Shoporders[0].Items[0].ItemGroupID,
					Insurances:     dt_old.Shoporders[0].Items[0].Insurances,
				},
			},
		},
	}
	// Silahkan ubah sesuai selera metode pembayaran
	dt.SelectedPaymentChannelData.ChannelID = TRANSFER_BANK
	dt.SelectedPaymentChannelData.ChannelItemOptionInfo.OptionInfo = TRANSFER_BANK_BNI_AUTO
	dt.SelectedPaymentChannelData.Version = 2
	//end metode pembayaran
	dt.PromotionData.AutoApplyShopVoucher = false
	dt.PromotionData.CheckShopVoucherEntrances = true
	dt.PromotionData.PlatformVouchers = append(dt.PromotionData.PlatformVouchers)
	dt.PromotionData.ShopVouchers = append(dt.PromotionData.ShopVouchers)
	dt.PromotionData.UseCoins = false
	dt.PromotionData.FreeShippingVoucherInfo.DisabledReason = nil
	dt.PromotionData.FreeShippingVoucherInfo.FreeShippingVoucherCode = ""
	dt.PromotionData.FreeShippingVoucherInfo.FreeShippingVoucherID = 0
	dt.PromotionData.FreeShippingVoucherInfo.BannerInfo.Msg = ""
	dt.PromotionData.FreeShippingVoucherInfo.BannerInfo.LearnMoreMsg = ""
	dt.FsvSelectionInfos = append(dt.FsvSelectionInfos)
	dt.DeviceInfo.DeviceFingerprint = ""
	dt.DeviceInfo.DeviceID = ""
	dt.DeviceInfo.TongdunBlackbox = ""
	dt.DeviceInfo.BuyerPaymentInfo = struct{}{}
	dt.BuyerInfo.ShareToFriendsInfo.DisplayToggle = true
	dt.BuyerInfo.ShareToFriendsInfo.EnableToggle = false
	dt.BuyerInfo.ShareToFriendsInfo.AllowToShare = false
	dt.BuyerInfo.KycInfo = nil
	dt.BuyerInfo.CheckoutEmail = ""
	dt.CartType = 0
	dt.ClientID = 0
	dt.TaxInfo.TaxID = ""
	dt.DropshippingInfo.Enabled = false
	dt.DropshippingInfo.Name = ""
	dt.DropshippingInfo.PhoneNumber = ""
	dt.ShippingOrders = []struct {
		Sync             bool "json:\"sync\""
		BuyerAddressData struct {
			Addressid   int    "json:\"addressid\""
			AddressType int    "json:\"address_type\""
			TaxAddress  string "json:\"tax_address\""
		} "json:\"buyer_address_data\""
		SelectedLogisticChannelid             int         "json:\"selected_logistic_channelid\""
		ShippingID                            int         "json:\"shipping_id\""
		ShoporderIndexes                      []int       "json:\"shoporder_indexes\""
		SelectedPreferredDeliveryTimeOptionID int         "json:\"selected_preferred_delivery_time_option_id\""
		SelectedPreferredDeliveryTimeSlotID   interface{} "json:\"selected_preferred_delivery_time_slot_id\""
	}{
		{
			Sync: true,
			BuyerAddressData: struct {
				Addressid   int    "json:\"addressid\""
				AddressType int    "json:\"address_type\""
				TaxAddress  string "json:\"tax_address\""
			}{
				Addressid:   dt_old.ShippingOrders[0].BuyerAddressData.Addressid,
				AddressType: 0,
				TaxAddress:  "",
			},
			SelectedLogisticChannelid:             dt_old.ShippingOrders[0].SelectedLogisticChannelid,
			ShippingID:                            1,
			ShoporderIndexes:                      []int{0},
			SelectedPreferredDeliveryTimeOptionID: 1,
			SelectedPreferredDeliveryTimeSlotID:   nil,
		},
	}
	dt.OrderUpdateInfo = struct{}{}
	dt.CheckoutPriceData = dt_old.CheckoutPriceData
	res1B, _ := json.Marshal(dt)
	// fmt.Println(string(res1B))
	req, err := http.NewRequest("POST", "https://shopee.co.id/api/v4/checkout/get", bytes.NewBuffer(res1B))
	if err != nil {
		fmt.Println("No response from request")
	}
	req.Header = http.Header{
		"Accept":        []string{"application/json"},
		"Content-Type":  []string{"application/json"},
		"Referer":       []string{"https://shopee.co.id/checkout"},
		"User-Agent":    []string{os.Getenv("USERAGENT")},
		"Cookie":        []string{os.Getenv("COOKIE")},
		"X-Csrftoken":   []string{c.Csrf},
		"X-Api-Source":  []string{"pc"},
		"if-none-match": []string{"*"},
	}
	// fmt.Println(req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("No response from request")
	}
	body, _ := ioutil.ReadAll(res.Body) // response body is []byte
	// fmt.Println()
	// fmt.Println("get checkout final")
	// fmt.Println()
	// fmt.Println(string(body))
	var h ResponseCheckOutGetFinal
	if err := json.Unmarshal(body, &h); err != nil {
		fmt.Println(err)
	}
	// fmt.Println(h)
	return h
}

func (c *Checkout) place_holder(item_old ResponseCheckOutGetFinal) map[string]interface{} {
	// var devicenya string
	client := http.Client{}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	timestamp := time.Now().In(loc).Unix()
	dt := &PayloadPlaceHolder{}
	dt.Status = 200
	dt.Headers = struct{}{}
	dt.ClientID = 0
	dt.CartType = 0
	dt.Timestamp = int(timestamp)
	dt.CheckoutPriceData = item_old.CheckoutPriceData
	dt.OrderUpdateInfo = struct{}{}
	dt.DropshippingInfo.Enabled = false
	dt.DropshippingInfo.Name = ""
	dt.DropshippingInfo.PhoneNumber = ""
	dt.PromotionData = item_old.PromotionData
	dt.SelectedPaymentChannelData = item_old.SelectedPaymentChannelData
	dt.Shoporders = []struct {
		Shop struct {
			Shopid          int    "json:\"shopid\""
			ShopName        string "json:\"shop_name\""
			CbOption        bool   "json:\"cb_option\""
			IsOfficialShop  bool   "json:\"is_official_shop\""
			RemarkType      int    "json:\"remark_type\""
			SupportEreceipt bool   "json:\"support_ereceipt\""
			SellerUserID    int    "json:\"seller_user_id\""
			ShopTag         int    "json:\"shop_tag\""
		} "json:\"shop\""
		Items []struct {
			Itemid      int64       "json:\"itemid\""
			Modelid     int64       "json:\"modelid\""
			Quantity    int         "json:\"quantity\""
			ItemGroupID interface{} "json:\"item_group_id\""
			Insurances  []struct {
				InsuranceProductID    string "json:\"insurance_product_id\""
				Name                  string "json:\"name\""
				Description           string "json:\"description\""
				ProductDetailPageURL  string "json:\"product_detail_page_url\""
				Premium               int    "json:\"premium\""
				PremiumBeforeDiscount int    "json:\"premium_before_discount\""
				InsuranceQuantity     int    "json:\"insurance_quantity\""
				Selected              bool   "json:\"selected\""
			} "json:\"insurances\""
			Shopid                  int    "json:\"shopid\""
			Shippable               bool   "json:\"shippable\""
			NonShippableErr         string "json:\"non_shippable_err\""
			NoneShippableReason     string "json:\"none_shippable_reason\""
			NoneShippableFullReason string "json:\"none_shippable_full_reason\""
			Price                   int    "json:\"price\""
			Name                    string "json:\"name\""
			ModelName               string "json:\"model_name\""
			AddOnDealID             int    "json:\"add_on_deal_id\""
			IsAddOnSubItem          bool   "json:\"is_add_on_sub_item\""
			IsPreOrder              bool   "json:\"is_pre_order\""
			IsStreamingPrice        bool   "json:\"is_streaming_price\""
			Image                   string "json:\"image\""
			Checkout                bool   "json:\"checkout\""
			Categories              []struct {
				Catids []int "json:\"catids\""
			} "json:\"categories\""
			IsSplZeroInterest bool "json:\"is_spl_zero_interest\""
		} "json:\"items\""
		TaxInfo struct {
			UseNewCustomTaxMsg  bool   "json:\"use_new_custom_tax_msg\""
			CustomTaxMsg        string "json:\"custom_tax_msg\""
			CustomTaxMsgShort   string "json:\"custom_tax_msg_short\""
			RemoveCustomTaxHint bool   "json:\"remove_custom_tax_hint\""
		} "json:\"tax_info\""
		TaxPayable                int           "json:\"tax_payable\""
		ShippingID                int           "json:\"shipping_id\""
		ShippingFeeDiscount       int           "json:\"shipping_fee_discount\""
		ShippingFee               int           "json:\"shipping_fee\""
		OrderTotalWithoutShipping int           "json:\"order_total_without_shipping\""
		OrderTotal                int64         "json:\"order_total\""
		BuyerRemark               string        "json:\"buyer_remark\""
		ExtAdInfoMappings         []interface{} "json:\"ext_ad_info_mappings\""
	}{
		{
			Shop:                      item_old.Shoporders[0].Shop,
			Items:                     item_old.Shoporders[0].Items,
			TaxInfo:                   item_old.Shoporders[0].TaxInfo,
			TaxPayable:                0,
			ShippingID:                1,
			ShippingFeeDiscount:       0,
			ShippingFee:               item_old.Shoporders[0].ShippingFee,
			OrderTotalWithoutShipping: item_old.Shoporders[0].OrderTotalWithoutShipping,
			OrderTotal:                item_old.Shoporders[0].OrderTotal,
			BuyerRemark:               "",
			ExtAdInfoMappings:         []interface{}{},
		},
	}
	dt.ShippingOrders = []struct {
		ShippingID                            int    "json:\"shipping_id\""
		ShoporderIndexes                      []int  "json:\"shoporder_indexes\""
		SelectedLogisticChannelid             int    "json:\"selected_logistic_channelid\""
		SelectedPreferredDeliveryTimeOptionID int    "json:\"selected_preferred_delivery_time_option_id\""
		BuyerRemark                           string "json:\"buyer_remark\""
		BuyerAddressData                      struct {
			Addressid   int    "json:\"addressid\""
			AddressType int    "json:\"address_type\""
			TaxAddress  string "json:\"tax_address\""
		} "json:\"buyer_address_data\""
		FulfillmentInfo struct {
			FulfillmentFlag      int    "json:\"fulfillment_flag\""
			FulfillmentSource    string "json:\"fulfillment_source\""
			ManagedBySbs         bool   "json:\"managed_by_sbs\""
			OrderFulfillmentType int    "json:\"order_fulfillment_type\""
			WarehouseAddressID   int    "json:\"warehouse_address_id\""
			IsFromOverseas       bool   "json:\"is_from_overseas\""
		} "json:\"fulfillment_info\""
		OrderTotal                           int64       "json:\"order_total\""
		OrderTotalWithoutShipping            int         "json:\"order_total_without_shipping\""
		SelectedLogisticChannelidWithWarning interface{} "json:\"selected_logistic_channelid_with_warning\""
		ShippingFee                          int         "json:\"shipping_fee\""
		ShippingFeeDiscount                  int         "json:\"shipping_fee_discount\""
		ShippingGroupDescription             string      "json:\"shipping_group_description\""
		ShippingGroupIcon                    string      "json:\"shipping_group_icon\""
		TaxPayable                           int         "json:\"tax_payable\""
		IsFsvApplied                         bool        "json:\"is_fsv_applied\""
	}{
		{
			ShippingID:                            1,
			ShoporderIndexes:                      item_old.ShippingOrders[0].ShoporderIndexes,
			SelectedLogisticChannelid:             item_old.ShippingOrders[0].SelectedLogisticChannelid,
			SelectedPreferredDeliveryTimeOptionID: item_old.ShippingOrders[0].SelectedPreferredDeliveryTimeOptionID,
			BuyerRemark:                           "",
			BuyerAddressData:                      item_old.ShippingOrders[0].BuyerAddressData,
			FulfillmentInfo:                       item_old.ShippingOrders[0].FulfillmentInfo,
			OrderTotal:                            int64(item_old.ShippingOrders[0].OrderTotal),
			OrderTotalWithoutShipping:             item_old.ShippingOrders[0].OrderTotalWithoutShipping,
			SelectedLogisticChannelidWithWarning:  nil,
			ShippingFee:                           item_old.ShippingOrders[0].ShippingFee,
			ShippingFeeDiscount:                   item_old.ShippingOrders[0].ShippingFeeDiscount,
			ShippingGroupDescription:              item_old.ShippingOrders[0].ShippingGroupDescription,
			ShippingGroupIcon:                     item_old.ShippingOrders[0].ShippingGroupIcon,
			TaxPayable:                            item_old.ShippingOrders[0].TaxPayable,
			IsFsvApplied:                          item_old.ShippingOrders[0].IsFsvApplied,
		},
	}
	dt.FsvSelectionInfos = item_old.FsvSelectionInfos
	dt.BuyerInfo.ShareToFriendsInfo.DisplayToggle = true
	dt.BuyerInfo.ShareToFriendsInfo.EnableToggle = true
	dt.BuyerInfo.ShareToFriendsInfo.AllowToShare = true
	dt.BuyerTxnFeeInfo = item_old.BuyerTxnFeeInfo
	dt.DisabledCheckoutInfo = item_old.DisabledCheckoutInfo
	dt.CanCheckout = true
	dt.Cft = append(dt.Cft, 107)
	dt.DeviceInfo = struct {
		DeviceSzFingerprint string "json:\"device_sz_fingerprint\""
	}{
		DeviceSzFingerprint: os.Getenv("KEYBELANJA"),
	}
	dt.CaptchaVersion = 1
	res1B, _ := json.Marshal(dt)
	// fmt.Println()
	// fmt.Println("PAYLOAD checkout")
	// fmt.Println()
	// fmt.Println(string(res1B))
	req, err := http.NewRequest("POST", "https://shopee.co.id/api/v4/checkout/place_order", bytes.NewBuffer(res1B))
	if err != nil {
		fmt.Println("No response from request")
	}
	req.Header = http.Header{
		"Accept":        []string{"application/json"},
		"Content-Type":  []string{"application/json"},
		"Referer":       []string{"https://shopee.co.id/checkout"},
		"User-Agent":    []string{os.Getenv("USERAGENT")},
		"Cookie":        []string{os.Getenv("COOKIE")},
		"X-Csrftoken":   []string{c.Csrf},
		"X-Api-Source":  []string{"pc"},
		"if-none-match": []string{"*"},
	}
	// fmt.Println(req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("No response from request")
	}
	body, _ := ioutil.ReadAll(res.Body) // response body is []byte
	fmt.Println()
	fmt.Println("response checkout")
	fmt.Println()
	fmt.Println(string(body))
	var data1 map[string]interface{}
	json.Unmarshal(body, &data1)
	return data1
}

func (c *Checkout) card_get() ResponseCardGet {
	// var devicenya string
	client := http.Client{}
	dt := &PayloadCardGet{}
	dt.PreSelectedItemList = append(dt.PreSelectedItemList)
	dt.Version = 3
	res1B, _ := json.Marshal(dt)
	// fmt.Println()
	// fmt.Println("PAYLOAD cardGet")
	// fmt.Println()
	// fmt.Println(string(res1B))
	req, err := http.NewRequest("POST", "https://shopee.co.id/api/v4/cart/get", bytes.NewBuffer(res1B))
	if err != nil {
		fmt.Println("No response from request")
	}
	req.Header = http.Header{
		"Accept":        []string{"application/json"},
		"Content-Type":  []string{"application/json"},
		"Referer":       []string{"https://shopee.co.id/checkout"},
		"User-Agent":    []string{os.Getenv("USERAGENT")},
		"Cookie":        []string{os.Getenv("COOKIE")},
		"X-Csrftoken":   []string{c.Csrf},
		"X-Api-Source":  []string{"pc"},
		"if-none-match": []string{"*"},
	}
	// fmt.Println(req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("No response from request")
	}
	body, _ := ioutil.ReadAll(res.Body) // response body is []byte
	// fmt.Println()
	// fmt.Println("response cardget")
	// fmt.Println()
	// fmt.Println(string(body))

	var h ResponseCardGet
	if err := json.Unmarshal(body, &h); err != nil { // Parse []byte to the go struct pointer
		fmt.Println(err)
	}
	// fmt.Println(h)
	// var data1 map[string]interface{}
	// json.Unmarshal(body, &data1)
	// fmt.Println(data1["checkout_price_data"])

	return h

}

func (c *Checkout) GoCheckOut(modelid int64) {
	if os.Getenv("BELANJA") != "sukses" {
		cd := c.card_get()
		geco_first := c.GetCheckout(modelid, cd)
		// fmt.Println(geco_first)
		geco_end := c.getCheckoutFinal(geco_first, modelid)
		// fmt.Println(geco_end)
		dt := c.place_holder(geco_end)
		if dt["checkoutid"] != 0 || dt["checkoutid"] != nil {
			f, err := os.OpenFile(".env", os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			if _, err = f.WriteString("BELANJA=sukses\n"); err != nil {
				panic(err)
			}
			fmt.Println("✅Barang sudah CHECKOUT✅")
		}
	}
}
