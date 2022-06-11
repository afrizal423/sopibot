package checkout

type PayloadCardGet struct {
	PreSelectedItemList []interface{} `json:"pre_selected_item_list"`
	Version             int           `json:"version"`
}
type ResponseCardGet struct {
	Error        int    `json:"error"`
	ErrorMessage string `json:"error_message"`
	Data         struct {
		AllPromotionRules interface{} `json:"all_promotion_rules"`
		ShopOrders        []struct {
			ShopHasVoucher interface{} `json:"shop_has_voucher"`
			Shop           struct {
				Shopname              string      `json:"shopname"`
				HolidayModeOn         bool        `json:"holiday_mode_on"`
				Shopid                int         `json:"shopid"`
				Username              string      `json:"username"`
				Status                int         `json:"status"`
				CbOption              int         `json:"cb_option"`
				ShowOfficialShopLabel bool        `json:"show_official_shop_label"`
				Portrait              string      `json:"portrait"`
				FollowingCount        int         `json:"following_count"`
				Userid                int         `json:"userid"`
				IsFreeShipping        bool        `json:"is_free_shipping"`
				IsShopeeVerified      bool        `json:"is_shopee_verified"`
				HasWelcomePackageItem bool        `json:"has_welcome_package_item"`
				EnabledChannelids     interface{} `json:"enabled_channelids"`
				PromotionRules        interface{} `json:"promotion_rules"`
				ShopTag               int         `json:"shop_tag"`
				AddinTime             int         `json:"addin_time"`
				ClickTime             int         `json:"click_time"`
				EntranceText          []struct {
					Description string `json:"description"`
					Attributes  []struct {
						Key            string `json:"key"`
						Value          int    `json:"value"`
						FormatCurrency bool   `json:"format_currency"`
					} `json:"attributes"`
					RuleID int `json:"rule_id"`
				} `json:"entrance_text"`
				DrawerEntries []struct {
					MinSpend                  int           `json:"min_spend"`
					ShippingDiscount          int           `json:"shipping_discount"`
					ShippingOptions           []string      `json:"shipping_options"`
					Labels                    []interface{} `json:"labels"`
					ShippingDiscountFormatted interface{}   `json:"shipping_discount_formatted"`
					RuleIds                   []int         `json:"rule_ids"`
				} `json:"drawer_entries"`
			} `json:"shop"`
			Items []struct {
				AddOnDealID   int `json:"add_on_deal_id"`
				AddOnDealInfo struct {
					AddOnDealID      int    `json:"add_on_deal_id"`
					Shopid           int    `json:"shopid"`
					StartTime        int    `json:"start_time"`
					EndTime          int    `json:"end_time"`
					CreateTime       int    `json:"create_time"`
					UpdateTime       int    `json:"update_time"`
					Status           int    `json:"status"`
					Region           string `json:"region"`
					SubItemLimit     int    `json:"sub_item_limit"`
					SubType          int    `json:"sub_type"`
					ImageHash        string `json:"image_hash"`
					PurchaseMinSpend int    `json:"purchase_min_spend"`
					PerGiftNum       int    `json:"per_gift_num"`
					TotalGiftTier    int    `json:"total_gift_tier"`
				} `json:"add_on_deal_info"`
				AddinTime                       int           `json:"addin_time"`
				BundleDealID                    interface{}   `json:"bundle_deal_id"`
				BundleDealInfo                  interface{}   `json:"bundle_deal_info"`
				CanUseWholesale                 bool          `json:"can_use_wholesale"`
				Catid                           int           `json:"catid"`
				CbOption                        int           `json:"cb_option"`
				Condition                       int           `json:"condition"`
				FreeReturnDay                   int           `json:"free_return_day"`
				Image                           string        `json:"image"`
				IsAddOnSubItem                  bool          `json:"is_add_on_sub_item"`
				IsWpItem                        bool          `json:"is_wp_item"`
				ItemGroupID                     string        `json:"item_group_id"`
				ItemStock                       int           `json:"item_stock"`
				Itemid                          int64         `json:"itemid"`
				ModelName                       string        `json:"model_name"`
				Modelid                         int64         `json:"modelid"`
				Models                          []interface{} `json:"models"`
				MsgUpdatedNewValue              int           `json:"msg_updated_new_value"`
				MsgUpdatedOldValue              int           `json:"msg_updated_old_value"`
				Name                            string        `json:"name"`
				Oldprice                        int64         `json:"oldprice"`
				OriginCartItemPrice             int64         `json:"origin_cart_item_price"`
				Price                           int64         `json:"price"`
				Quantity                        int           `json:"quantity"`
				Shopid                          int           `json:"shopid"`
				Status                          int           `json:"status"`
				Stock                           int           `json:"stock"`
				UpdatedField                    string        `json:"updated_field"`
				WholesaleTierList               []interface{} `json:"wholesale_tier_list"`
				CurrentPromotionReservedStock   int           `json:"current_promotion_reserved_stock"`
				NormalStock                     int           `json:"normal_stock"`
				CurrentPromotionHasReserveStock bool          `json:"current_promotion_has_reserve_stock"`
				IsFreeGift                      bool          `json:"is_free_gift"`
				TierVariations                  []interface{} `json:"tier_variations"`
				PriceBeforeDiscount             interface{}   `json:"price_before_discount"`
				Flag                            int           `json:"flag"`
				WelcomePackageType              int           `json:"welcome_package_type"`
				CanUseBundleDeal                bool          `json:"can_use_bundle_deal"`
				Offerid                         interface{}   `json:"offerid"`
				Currency                        string        `json:"currency"`
				PromotionType                   int           `json:"promotion_type"`
				DeepDiscountCampaignPeriod      interface{}   `json:"deep_discount_campaign_period"`
				DeepDiscountCampaignTextPc      interface{}   `json:"deep_discount_campaign_text_pc"`
				DeepDiscountCampaignTextMobile  interface{}   `json:"deep_discount_campaign_text_mobile"`
				DeepDiscountCampaignStartTime   interface{}   `json:"deep_discount_campaign_start_time"`
				IsInPreview                     bool          `json:"is_in_preview"`
				ExclusivePriceLabel             interface{}   `json:"exclusive_price_label"`
				IsDefaultModelItem              bool          `json:"is_default_model_item"`
				LogisticsInfo                   interface{}   `json:"logistics_info"`
				CartItemChangeTime              int           `json:"cart_item_change_time"`
				ItemDiscountReal                interface{}   `json:"item_discount_real"`
				IsPreOrder                      bool          `json:"is_pre_order"`
				FulfillmentSource               interface{}   `json:"fulfillment_source"`
				AppliedPromotionID              int           `json:"applied_promotion_id"`
				ProductPromotionType            int           `json:"product_promotion_type"`
				ProductPromotionID              int           `json:"product_promotion_id"`
				IsWholesalePrice                bool          `json:"is_wholesale_price"`
				Checkout                        interface{}   `json:"checkout"`
				NeedRefreshUpdateItem           bool          `json:"need_refresh_update_item"`
				TotalCanBuyQuantity             int           `json:"total_can_buy_quantity"`
				ExtraDiscountLabel              interface{}   `json:"extra_discount_label"`
				IsStreamingPrice                interface{}   `json:"is_streaming_price"`
				CampaignImage                   string        `json:"campaign_image"`
				CartItemType                    int           `json:"cart_item_type"`
				SubItemRawQuantity              interface{}   `json:"sub_item_raw_quantity"`
				PackageType                     int           `json:"package_type"`
				MembershipOfferID               interface{}   `json:"membership_offer_id"`
				IsAhoraItem                     bool          `json:"is_ahora_item"`
				IsSpaylaterItem                 interface{}   `json:"is_spaylater_item"`
			} `json:"items"`
		} `json:"shop_orders"`
		Fsv struct {
			FsvMessage string `json:"fsv_message"`
			URL        string `json:"url"`
			Error      int    `json:"error"`
			ErrorMsg   string `json:"error_msg"`
		} `json:"fsv"`
		ShopOrderIds []struct {
			Shopid     int `json:"shopid"`
			ItemBriefs []struct {
				Itemid      int64 `json:"itemid"`
				Modelid     int64 `json:"modelid"`
				ItemGroupID int64 `json:"item_group_id"`
				Quantity    int   `json:"quantity"`
			} `json:"item_briefs"`
		} `json:"shop_order_ids"`
		LogisticsChannels              interface{} `json:"logistics_channels"`
		AddOnDeals                     interface{} `json:"add_on_deals"`
		IsFreeShippingVoucherToggledOn bool        `json:"is_free_shipping_voucher_toggled_on"`
		CoinMinRedeem                  int         `json:"coin_min_redeem"`
		BuyerLocationGroupIds          interface{} `json:"buyer_location_group_ids"`
		ShopOrderIDList                []struct {
			Shopid     int `json:"shopid"`
			ItemBriefs []struct {
				Itemid             int64       `json:"itemid"`
				Modelid            int64       `json:"modelid"`
				ItemGroupID        string      `json:"item_group_id"`
				Status             interface{} `json:"status"`
				Offerid            interface{} `json:"offerid"`
				Price              interface{} `json:"price"`
				Quantity           int         `json:"quantity"`
				IsAddOnSubItem     interface{} `json:"is_add_on_sub_item"`
				AddOnDealID        interface{} `json:"add_on_deal_id"`
				CartItemChangeTime interface{} `json:"cart_item_change_time"`
				IsWholesaleItem    interface{} `json:"is_wholesale_item"`
				AppliedPromotionID interface{} `json:"applied_promotion_id"`
				MembershipOfferID  interface{} `json:"membership_offer_id"`
			} `json:"item_briefs"`
			AddinTime int         `json:"addin_time"`
			ClickTime interface{} `json:"click_time"`
		} `json:"shop_order_id_list"`
		DiscountTabShopOrderIds struct {
			Discount []interface{} `json:"discount"`
			Other    []interface{} `json:"other"`
		} `json:"discount_tab_shop_order_ids"`
		DisableTab        bool          `json:"disable_tab"`
		ShowTabs          []interface{} `json:"show_tabs"`
		RnToggle          int           `json:"rn_toggle"`
		PackagedItemInfos []struct {
			AvailSubAmount int         `json:"avail_sub_amount"`
			PackageID      string      `json:"package_id"`
			MeetMinSpend   bool        `json:"meet_min_spend"`
			ShowPromoBar   bool        `json:"show_promo_bar"`
			PurchaseType   interface{} `json:"purchase_type"`
			PackageType    int         `json:"package_type"`
			NextTier       interface{} `json:"next_tier"`
		} `json:"packaged_item_infos"`
	} `json:"data"`
}
