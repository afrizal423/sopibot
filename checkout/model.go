package checkout

type ResponseCheckOutGetFinal struct {
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
		TotalPayable                    int64       `json:"total_payable"`
	} `json:"checkout_price_data"`
	OrderUpdateInfo struct {
	} `json:"order_update_info"`
	DropshippingInfo struct {
		Enabled     bool   `json:"enabled"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	} `json:"dropshipping_info"`
	PromotionData struct {
		CanUseCoins             bool          `json:"can_use_coins"`
		UseCoins                bool          `json:"use_coins"`
		PlatformVouchers        []interface{} `json:"platform_vouchers"`
		FreeShippingVoucherInfo struct {
			FreeShippingVoucherID   int         `json:"free_shipping_voucher_id"`
			FreeShippingVoucherCode string      `json:"free_shipping_voucher_code"`
			DisabledReason          interface{} `json:"disabled_reason"`
			BannerInfo              struct {
				Msg          string `json:"msg"`
				LearnMoreMsg string `json:"learn_more_msg"`
			} `json:"banner_info"`
		} `json:"free_shipping_voucher_info"`
		ShopVoucherEntrances []struct {
			Shopid int  `json:"shopid"`
			Status bool `json:"status"`
		} `json:"shop_voucher_entrances"`
		AppliedVoucherCode interface{} `json:"applied_voucher_code"`
		VoucherCode        interface{} `json:"voucher_code"`
		VoucherInfo        struct {
			CoinEarned         int         `json:"coin_earned"`
			VoucherCode        interface{} `json:"voucher_code"`
			CoinPercentage     int         `json:"coin_percentage"`
			DiscountPercentage int         `json:"discount_percentage"`
			DiscountValue      int         `json:"discount_value"`
			Promotionid        int         `json:"promotionid"`
			RewardType         int         `json:"reward_type"`
			UsedPrice          int         `json:"used_price"`
		} `json:"voucher_info"`
		InvalidMessage string `json:"invalid_message"`
		PriceDiscount  int    `json:"price_discount"`
		CoinInfo       struct {
			CoinOffset        int `json:"coin_offset"`
			CoinUsed          int `json:"coin_used"`
			CoinEarnByVoucher int `json:"coin_earn_by_voucher"`
			CoinEarn          int `json:"coin_earn"`
		} `json:"coin_info"`
		CardPromotionID      interface{} `json:"card_promotion_id"`
		CardPromotionEnabled bool        `json:"card_promotion_enabled"`
		PromotionMsg         string      `json:"promotion_msg"`
	} `json:"promotion_data"`
	SelectedPaymentChannelData struct {
		Version               int    `json:"version"`
		OptionInfo            string `json:"option_info"`
		ChannelID             int    `json:"channel_id"`
		ChannelItemOptionInfo struct {
			OptionInfo string `json:"option_info"`
		} `json:"channel_item_option_info"`
		AdditionalInfo struct {
			Reason          string `json:"reason"`
			ChannelBlackbox string `json:"channel_blackbox"`
		} `json:"additional_info"`
		TextInfo struct {
		} `json:"text_info"`
	} `json:"selected_payment_channel_data"`
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
			Itemid      int64       `json:"itemid"`
			Modelid     int64       `json:"modelid"`
			Quantity    int         `json:"quantity"`
			ItemGroupID interface{} `json:"item_group_id"`
			Insurances  []struct {
				InsuranceProductID    string `json:"insurance_product_id"`
				Name                  string `json:"name"`
				Description           string `json:"description"`
				ProductDetailPageURL  string `json:"product_detail_page_url"`
				Premium               int    `json:"premium"`
				PremiumBeforeDiscount int    `json:"premium_before_discount"`
				InsuranceQuantity     int    `json:"insurance_quantity"`
				Selected              bool   `json:"selected"`
			} `json:"insurances"`
			Shopid                  int    `json:"shopid"`
			Shippable               bool   `json:"shippable"`
			NonShippableErr         string `json:"non_shippable_err"`
			NoneShippableReason     string `json:"none_shippable_reason"`
			NoneShippableFullReason string `json:"none_shippable_full_reason"`
			Price                   int    `json:"price"`
			Name                    string `json:"name"`
			ModelName               string `json:"model_name"`
			AddOnDealID             int    `json:"add_on_deal_id"`
			IsAddOnSubItem          bool   `json:"is_add_on_sub_item"`
			IsPreOrder              bool   `json:"is_pre_order"`
			IsStreamingPrice        bool   `json:"is_streaming_price"`
			Image                   string `json:"image"`
			Checkout                bool   `json:"checkout"`
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
		OrderTotal                int64       `json:"order_total"`
		BuyerRemark               interface{} `json:"buyer_remark"`
	} `json:"shoporders"`
	ShippingOrders []struct {
		ShippingID                            int         `json:"shipping_id"`
		ShoporderIndexes                      []int       `json:"shoporder_indexes"`
		SelectedLogisticChannelid             int         `json:"selected_logistic_channelid"`
		SelectedPreferredDeliveryTimeOptionID int         `json:"selected_preferred_delivery_time_option_id"`
		BuyerRemark                           interface{} `json:"buyer_remark"`
		BuyerAddressData                      struct {
			Addressid   int    `json:"addressid"`
			AddressType int    `json:"address_type"`
			TaxAddress  string `json:"tax_address"`
		} `json:"buyer_address_data"`
		FulfillmentInfo struct {
			FulfillmentFlag      int    `json:"fulfillment_flag"`
			FulfillmentSource    string `json:"fulfillment_source"`
			ManagedBySbs         bool   `json:"managed_by_sbs"`
			OrderFulfillmentType int    `json:"order_fulfillment_type"`
			WarehouseAddressID   int    `json:"warehouse_address_id"`
			IsFromOverseas       bool   `json:"is_from_overseas"`
		} `json:"fulfillment_info"`
		Logistics struct {
			IntegratedChannelids            []int         `json:"integrated_channelids"`
			NonIntegratedChannelids         []interface{} `json:"non_integrated_channelids"`
			VoucherWalletCheckingChannelIds []int         `json:"voucher_wallet_checking_channel_ids"`
			LogisticChannels                struct {
				Num8003 struct {
					ChannelData struct {
						AddressType   int    `json:"address_type"`
						Channelid     int    `json:"channelid"`
						CodSupported  bool   `json:"cod_supported"`
						Enabled       bool   `json:"enabled"`
						IsMaskChannel int    `json:"is_mask_channel"`
						Name          string `json:"name"`
						Priority      int    `json:"priority"`
						Warning       string `json:"warning"`
						WarningMsg    string `json:"warning_msg"`
					} `json:"channel_data"`
					CodData struct {
						CodAvailable bool `json:"cod_available"`
					} `json:"cod_data"`
					DeliveryData struct {
						DelayMessage string `json:"delay_message"`
						DetailInfo   struct {
							Apt      float64 `json:"apt"`
							CdtMax   float64 `json:"cdt_max"`
							CdtMin   float64 `json:"cdt_min"`
							EdtMaxDt string  `json:"edt_max_dt"`
							EdtMinDt string  `json:"edt_min_dt"`
							HeCdt    int     `json:"he_cdt"`
							HePt     int     `json:"he_pt"`
						} `json:"detail_info"`
						DisplayMode               string `json:"display_mode"`
						MaxDays                   int    `json:"max_days"`
						MinDays                   int    `json:"min_days"`
						EstimatedDeliveryTimeMax  int    `json:"estimated_delivery_time_max"`
						EstimatedDeliveryTimeMin  int    `json:"estimated_delivery_time_min"`
						EstimatedDeliveryDateFrom int    `json:"estimated_delivery_date_from"`
						EstimatedDeliveryDateTo   int    `json:"estimated_delivery_date_to"`
						HasEdt                    bool   `json:"has_edt"`
						IsCrossBorder             bool   `json:"is_cross_border"`
						IsRapidSLA                bool   `json:"is_rapid_sla"`
						IsShopee24H               bool   `json:"is_shopee_24h"`
					} `json:"delivery_data"`
					ShippingFeeData struct {
						ChargeableShippingFee     int `json:"chargeable_shipping_fee"`
						ShippingFeeBeforeDiscount int `json:"shipping_fee_before_discount"`
					} `json:"shipping_fee_data"`
				} `json:"8003"`
				Num8006 struct {
					ChannelData struct {
						AddressType   int    `json:"address_type"`
						Channelid     int    `json:"channelid"`
						CodSupported  bool   `json:"cod_supported"`
						Enabled       bool   `json:"enabled"`
						IsMaskChannel int    `json:"is_mask_channel"`
						Name          string `json:"name"`
						Priority      int    `json:"priority"`
						Warning       string `json:"warning"`
						WarningMsg    string `json:"warning_msg"`
					} `json:"channel_data"`
					CodData struct {
						CodAvailable bool `json:"cod_available"`
					} `json:"cod_data"`
					DeliveryData struct {
						DelayMessage string `json:"delay_message"`
						DetailInfo   struct {
							Apt      float64 `json:"apt"`
							CdtMax   float64 `json:"cdt_max"`
							CdtMin   float64 `json:"cdt_min"`
							EdtMaxDt string  `json:"edt_max_dt"`
							EdtMinDt string  `json:"edt_min_dt"`
							HeCdt    int     `json:"he_cdt"`
							HePt     int     `json:"he_pt"`
						} `json:"detail_info"`
						DisplayMode               string `json:"display_mode"`
						MaxDays                   int    `json:"max_days"`
						MinDays                   int    `json:"min_days"`
						EstimatedDeliveryTimeMax  int    `json:"estimated_delivery_time_max"`
						EstimatedDeliveryTimeMin  int    `json:"estimated_delivery_time_min"`
						EstimatedDeliveryDateFrom int    `json:"estimated_delivery_date_from"`
						EstimatedDeliveryDateTo   int    `json:"estimated_delivery_date_to"`
						HasEdt                    bool   `json:"has_edt"`
						IsCrossBorder             bool   `json:"is_cross_border"`
						IsRapidSLA                bool   `json:"is_rapid_sla"`
						IsShopee24H               bool   `json:"is_shopee_24h"`
					} `json:"delivery_data"`
					ShippingFeeData struct {
						ChargeableShippingFee     int64 `json:"chargeable_shipping_fee"`
						ShippingFeeBeforeDiscount int64 `json:"shipping_fee_before_discount"`
					} `json:"shipping_fee_data"`
				} `json:"8006"`
			} `json:"logistic_channels"`
			LogisticServiceTypes struct {
				Regular struct {
					ChannelIds []int  `json:"channel_ids"`
					Enabled    bool   `json:"enabled"`
					Identifier string `json:"identifier"`
					MaxCost    int    `json:"max_cost"`
					MinCost    int    `json:"min_cost"`
					Name       string `json:"name"`
					Priority   int    `json:"priority"`
					SLAMsg     string `json:"sla_msg"`
				} `json:"regular"`
				RegularCargo struct {
					ChannelIds []int  `json:"channel_ids"`
					Enabled    bool   `json:"enabled"`
					Identifier string `json:"identifier"`
					MaxCost    int64  `json:"max_cost"`
					MinCost    int64  `json:"min_cost"`
					Name       string `json:"name"`
					Priority   int    `json:"priority"`
					SLAMsg     string `json:"sla_msg"`
				} `json:"regular_cargo"`
			} `json:"logistic_service_types"`
		} `json:"logistics"`
		OrderTotal                           int         `json:"order_total"`
		OrderTotalWithoutShipping            int         `json:"order_total_without_shipping"`
		SelectedLogisticChannelidWithWarning interface{} `json:"selected_logistic_channelid_with_warning"`
		ShippingFee                          int         `json:"shipping_fee"`
		ShippingFeeDiscount                  int         `json:"shipping_fee_discount"`
		ShippingGroupDescription             string      `json:"shipping_group_description"`
		ShippingGroupIcon                    string      `json:"shipping_group_icon"`
		TaxPayable                           int         `json:"tax_payable"`
		IsFsvApplied                         bool        `json:"is_fsv_applied"`
	} `json:"shipping_orders"`
	FsvSelectionInfos []interface{} `json:"fsv_selection_infos"`
	BuyerInfo         struct {
		ShareToFriendsInfo struct {
			DisplayToggle bool `json:"display_toggle"`
			EnableToggle  bool `json:"enable_toggle"`
			AllowToShare  bool `json:"allow_to_share"`
		} `json:"share_to_friends_info"`
		KycInfo       interface{} `json:"kyc_info"`
		CheckoutEmail string      `json:"checkout_email"`
	} `json:"buyer_info"`
	PaymentChannelInfo struct {
		Channels []struct {
			Flag                int64 `json:"flag"`
			SpmChannelID        int   `json:"spm_channel_id"`
			Balance             int   `json:"balance,omitempty"`
			PreselectDisabled   bool  `json:"preselect_disabled"`
			SelectedChannelItem struct {
				ChannelItemID   int `json:"channel_item_id"`
				ChannelItemData struct {
					LogoURL         string `json:"logo_url"`
					BankName        string `json:"bank_name"`
					AgencyShortname string `json:"agency_shortname"`
					Identifier      string `json:"identifier"`
					BankID          int    `json:"bank_id"`
					CardNumber      string `json:"card_number"`
					ExpiryDate      string `json:"expiry_date"`
					AgencyUen       string `json:"agency_uen"`
				} `json:"channel_item_data"`
				ErrorCode    int    `json:"error_code"`
				ErrorMessage string `json:"error_message"`
			} `json:"selected_channel_item"`
			ChannelBlackbox string `json:"channel_blackbox"`
			Groups          struct {
				SpmChannel   bool `json:"spm_channel"`
				Cod          bool `json:"cod"`
				BankTransfer bool `json:"bank_transfer"`
				Wallet       bool `json:"wallet"`
				Installment  bool `json:"installment"`
				GeneralCard  bool `json:"general_card"`
				Airpay       bool `json:"airpay"`
				Seabank      bool `json:"seabank"`
				Spl          bool `json:"spl"`
				Immediate    bool `json:"immediate"`
				CreditCard   bool `json:"credit_card"`
				BankAccount  bool `json:"bank_account"`
			} `json:"groups"`
			OptionKeysName          []interface{} `json:"option_keys_name"`
			ChannelID               int           `json:"channel_id,omitempty"`
			PinResetNeeded          bool          `json:"pin_reset_needed,omitempty"`
			HasWallet               bool          `json:"has_wallet"`
			Currency                string        `json:"currency"`
			DisabledReasonKey       int           `json:"disabled_reason_key"`
			DisabledReasonLearnMore string        `json:"disabled_reason_learn_more"`
			TopupNeeded             bool          `json:"topup_needed,omitempty"`
			Name                    string        `json:"name"`
			InfoText                string        `json:"info_text"`
			ActivationNeeded        bool          `json:"activation_needed,omitempty"`
			SpmOptionInfo           string        `json:"spm_option_info"`
			Country                 string        `json:"country"`
			ExtraData               struct {
				IncludeShippingAndDiscounts bool          `json:"include_shipping_and_discounts"`
				VoucherPaymentType          int           `json:"voucher_payment_type"`
				AlternateChannelIds         []interface{} `json:"alternate_channel_ids"`
				BannedCategories            []interface{} `json:"banned_categories"`
				ExpiryDuration              int           `json:"expiry_duration"`
				EligibleTransfer            string        `json:"eligible_transfer"`
				V1ChannelGroupingInfo       struct {
					ChannelName         string        `json:"channel_name"`
					IconPath            string        `json:"icon_path"`
					IconBackground      string        `json:"icon_background"`
					GroupID             int           `json:"group_id"`
					Version             int           `json:"version"`
					ChannelIdsToCombine []interface{} `json:"channel_ids_to_combine"`
				} `json:"v1_channel_grouping_info"`
				MaxCheckoutPrice int    `json:"max_checkout_price"`
				SubDescription1  string `json:"sub_description_1"`
				ExpiryExtension  int    `json:"expiry_extension"`
				PriceLimit       int    `json:"price_limit"`
				MinAmount        int    `json:"min_amount"`
				SubDescription2  string `json:"sub_description_2"`
			} `json:"extra_data"`
			KycEnforcementNeeded bool   `json:"kyc_enforcement_needed,omitempty"`
			WalletBalance        int    `json:"wallet_balance"`
			AlwaysShow           bool   `json:"always_show"`
			Enabled              bool   `json:"enabled"`
			PaymentResultText    string `json:"payment_result_text"`
			BeChannelID          int    `json:"be_channel_id"`
			DisabledReasonTitle  string `json:"disabled_reason_title"`
			ChannelBehavior      struct {
				DisableInstruction bool `json:"disable_instruction"`
				DisableCancel      bool `json:"disable_cancel"`
			} `json:"channel_behavior,omitempty"`
			Version            int    `json:"version"`
			IsNew              bool   `json:"is_new"`
			InfoLink           string `json:"info_link"`
			ChannelProvider    string `json:"channel_provider"`
			IconBackground     string `json:"icon_background"`
			SubDescriptionInfo struct {
				First struct {
					Text      string `json:"text"`
					Highlight bool   `json:"highlight"`
				} `json:"first"`
				Second struct {
					Text      string `json:"text"`
					Highlight bool   `json:"highlight"`
				} `json:"second"`
			} `json:"sub_description_info"`
			Category                          int           `json:"category"`
			Description                       string        `json:"description"`
			PaymentChannelBanners             []interface{} `json:"payment_channel_banners"`
			NameLabel                         string        `json:"name_label"`
			DisabledReason                    string        `json:"disabled_reason"`
			DisabledReasonLearnMoreButtonText string        `json:"disabled_reason_learn_more_button_text"`
			PromotionInfo                     struct {
				VoucherPaymentType []interface{} `json:"voucher_payment_type"`
				Text               string        `json:"text"`
			} `json:"promotion_info"`
			PinSetUpNeeded    bool   `json:"pin_set_up_needed,omitempty"`
			Subcategory       int    `json:"subcategory"`
			Priority          int    `json:"priority"`
			Icon              string `json:"icon"`
			Channelid         int    `json:"channelid,omitempty"`
			ChannelBehaviorss struct {
			} `json:"channel_behavior,omitempty"`
			Banks []struct {
				OptionInfo         string `json:"option_info"`
				Description        string `json:"description"`
				BankName           string `json:"bank_name"`
				DisabledReason     string `json:"disabled_reason"`
				Enabled            bool   `json:"enabled"`
				Icon               string `json:"icon"`
				SubDescriptionInfo struct {
					Normal    string `json:"normal"`
					Important string `json:"important"`
				} `json:"sub_description_info"`
				DeepLink              string `json:"deep_link"`
				PopupConfirmationData struct {
					NeedPopup           bool   `json:"need_popup"`
					PopupMessage        string `json:"popup_message"`
					OkButtonMessage     string `json:"ok_button_message"`
					CancelButtonMessage string `json:"cancel_button_message"`
				} `json:"popup_confirmation_data"`
				Hide           bool   `json:"hide"`
				IconLabel      string `json:"icon_label"`
				IconBackground string `json:"icon_background"`
				HandlingFee    int    `json:"handling_fee"`
				BeChannelID    int    `json:"be_channel_id"`
			} `json:"banks,omitempty"`
			PaymentChannelHints []interface{} `json:"payment_channel_hints,omitempty"`
			Promotions          []struct {
				DiscountType    int    `json:"discount_type"`
				BankID          int    `json:"bank_id"`
				CardPromotionID int    `json:"card_promotion_id"`
				PrimaryColor    string `json:"primary_color"`
				IsFullyRedeemed bool   `json:"is_fully_redeemed"`
				EndTime         int    `json:"end_time"`
				DiscountValue   string `json:"discount_value"`
				URL             string `json:"url"`
				Title           string `json:"title"`
				BankLogo        string `json:"bank_logo"`
				StartTime       int    `json:"start_time"`
				Description     string `json:"description"`
			} `json:"promotions,omitempty"`
			Cards      []interface{} `json:"cards,omitempty"`
			AddCardURL struct {
				API string `json:"api"`
				UI  string `json:"ui"`
			} `json:"add_card_url,omitempty"`
			InstallmentBanks []interface{} `json:"installment_banks,omitempty"`
			AddCardURLs      struct {
				UI string `json:"ui"`
			} `json:"add_card_url,omitempty"`
			AddCardButtonLabel string `json:"add_card_button_label,omitempty"`
			HideAddCardButton  bool   `json:"hide_add_card_button,omitempty"`
			AddCardURLss       struct {
				UI string `json:"ui"`
			} `json:"add_card_url,omitempty"`
			ChannelBehaviors struct {
			} `json:"channel_behavior,omitempty"`
			Subchannels []struct {
				InstallmentPlans         []interface{} `json:"installment_plans"`
				InstallmentPromotionText string        `json:"installment_promotion_text"`
				WalletNoPin              bool          `json:"wallet_no_pin"`
				AlwaysShow               bool          `json:"always_show"`
				BlockedByPlp             bool          `json:"blocked_by_plp"`
				PaymentResultText        string        `json:"payment_result_text"`
				DisabledReason           string        `json:"disabled_reason"`
				ExtraData                struct {
					PriceLimit            int           `json:"price_limit"`
					BannedCategories      []interface{} `json:"banned_categories"`
					SubDescription2       string        `json:"sub_description_2"`
					ExpiryDuration        int           `json:"expiry_duration"`
					EligibleTransfer      string        `json:"eligible_transfer"`
					AlternateChannelIds   []interface{} `json:"alternate_channel_ids"`
					MaxCheckoutPrice      int64         `json:"max_checkout_price"`
					VoucherPaymentType    int           `json:"voucher_payment_type"`
					V1ChannelGroupingInfo struct {
						ChannelIdsToCombine []interface{} `json:"channel_ids_to_combine"`
						ChannelName         string        `json:"channel_name"`
						IconPath            string        `json:"icon_path"`
						IconBackground      string        `json:"icon_background"`
						GroupID             int           `json:"group_id"`
						Version             int           `json:"version"`
					} `json:"v1_channel_grouping_info"`
					ExpiryExtension             int    `json:"expiry_extension"`
					IncludeShippingAndDiscounts bool   `json:"include_shipping_and_discounts"`
					SubDescription1             string `json:"sub_description_1"`
					MinAmount                   int    `json:"min_amount"`
				} `json:"extra_data"`
				NameLabel          string        `json:"name_label"`
				TopupNeeded        bool          `json:"topup_needed"`
				BankAccounts       []interface{} `json:"bank_accounts"`
				Subchannels        []interface{} `json:"subchannels"`
				CcbText            string        `json:"ccb_text"`
				DisabledReasonKey  int           `json:"disabled_reason_key"`
				Subcategory        int           `json:"subcategory"`
				Enabled            bool          `json:"enabled"`
				DefaultBankAccount struct {
					Version               int `json:"version"`
					ChannelID             int `json:"channel_id"`
					ChannelItemOptionInfo struct {
						ChannelItemID         int    `json:"channel_item_id"`
						ExternalChannelItemID int    `json:"external_channel_item_id"`
						OptionInfo            string `json:"option_info"`
					} `json:"channel_item_option_info"`
				} `json:"default_bank_account"`
				CcbLink            string `json:"ccb_link"`
				Channelid          int    `json:"channelid"`
				IsNew              bool   `json:"is_new"`
				Description        string `json:"description"`
				WalletBalance      int    `json:"wallet_balance"`
				IsReadyToPayConfig struct {
					ResponseSuccessCases []interface{} `json:"response_success_cases"`
					DisableText          string        `json:"disable_text"`
					Request              string        `json:"request"`
				} `json:"is_ready_to_pay_config"`
				BeChannelID                       int           `json:"be_channel_id"`
				Priority                          int           `json:"priority"`
				DisabledReasonTitle               string        `json:"disabled_reason_title"`
				ChannelID                         int           `json:"channel_id"`
				DisabledReasonLearnMoreButtonText string        `json:"disabled_reason_learn_more_button_text"`
				Cards                             []interface{} `json:"cards"`
				BillingAddresses                  []interface{} `json:"billing_addresses"`
				PaymentReasonInfo                 struct {
					NeedReason      bool `json:"need_reason"`
					ReasonMaxLength int  `json:"reason_max_length"`
				} `json:"payment_reason_info"`
				PinSetUpNeeded      bool   `json:"pin_set_up_needed"`
				HideAddCardButton   bool   `json:"hide_add_card_button"`
				Version             int    `json:"version"`
				Country             string `json:"country"`
				Currency            string `json:"currency"`
				NeedTopupAmount     int    `json:"need_topup_amount"`
				PinResetNeeded      bool   `json:"pin_reset_needed"`
				LinkBankAccountInfo struct {
					LinkIcon string `json:"link_icon"`
					LinkText string `json:"link_text"`
				} `json:"link_bank_account_info"`
				SelectedChannelItem struct {
					ErrorMessage    string `json:"error_message"`
					ChannelItemID   int    `json:"channel_item_id"`
					ChannelItemData struct {
						LogoURL         string `json:"logo_url"`
						BankID          int    `json:"bank_id"`
						BankName        string `json:"bank_name"`
						ExpiryDate      string `json:"expiry_date"`
						AgencyShortname string `json:"agency_shortname"`
						Identifier      string `json:"identifier"`
						AgencyUen       string `json:"agency_uen"`
						CardNumber      string `json:"card_number"`
					} `json:"channel_item_data"`
					ErrorCode int `json:"error_code"`
				} `json:"selected_channel_item"`
				PromotionInfo struct {
					VoucherPaymentType []interface{} `json:"voucher_payment_type"`
					Text               string        `json:"text"`
				} `json:"promotion_info"`
				Balance             int           `json:"balance"`
				AddBankAccountURL   string        `json:"add_bank_account_url"`
				DebitingAgentList   []interface{} `json:"debiting_agent_list"`
				WhitelistedFields   []string      `json:"whitelisted_fields"`
				IconBackground      string        `json:"icon_background"`
				Name                string        `json:"name"`
				SpmItemsKeyName     string        `json:"spm_items_key_name"`
				Counters            []interface{} `json:"counters"`
				Instructions        []interface{} `json:"instructions"`
				PaymentChannelHints []interface{} `json:"payment_channel_hints"`
				AddBankAccountText  string        `json:"add_bank_account_text"`
				PaymentDisclaimer   struct {
					Body    string `json:"body"`
					URL     string `json:"url"`
					Enabled bool   `json:"enabled"`
					Preview string `json:"preview"`
					Header  string `json:"header"`
				} `json:"payment_disclaimer"`
				ChannelProvider         string        `json:"channel_provider"`
				Flag                    int64         `json:"flag"`
				InfoLink                string        `json:"info_link"`
				DisabledReasonLearnMore string        `json:"disabled_reason_learn_more"`
				KycEnforcementNeeded    bool          `json:"kyc_enforcement_needed"`
				Banks                   []interface{} `json:"banks"`
				Icon                    string        `json:"icon"`
				PartnerAccountInfo      struct {
					RulesSameFlag    bool          `json:"rules_same_flag"`
					Balance          int           `json:"balance"`
					ShouldShowButton bool          `json:"should_show_button"`
					PartnersLink     string        `json:"partners_link"`
					ForceGreyOut     bool          `json:"force_grey_out"`
					UserStatus       int           `json:"user_status"`
					AdditionalInfo   int           `json:"additional_info"`
					IsNewAccount     bool          `json:"is_new_account"`
					DisplayText      string        `json:"display_text"`
					DisplayColor     string        `json:"display_color"`
					TrackingInfo     string        `json:"tracking_info"`
					RulesSkuNames    []interface{} `json:"rules_sku_names"`
				} `json:"partner_account_info"`
				PaymentChannelBanners       []interface{} `json:"payment_channel_banners"`
				SkipPlpCheckoutVerification bool          `json:"skip_plp_checkout_verification"`
				BiInfo                      struct {
					HasLocalSeller bool `json:"has_local_seller"`
					IsKycVerified  bool `json:"is_kyc_verified"`
					HasCbSeller    bool `json:"has_cb_seller"`
				} `json:"bi_info"`
				ParentChannelid   int           `json:"parent_channelid"`
				OptionKeysName    []interface{} `json:"option_keys_name"`
				PreselectDisabled bool          `json:"preselect_disabled"`
				SpmChannelID      int           `json:"spm_channel_id"`
				InfoText          string        `json:"info_text"`
				ChannelBehavior   []interface{} `json:"channel_behavior"`
				ChannelBlackbox   string        `json:"channel_blackbox"`
				Groups            struct {
					Immediate    bool `json:"immediate"`
					CreditCard   bool `json:"credit_card"`
					Wallet       bool `json:"wallet"`
					BankAccount  bool `json:"bank_account"`
					SpmChannel   bool `json:"spm_channel"`
					BankTransfer bool `json:"bank_transfer"`
					Seabank      bool `json:"seabank"`
					Spl          bool `json:"spl"`
					Installment  bool `json:"installment"`
					GeneralCard  bool `json:"general_card"`
					Cod          bool `json:"cod"`
					Airpay       bool `json:"airpay"`
				} `json:"groups"`
				ActivationNeeded              bool          `json:"activation_needed"`
				InstallmentPromotionLink      string        `json:"installment_promotion_link"`
				HasWallet                     bool          `json:"has_wallet"`
				DirectDebitIntegratedBankList []interface{} `json:"direct_debit_integrated_bank_list"`
				AddCardURL                    []interface{} `json:"add_card_url"`
				LinkedBankAccount             struct {
					LinkBankAccountPopup struct {
						CancelButtonMessage string `json:"cancel_button_message"`
						NeedPopup           bool   `json:"need_popup"`
						PopupMessage        string `json:"popup_message"`
						OkButtonMessage     string `json:"ok_button_message"`
					} `json:"link_bank_account_popup"`
					ConsentTokenDisplayNumber string `json:"consent_token_display_number"`
					Default                   bool   `json:"default"`
					BankName                  string `json:"bank_name"`
					VerificationMessage       string `json:"verification_message"`
					BankAccountNumber         string `json:"bank_account_number"`
					ChannelItemID             int    `json:"channel_item_id"`
					Enabled                   bool   `json:"enabled"`
					NeedVerification          bool   `json:"need_verification"`
					IsLinked                  bool   `json:"is_linked"`
					IsExpired                 bool   `json:"is_expired"`
					DisabledReason            string `json:"disabled_reason"`
					DisabledReasonKey         string `json:"disabled_reason_key"`
					Icon                      string `json:"icon"`
					VerificationURL           string `json:"verification_url"`
					LinkBankAccountURL        string `json:"link_bank_account_url"`
				} `json:"linked_bank_account"`
				AlwaysExpanded     bool   `json:"always_expanded"`
				Category           int    `json:"category"`
				SpmOptionInfo      string `json:"spm_option_info"`
				SubDescriptionInfo struct {
					Second struct {
						Text      string `json:"text"`
						Highlight bool   `json:"highlight"`
					} `json:"second"`
					First struct {
						Text      string `json:"text"`
						Highlight bool   `json:"highlight"`
					} `json:"first"`
				} `json:"sub_description_info"`
				Promotions         []interface{} `json:"promotions"`
				InstallmentBanks   []interface{} `json:"installment_banks"`
				AddCardButtonLabel string        `json:"add_card_button_label"`
			} `json:"subchannels,omitempty"`
		} `json:"channels"`
		GroupingInfo struct {
			Groups []struct {
				GroupID     int   `json:"group_id"`
				ToCombine   []int `json:"to_combine"`
				DisplayInfo struct {
					ChannelID      int    `json:"channel_id"`
					IconBackground string `json:"icon_background"`
					Name           string `json:"name"`
					IsNew          bool   `json:"is_new"`
					Icon           string `json:"icon"`
					InfoLink       string `json:"info_link"`
					InfoText       string `json:"info_text"`
					PromotionInfo  struct {
						Text string `json:"text"`
					} `json:"promotion_info"`
					Version int `json:"version"`
				} `json:"display_info"`
			} `json:"groups"`
		} `json:"grouping_info"`
		PromotionInfo struct {
			Text string `json:"text"`
		} `json:"promotion_info"`
	} `json:"payment_channel_info"`
	BuyerTxnFeeInfo struct {
		Title        string `json:"title"`
		Description  string `json:"description"`
		LearnMoreURL string `json:"learn_more_url"`
	} `json:"buyer_txn_fee_info"`
	DisabledCheckoutInfo struct {
		Description string        `json:"description"`
		AutoPopup   bool          `json:"auto_popup"`
		ErrorInfos  []interface{} `json:"error_infos"`
	} `json:"disabled_checkout_info"`
	CanCheckout bool `json:"can_checkout"`
}

type ResponseCheckOutGetTanpaAlfa struct {
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
		TotalPayable                    int64       `json:"total_payable"`
	} `json:"checkout_price_data"`
	OrderUpdateInfo struct {
	} `json:"order_update_info"`
	DropshippingInfo struct {
		Enabled     bool   `json:"enabled"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	} `json:"dropshipping_info"`
	PromotionData struct {
		CanUseCoins             bool          `json:"can_use_coins"`
		UseCoins                bool          `json:"use_coins"`
		PlatformVouchers        []interface{} `json:"platform_vouchers"`
		FreeShippingVoucherInfo struct {
			FreeShippingVoucherID   int         `json:"free_shipping_voucher_id"`
			FreeShippingVoucherCode string      `json:"free_shipping_voucher_code"`
			DisabledReason          interface{} `json:"disabled_reason"`
			BannerInfo              struct {
				Msg          string `json:"msg"`
				LearnMoreMsg string `json:"learn_more_msg"`
			} `json:"banner_info"`
		} `json:"free_shipping_voucher_info"`
		ShopVoucherEntrances []struct {
			Shopid int  `json:"shopid"`
			Status bool `json:"status"`
		} `json:"shop_voucher_entrances"`
		AppliedVoucherCode interface{} `json:"applied_voucher_code"`
		VoucherCode        interface{} `json:"voucher_code"`
		VoucherInfo        struct {
			CoinEarned         int         `json:"coin_earned"`
			VoucherCode        interface{} `json:"voucher_code"`
			CoinPercentage     int         `json:"coin_percentage"`
			DiscountPercentage int         `json:"discount_percentage"`
			DiscountValue      int         `json:"discount_value"`
			Promotionid        int         `json:"promotionid"`
			RewardType         int         `json:"reward_type"`
			UsedPrice          int         `json:"used_price"`
		} `json:"voucher_info"`
		InvalidMessage string `json:"invalid_message"`
		PriceDiscount  int    `json:"price_discount"`
		CoinInfo       struct {
			CoinOffset        int `json:"coin_offset"`
			CoinUsed          int `json:"coin_used"`
			CoinEarnByVoucher int `json:"coin_earn_by_voucher"`
			CoinEarn          int `json:"coin_earn"`
		} `json:"coin_info"`
		CardPromotionID      interface{} `json:"card_promotion_id"`
		CardPromotionEnabled bool        `json:"card_promotion_enabled"`
		PromotionMsg         string      `json:"promotion_msg"`
	} `json:"promotion_data"`
	SelectedPaymentChannelData struct {
		OptionInfo string      `json:"option_info"`
		TextInfo   interface{} `json:"text_info"`
	} `json:"selected_payment_channel_data"`
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
			Itemid      int64       `json:"itemid"`
			Modelid     int64       `json:"modelid"`
			Quantity    int         `json:"quantity"`
			ItemGroupID interface{} `json:"item_group_id"`
			Insurances  []struct {
				InsuranceProductID    string `json:"insurance_product_id"`
				Name                  string `json:"name"`
				Description           string `json:"description"`
				ProductDetailPageURL  string `json:"product_detail_page_url"`
				Premium               int    `json:"premium"`
				PremiumBeforeDiscount int    `json:"premium_before_discount"`
				InsuranceQuantity     int    `json:"insurance_quantity"`
				Selected              bool   `json:"selected"`
			} `json:"insurances"`
			Shopid                  int    `json:"shopid"`
			Shippable               bool   `json:"shippable"`
			NonShippableErr         string `json:"non_shippable_err"`
			NoneShippableReason     string `json:"none_shippable_reason"`
			NoneShippableFullReason string `json:"none_shippable_full_reason"`
			Price                   int    `json:"price"`
			Name                    string `json:"name"`
			ModelName               string `json:"model_name"`
			AddOnDealID             int    `json:"add_on_deal_id"`
			IsAddOnSubItem          bool   `json:"is_add_on_sub_item"`
			IsPreOrder              bool   `json:"is_pre_order"`
			IsStreamingPrice        bool   `json:"is_streaming_price"`
			Image                   string `json:"image"`
			Checkout                bool   `json:"checkout"`
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
	ShippingOrders []struct {
		ShippingID                            int         `json:"shipping_id"`
		ShoporderIndexes                      []int       `json:"shoporder_indexes"`
		SelectedLogisticChannelid             int         `json:"selected_logistic_channelid"`
		SelectedPreferredDeliveryTimeOptionID int         `json:"selected_preferred_delivery_time_option_id"`
		BuyerRemark                           interface{} `json:"buyer_remark"`
		BuyerAddressData                      struct {
			Addressid   int    `json:"addressid"`
			AddressType int    `json:"address_type"`
			TaxAddress  string `json:"tax_address"`
		} `json:"buyer_address_data"`
		FulfillmentInfo struct {
			FulfillmentFlag      int    `json:"fulfillment_flag"`
			FulfillmentSource    string `json:"fulfillment_source"`
			ManagedBySbs         bool   `json:"managed_by_sbs"`
			OrderFulfillmentType int    `json:"order_fulfillment_type"`
			WarehouseAddressID   int    `json:"warehouse_address_id"`
			IsFromOverseas       bool   `json:"is_from_overseas"`
		} `json:"fulfillment_info"`
		Logistics struct {
			IntegratedChannelids            []int         `json:"integrated_channelids"`
			NonIntegratedChannelids         []interface{} `json:"non_integrated_channelids"`
			VoucherWalletCheckingChannelIds []int         `json:"voucher_wallet_checking_channel_ids"`
			LogisticChannels                struct {
				Num8001 struct {
					ChannelData struct {
						AddressType   int    `json:"address_type"`
						Channelid     int    `json:"channelid"`
						CodSupported  bool   `json:"cod_supported"`
						Enabled       bool   `json:"enabled"`
						IsMaskChannel int    `json:"is_mask_channel"`
						Name          string `json:"name"`
						Priority      int    `json:"priority"`
						Warning       string `json:"warning"`
						WarningMsg    string `json:"warning_msg"`
					} `json:"channel_data"`
					CodData struct {
						CodAvailable bool `json:"cod_available"`
					} `json:"cod_data"`
					DeliveryData struct {
						DelayMessage string `json:"delay_message"`
						DetailInfo   struct {
							Apt      float64 `json:"apt"`
							CdtMax   float64 `json:"cdt_max"`
							CdtMin   float64 `json:"cdt_min"`
							EdtMaxDt string  `json:"edt_max_dt"`
							EdtMinDt string  `json:"edt_min_dt"`
							HeCdt    int     `json:"he_cdt"`
							HePt     int     `json:"he_pt"`
						} `json:"detail_info"`
						DisplayMode               string `json:"display_mode"`
						MaxDays                   int    `json:"max_days"`
						MinDays                   int    `json:"min_days"`
						EstimatedDeliveryTimeMax  int    `json:"estimated_delivery_time_max"`
						EstimatedDeliveryTimeMin  int    `json:"estimated_delivery_time_min"`
						EstimatedDeliveryDateFrom int    `json:"estimated_delivery_date_from"`
						EstimatedDeliveryDateTo   int    `json:"estimated_delivery_date_to"`
						HasEdt                    bool   `json:"has_edt"`
						IsCrossBorder             bool   `json:"is_cross_border"`
						IsRapidSLA                bool   `json:"is_rapid_sla"`
						IsShopee24H               bool   `json:"is_shopee_24h"`
					} `json:"delivery_data"`
					ShippingFeeData struct {
						ChargeableShippingFee     int `json:"chargeable_shipping_fee"`
						ShippingFeeBeforeDiscount int `json:"shipping_fee_before_discount"`
					} `json:"shipping_fee_data"`
				} `json:"8001"`
				Num8003 struct {
					ChannelData struct {
						AddressType   int    `json:"address_type"`
						Channelid     int    `json:"channelid"`
						CodSupported  bool   `json:"cod_supported"`
						Enabled       bool   `json:"enabled"`
						IsMaskChannel int    `json:"is_mask_channel"`
						Name          string `json:"name"`
						Priority      int    `json:"priority"`
						Warning       string `json:"warning"`
						WarningMsg    string `json:"warning_msg"`
					} `json:"channel_data"`
					CodData struct {
						CodAvailable bool `json:"cod_available"`
					} `json:"cod_data"`
					DeliveryData struct {
						DelayMessage string `json:"delay_message"`
						DetailInfo   struct {
							Apt      float64 `json:"apt"`
							CdtMax   float64 `json:"cdt_max"`
							CdtMin   float64 `json:"cdt_min"`
							EdtMaxDt string  `json:"edt_max_dt"`
							EdtMinDt string  `json:"edt_min_dt"`
							HeCdt    int     `json:"he_cdt"`
							HePt     int     `json:"he_pt"`
						} `json:"detail_info"`
						DisplayMode               string `json:"display_mode"`
						MaxDays                   int    `json:"max_days"`
						MinDays                   int    `json:"min_days"`
						EstimatedDeliveryTimeMax  int    `json:"estimated_delivery_time_max"`
						EstimatedDeliveryTimeMin  int    `json:"estimated_delivery_time_min"`
						EstimatedDeliveryDateFrom int    `json:"estimated_delivery_date_from"`
						EstimatedDeliveryDateTo   int    `json:"estimated_delivery_date_to"`
						HasEdt                    bool   `json:"has_edt"`
						IsCrossBorder             bool   `json:"is_cross_border"`
						IsRapidSLA                bool   `json:"is_rapid_sla"`
						IsShopee24H               bool   `json:"is_shopee_24h"`
					} `json:"delivery_data"`
					ShippingFeeData struct {
						ChargeableShippingFee     int `json:"chargeable_shipping_fee"`
						ShippingFeeBeforeDiscount int `json:"shipping_fee_before_discount"`
					} `json:"shipping_fee_data"`
				} `json:"8003"`
				Num8005 struct {
					ChannelData struct {
						AddressType   int    `json:"address_type"`
						Channelid     int    `json:"channelid"`
						CodSupported  bool   `json:"cod_supported"`
						Enabled       bool   `json:"enabled"`
						IsMaskChannel int    `json:"is_mask_channel"`
						Name          string `json:"name"`
						Priority      int    `json:"priority"`
						Warning       string `json:"warning"`
						WarningMsg    string `json:"warning_msg"`
					} `json:"channel_data"`
					CodData struct {
						CodAvailable bool `json:"cod_available"`
					} `json:"cod_data"`
					DeliveryData struct {
						DelayMessage string `json:"delay_message"`
						DetailInfo   struct {
							Apt      float64 `json:"apt"`
							CdtMax   float64 `json:"cdt_max"`
							CdtMin   float64 `json:"cdt_min"`
							EdtMaxDt string  `json:"edt_max_dt"`
							EdtMinDt string  `json:"edt_min_dt"`
							HeCdt    int     `json:"he_cdt"`
							HePt     int     `json:"he_pt"`
						} `json:"detail_info"`
						DisplayMode               string `json:"display_mode"`
						MaxDays                   int    `json:"max_days"`
						MinDays                   int    `json:"min_days"`
						EstimatedDeliveryTimeMax  int    `json:"estimated_delivery_time_max"`
						EstimatedDeliveryTimeMin  int    `json:"estimated_delivery_time_min"`
						EstimatedDeliveryDateFrom int    `json:"estimated_delivery_date_from"`
						EstimatedDeliveryDateTo   int    `json:"estimated_delivery_date_to"`
						HasEdt                    bool   `json:"has_edt"`
						IsCrossBorder             bool   `json:"is_cross_border"`
						IsRapidSLA                bool   `json:"is_rapid_sla"`
						IsShopee24H               bool   `json:"is_shopee_24h"`
					} `json:"delivery_data"`
					ShippingFeeData struct {
						ChargeableShippingFee     int `json:"chargeable_shipping_fee"`
						ShippingFeeBeforeDiscount int `json:"shipping_fee_before_discount"`
					} `json:"shipping_fee_data"`
				} `json:"8005"`
				Num8006 struct {
					ChannelData struct {
						AddressType   int    `json:"address_type"`
						Channelid     int    `json:"channelid"`
						CodSupported  bool   `json:"cod_supported"`
						Enabled       bool   `json:"enabled"`
						IsMaskChannel int    `json:"is_mask_channel"`
						Name          string `json:"name"`
						Priority      int    `json:"priority"`
						Warning       string `json:"warning"`
						WarningMsg    string `json:"warning_msg"`
					} `json:"channel_data"`
					CodData struct {
						CodAvailable bool `json:"cod_available"`
					} `json:"cod_data"`
					DeliveryData struct {
						DelayMessage string `json:"delay_message"`
						DetailInfo   struct {
							Apt      float64 `json:"apt"`
							CdtMax   float64 `json:"cdt_max"`
							CdtMin   float64 `json:"cdt_min"`
							EdtMaxDt string  `json:"edt_max_dt"`
							EdtMinDt string  `json:"edt_min_dt"`
							HeCdt    int     `json:"he_cdt"`
							HePt     int     `json:"he_pt"`
						} `json:"detail_info"`
						DisplayMode               string `json:"display_mode"`
						MaxDays                   int    `json:"max_days"`
						MinDays                   int    `json:"min_days"`
						EstimatedDeliveryTimeMax  int    `json:"estimated_delivery_time_max"`
						EstimatedDeliveryTimeMin  int    `json:"estimated_delivery_time_min"`
						EstimatedDeliveryDateFrom int    `json:"estimated_delivery_date_from"`
						EstimatedDeliveryDateTo   int    `json:"estimated_delivery_date_to"`
						HasEdt                    bool   `json:"has_edt"`
						IsCrossBorder             bool   `json:"is_cross_border"`
						IsRapidSLA                bool   `json:"is_rapid_sla"`
						IsShopee24H               bool   `json:"is_shopee_24h"`
					} `json:"delivery_data"`
					ShippingFeeData struct {
						ChargeableShippingFee     int `json:"chargeable_shipping_fee"`
						ShippingFeeBeforeDiscount int `json:"shipping_fee_before_discount"`
					} `json:"shipping_fee_data"`
				} `json:"8006"`
				Num80021 struct {
					ChannelData struct {
						AddressType   int    `json:"address_type"`
						Channelid     int    `json:"channelid"`
						CodSupported  bool   `json:"cod_supported"`
						Enabled       bool   `json:"enabled"`
						IsMaskChannel int    `json:"is_mask_channel"`
						Name          string `json:"name"`
						Priority      int    `json:"priority"`
						Warning       string `json:"warning"`
						WarningMsg    string `json:"warning_msg"`
					} `json:"channel_data"`
					CodData struct {
						CodAvailable bool `json:"cod_available"`
					} `json:"cod_data"`
					DeliveryData struct {
						DelayMessage string `json:"delay_message"`
						DetailInfo   struct {
							Apt      float64 `json:"apt"`
							CdtMax   float64 `json:"cdt_max"`
							CdtMin   float64 `json:"cdt_min"`
							EdtMaxDt string  `json:"edt_max_dt"`
							EdtMinDt string  `json:"edt_min_dt"`
							HeCdt    int     `json:"he_cdt"`
							HePt     int     `json:"he_pt"`
						} `json:"detail_info"`
						DisplayMode               string `json:"display_mode"`
						MaxDays                   int    `json:"max_days"`
						MinDays                   int    `json:"min_days"`
						EstimatedDeliveryTimeMax  int    `json:"estimated_delivery_time_max"`
						EstimatedDeliveryTimeMin  int    `json:"estimated_delivery_time_min"`
						EstimatedDeliveryDateFrom int    `json:"estimated_delivery_date_from"`
						EstimatedDeliveryDateTo   int    `json:"estimated_delivery_date_to"`
						HasEdt                    bool   `json:"has_edt"`
						IsCrossBorder             bool   `json:"is_cross_border"`
						IsRapidSLA                bool   `json:"is_rapid_sla"`
						IsShopee24H               bool   `json:"is_shopee_24h"`
					} `json:"delivery_data"`
					ShippingFeeData struct {
						ChargeableShippingFee     int `json:"chargeable_shipping_fee"`
						ShippingFeeBeforeDiscount int `json:"shipping_fee_before_discount"`
					} `json:"shipping_fee_data"`
				} `json:"80021"`
			} `json:"logistic_channels"`
			LogisticServiceTypes struct {
				Regular struct {
					ChannelIds []int  `json:"channel_ids"`
					Enabled    bool   `json:"enabled"`
					Identifier string `json:"identifier"`
					MaxCost    int    `json:"max_cost"`
					MinCost    int    `json:"min_cost"`
					Name       string `json:"name"`
					Priority   int    `json:"priority"`
					SLAMsg     string `json:"sla_msg"`
				} `json:"regular"`
				RegularCargo struct {
					ChannelIds []int  `json:"channel_ids"`
					Enabled    bool   `json:"enabled"`
					Identifier string `json:"identifier"`
					MaxCost    int    `json:"max_cost"`
					MinCost    int    `json:"min_cost"`
					Name       string `json:"name"`
					Priority   int    `json:"priority"`
					SLAMsg     string `json:"sla_msg"`
				} `json:"regular_cargo"`
				SameDay struct {
					ChannelIds []int  `json:"channel_ids"`
					Enabled    bool   `json:"enabled"`
					Identifier string `json:"identifier"`
					MaxCost    int    `json:"max_cost"`
					MinCost    int    `json:"min_cost"`
					Name       string `json:"name"`
					Priority   int    `json:"priority"`
					SLAMsg     string `json:"sla_msg"`
				} `json:"same_day"`
				SelfCollection struct {
					ChannelIds []int  `json:"channel_ids"`
					Enabled    bool   `json:"enabled"`
					Identifier string `json:"identifier"`
					MaxCost    int    `json:"max_cost"`
					MinCost    int    `json:"min_cost"`
					Name       string `json:"name"`
					Priority   int    `json:"priority"`
					SLAMsg     string `json:"sla_msg"`
				} `json:"self_collection"`
			} `json:"logistic_service_types"`
		} `json:"logistics"`
		OrderTotal                           int         `json:"order_total"`
		OrderTotalWithoutShipping            int         `json:"order_total_without_shipping"`
		SelectedLogisticChannelidWithWarning interface{} `json:"selected_logistic_channelid_with_warning"`
		ShippingFee                          int         `json:"shipping_fee"`
		ShippingFeeDiscount                  int         `json:"shipping_fee_discount"`
		ShippingGroupDescription             string      `json:"shipping_group_description"`
		ShippingGroupIcon                    string      `json:"shipping_group_icon"`
		TaxPayable                           int         `json:"tax_payable"`
		IsFsvApplied                         bool        `json:"is_fsv_applied"`
	} `json:"shipping_orders"`
	FsvSelectionInfos []interface{} `json:"fsv_selection_infos"`
	BuyerInfo         struct {
		ShareToFriendsInfo struct {
			DisplayToggle bool `json:"display_toggle"`
			EnableToggle  bool `json:"enable_toggle"`
			AllowToShare  bool `json:"allow_to_share"`
		} `json:"share_to_friends_info"`
		KycInfo       interface{} `json:"kyc_info"`
		CheckoutEmail string      `json:"checkout_email"`
	} `json:"buyer_info"`
	PaymentChannelInfo struct {
		Channels []struct {
			Subcategory         int    `json:"subcategory"`
			Icon                string `json:"icon"`
			PaymentResultText   string `json:"payment_result_text"`
			SelectedChannelItem struct {
				ErrorCode       int    `json:"error_code"`
				ErrorMessage    string `json:"error_message"`
				ChannelItemID   int    `json:"channel_item_id"`
				ChannelItemData struct {
					Identifier      string `json:"identifier"`
					BankID          int    `json:"bank_id"`
					CardNumber      string `json:"card_number"`
					ExpiryDate      string `json:"expiry_date"`
					AgencyUen       string `json:"agency_uen"`
					AgencyShortname string `json:"agency_shortname"`
					LogoURL         string `json:"logo_url"`
					BankName        string `json:"bank_name"`
				} `json:"channel_item_data"`
			} `json:"selected_channel_item"`
			DisabledReasonTitle               string `json:"disabled_reason_title"`
			DisabledReasonLearnMoreButtonText string `json:"disabled_reason_learn_more_button_text"`
			ActivationNeeded                  bool   `json:"activation_needed,omitempty"`
			ExtraData                         struct {
				BannedCategories            []interface{} `json:"banned_categories"`
				ExpiryDuration              int           `json:"expiry_duration"`
				AlternateChannelIds         []interface{} `json:"alternate_channel_ids"`
				IncludeShippingAndDiscounts bool          `json:"include_shipping_and_discounts"`
				EligibleTransfer            string        `json:"eligible_transfer"`
				V1ChannelGroupingInfo       struct {
					ChannelIdsToCombine []interface{} `json:"channel_ids_to_combine"`
					ChannelName         string        `json:"channel_name"`
					IconPath            string        `json:"icon_path"`
					IconBackground      string        `json:"icon_background"`
					GroupID             int           `json:"group_id"`
					Version             int           `json:"version"`
				} `json:"v1_channel_grouping_info"`
				MinAmount          int    `json:"min_amount"`
				SubDescription1    string `json:"sub_description_1"`
				SubDescription2    string `json:"sub_description_2"`
				ExpiryExtension    int    `json:"expiry_extension"`
				MaxCheckoutPrice   int    `json:"max_checkout_price"`
				PriceLimit         int    `json:"price_limit"`
				VoucherPaymentType int    `json:"voucher_payment_type"`
			} `json:"extra_data"`
			PinSetUpNeeded  bool `json:"pin_set_up_needed,omitempty"`
			ChannelBehavior struct {
				DisableCancel      bool `json:"disable_cancel"`
				DisableInstruction bool `json:"disable_instruction"`
			} `json:"channel_behavior,omitempty"`
			Priority        int           `json:"priority"`
			InfoText        string        `json:"info_text"`
			ChannelProvider string        `json:"channel_provider"`
			OptionKeysName  []interface{} `json:"option_keys_name"`
			ChannelID       int           `json:"channel_id,omitempty"`
			SpmChannelID    int           `json:"spm_channel_id"`
			PromotionInfo   struct {
				VoucherPaymentType []interface{} `json:"voucher_payment_type"`
				Text               string        `json:"text"`
			} `json:"promotion_info"`
			BeChannelID        int    `json:"be_channel_id"`
			ChannelBlackbox    string `json:"channel_blackbox"`
			InfoLink           string `json:"info_link"`
			SubDescriptionInfo struct {
				First struct {
					Text      string `json:"text"`
					Highlight bool   `json:"highlight"`
				} `json:"first"`
				Second struct {
					Text      string `json:"text"`
					Highlight bool   `json:"highlight"`
				} `json:"second"`
			} `json:"sub_description_info"`
			Version                 int           `json:"version"`
			TopupNeeded             bool          `json:"topup_needed,omitempty"`
			Name                    string        `json:"name"`
			Description             string        `json:"description"`
			Country                 string        `json:"country"`
			Flag                    int64         `json:"flag"`
			IsNew                   bool          `json:"is_new"`
			SpmOptionInfo           string        `json:"spm_option_info"`
			DisabledReason          string        `json:"disabled_reason"`
			DisabledReasonKey       int           `json:"disabled_reason_key"`
			PaymentChannelBanners   []interface{} `json:"payment_channel_banners"`
			Currency                string        `json:"currency"`
			Category                int           `json:"category"`
			Enabled                 bool          `json:"enabled"`
			KycEnforcementNeeded    bool          `json:"kyc_enforcement_needed,omitempty"`
			HasWallet               bool          `json:"has_wallet"`
			AlwaysShow              bool          `json:"always_show"`
			NameLabel               string        `json:"name_label"`
			DisabledReasonLearnMore string        `json:"disabled_reason_learn_more"`
			Balance                 int           `json:"balance,omitempty"`
			WalletBalance           int           `json:"wallet_balance"`
			Groups                  struct {
				Seabank      bool `json:"seabank"`
				Immediate    bool `json:"immediate"`
				BankAccount  bool `json:"bank_account"`
				SpmChannel   bool `json:"spm_channel"`
				GeneralCard  bool `json:"general_card"`
				Cod          bool `json:"cod"`
				BankTransfer bool `json:"bank_transfer"`
				Airpay       bool `json:"airpay"`
				Spl          bool `json:"spl"`
				CreditCard   bool `json:"credit_card"`
				Wallet       bool `json:"wallet"`
				Installment  bool `json:"installment"`
			} `json:"groups"`
			PreselectDisabled bool   `json:"preselect_disabled"`
			IconBackground    string `json:"icon_background"`
			PinResetNeeded    bool   `json:"pin_reset_needed,omitempty"`
			ChannelBehavior2  struct {
			} `json:"channel_behavior,omitempty"`
			Channelid int `json:"channelid,omitempty"`
			Banks     []struct {
				SubDescriptionInfo struct {
					Important string `json:"important"`
					Normal    string `json:"normal"`
				} `json:"sub_description_info"`
				IconLabel             string `json:"icon_label"`
				Hide                  bool   `json:"hide"`
				OptionInfo            string `json:"option_info"`
				Description           string `json:"description"`
				Enabled               bool   `json:"enabled"`
				IconBackground        string `json:"icon_background"`
				DisabledReason        string `json:"disabled_reason"`
				Icon                  string `json:"icon"`
				DeepLink              string `json:"deep_link"`
				HandlingFee           int    `json:"handling_fee"`
				BankName              string `json:"bank_name"`
				PopupConfirmationData struct {
					OkButtonMessage     string `json:"ok_button_message"`
					CancelButtonMessage string `json:"cancel_button_message"`
					NeedPopup           bool   `json:"need_popup"`
					PopupMessage        string `json:"popup_message"`
				} `json:"popup_confirmation_data"`
				BeChannelID int `json:"be_channel_id"`
			} `json:"banks,omitempty"`
			AddCardURL struct {
				API string `json:"api"`
				UI  string `json:"ui"`
			} `json:"add_card_url,omitempty"`
			Cards      []interface{} `json:"cards,omitempty"`
			Promotions []struct {
				URL             string `json:"url"`
				Title           string `json:"title"`
				DiscountType    int    `json:"discount_type"`
				BankID          int    `json:"bank_id"`
				CardPromotionID int    `json:"card_promotion_id"`
				BankLogo        string `json:"bank_logo"`
				StartTime       int    `json:"start_time"`
				Description     string `json:"description"`
				EndTime         int    `json:"end_time"`
				PrimaryColor    string `json:"primary_color"`
				IsFullyRedeemed bool   `json:"is_fully_redeemed"`
				DiscountValue   string `json:"discount_value"`
			} `json:"promotions,omitempty"`
			PaymentChannelHints []interface{} `json:"payment_channel_hints,omitempty"`
			InstallmentBanks    []interface{} `json:"installment_banks,omitempty"`
			AddCardURL2         struct {
				UI string `json:"ui"`
			} `json:"add_card_url,omitempty"`
			AddCardButtonLabel string `json:"add_card_button_label,omitempty"`
			HideAddCardButton  bool   `json:"hide_add_card_button,omitempty"`
			AddCardURL3        struct {
				UI string `json:"ui"`
			} `json:"add_card_url,omitempty"`
			Subchannels []struct {
				KycEnforcementNeeded  bool          `json:"kyc_enforcement_needed"`
				PaymentChannelBanners []interface{} `json:"payment_channel_banners"`
				ParentChannelid       int           `json:"parent_channelid"`
				PaymentChannelHints   []interface{} `json:"payment_channel_hints"`
				Icon                  string        `json:"icon"`
				SubDescriptionInfo    struct {
					Second struct {
						Text      string `json:"text"`
						Highlight bool   `json:"highlight"`
					} `json:"second"`
					First struct {
						Text      string `json:"text"`
						Highlight bool   `json:"highlight"`
					} `json:"first"`
				} `json:"sub_description_info"`
				SkipPlpCheckoutVerification bool          `json:"skip_plp_checkout_verification"`
				Instructions                []interface{} `json:"instructions"`
				InstallmentPromotionLink    string        `json:"installment_promotion_link"`
				InfoLink                    string        `json:"info_link"`
				DisabledReasonTitle         string        `json:"disabled_reason_title"`
				Banks                       []interface{} `json:"banks"`
				PaymentResultText           string        `json:"payment_result_text"`
				DisabledReasonLearnMore     string        `json:"disabled_reason_learn_more"`
				NeedTopupAmount             int           `json:"need_topup_amount"`
				Subchannels                 []interface{} `json:"subchannels"`
				PaymentReasonInfo           struct {
					ReasonMaxLength int  `json:"reason_max_length"`
					NeedReason      bool `json:"need_reason"`
				} `json:"payment_reason_info"`
				Version           int   `json:"version"`
				Priority          int   `json:"priority"`
				Flag              int64 `json:"flag"`
				PaymentDisclaimer struct {
					Enabled bool   `json:"enabled"`
					Preview string `json:"preview"`
					Header  string `json:"header"`
					Body    string `json:"body"`
					URL     string `json:"url"`
				} `json:"payment_disclaimer"`
				ChannelBlackbox string `json:"channel_blackbox"`
				BiInfo          struct {
					IsKycVerified  bool `json:"is_kyc_verified"`
					HasCbSeller    bool `json:"has_cb_seller"`
					HasLocalSeller bool `json:"has_local_seller"`
				} `json:"bi_info"`
				ChannelProvider string `json:"channel_provider"`
				Currency        string `json:"currency"`
				PromotionInfo   struct {
					VoucherPaymentType []interface{} `json:"voucher_payment_type"`
					Text               string        `json:"text"`
				} `json:"promotion_info"`
				SpmItemsKeyName   string `json:"spm_items_key_name"`
				LinkedBankAccount struct {
					ChannelItemID        int    `json:"channel_item_id"`
					Default              bool   `json:"default"`
					DisabledReasonKey    string `json:"disabled_reason_key"`
					Icon                 string `json:"icon"`
					BankName             string `json:"bank_name"`
					BankAccountNumber    string `json:"bank_account_number"`
					LinkBankAccountURL   string `json:"link_bank_account_url"`
					IsExpired            bool   `json:"is_expired"`
					NeedVerification     bool   `json:"need_verification"`
					VerificationMessage  string `json:"verification_message"`
					IsLinked             bool   `json:"is_linked"`
					Enabled              bool   `json:"enabled"`
					LinkBankAccountPopup struct {
						PopupMessage        string `json:"popup_message"`
						OkButtonMessage     string `json:"ok_button_message"`
						CancelButtonMessage string `json:"cancel_button_message"`
						NeedPopup           bool   `json:"need_popup"`
					} `json:"link_bank_account_popup"`
					ConsentTokenDisplayNumber string `json:"consent_token_display_number"`
					DisabledReason            string `json:"disabled_reason"`
					VerificationURL           string `json:"verification_url"`
				} `json:"linked_bank_account"`
				Description                       string        `json:"description"`
				SpmOptionInfo                     string        `json:"spm_option_info"`
				DisabledReasonLearnMoreButtonText string        `json:"disabled_reason_learn_more_button_text"`
				WalletBalance                     int           `json:"wallet_balance"`
				AlwaysExpanded                    bool          `json:"always_expanded"`
				Name                              string        `json:"name"`
				DisabledReasonKey                 int           `json:"disabled_reason_key"`
				Balance                           int           `json:"balance"`
				BeChannelID                       int           `json:"be_channel_id"`
				Cards                             []interface{} `json:"cards"`
				Promotions                        []interface{} `json:"promotions"`
				PinResetNeeded                    bool          `json:"pin_reset_needed"`
				PinSetUpNeeded                    bool          `json:"pin_set_up_needed"`
				AddBankAccountURL                 string        `json:"add_bank_account_url"`
				DirectDebitIntegratedBankList     []interface{} `json:"direct_debit_integrated_bank_list"`
				WhitelistedFields                 []string      `json:"whitelisted_fields"`
				NameLabel                         string        `json:"name_label"`
				ChannelBehavior                   []interface{} `json:"channel_behavior"`
				BankAccounts                      []interface{} `json:"bank_accounts"`
				DefaultBankAccount                struct {
					Version               int `json:"version"`
					ChannelID             int `json:"channel_id"`
					ChannelItemOptionInfo struct {
						ChannelItemID         int    `json:"channel_item_id"`
						ExternalChannelItemID int    `json:"external_channel_item_id"`
						OptionInfo            string `json:"option_info"`
					} `json:"channel_item_option_info"`
				} `json:"default_bank_account"`
				DebitingAgentList []interface{} `json:"debiting_agent_list"`
				Groups            struct {
					Immediate    bool `json:"immediate"`
					Wallet       bool `json:"wallet"`
					BankAccount  bool `json:"bank_account"`
					BankTransfer bool `json:"bank_transfer"`
					Seabank      bool `json:"seabank"`
					Spl          bool `json:"spl"`
					CreditCard   bool `json:"credit_card"`
					Installment  bool `json:"installment"`
					GeneralCard  bool `json:"general_card"`
					SpmChannel   bool `json:"spm_channel"`
					Cod          bool `json:"cod"`
					Airpay       bool `json:"airpay"`
				} `json:"groups"`
				TopupNeeded        bool `json:"topup_needed"`
				PartnerAccountInfo struct {
					IsNewAccount     bool          `json:"is_new_account"`
					ShouldShowButton bool          `json:"should_show_button"`
					DisplayText      string        `json:"display_text"`
					DisplayColor     string        `json:"display_color"`
					PartnersLink     string        `json:"partners_link"`
					TrackingInfo     string        `json:"tracking_info"`
					Balance          int           `json:"balance"`
					ForceGreyOut     bool          `json:"force_grey_out"`
					UserStatus       int           `json:"user_status"`
					AdditionalInfo   int           `json:"additional_info"`
					RulesSameFlag    bool          `json:"rules_same_flag"`
					RulesSkuNames    []interface{} `json:"rules_sku_names"`
				} `json:"partner_account_info"`
				BillingAddresses         []interface{} `json:"billing_addresses"`
				Counters                 []interface{} `json:"counters"`
				AddCardURL               []interface{} `json:"add_card_url"`
				HideAddCardButton        bool          `json:"hide_add_card_button"`
				AddBankAccountText       string        `json:"add_bank_account_text"`
				CcbText                  string        `json:"ccb_text"`
				CcbLink                  string        `json:"ccb_link"`
				AlwaysShow               bool          `json:"always_show"`
				Subcategory              int           `json:"subcategory"`
				Enabled                  bool          `json:"enabled"`
				WalletNoPin              bool          `json:"wallet_no_pin"`
				AddCardButtonLabel       string        `json:"add_card_button_label"`
				OptionKeysName           []interface{} `json:"option_keys_name"`
				Channelid                int           `json:"channelid"`
				InstallmentPromotionText string        `json:"installment_promotion_text"`
				SelectedChannelItem      struct {
					ErrorCode       int    `json:"error_code"`
					ErrorMessage    string `json:"error_message"`
					ChannelItemID   int    `json:"channel_item_id"`
					ChannelItemData struct {
						Identifier      string `json:"identifier"`
						BankName        string `json:"bank_name"`
						AgencyShortname string `json:"agency_shortname"`
						LogoURL         string `json:"logo_url"`
						BankID          int    `json:"bank_id"`
						CardNumber      string `json:"card_number"`
						ExpiryDate      string `json:"expiry_date"`
						AgencyUen       string `json:"agency_uen"`
					} `json:"channel_item_data"`
				} `json:"selected_channel_item"`
				SpmChannelID       int `json:"spm_channel_id"`
				IsReadyToPayConfig struct {
					Request              string        `json:"request"`
					ResponseSuccessCases []interface{} `json:"response_success_cases"`
					DisableText          string        `json:"disable_text"`
				} `json:"is_ready_to_pay_config"`
				PreselectDisabled bool          `json:"preselect_disabled"`
				InfoText          string        `json:"info_text"`
				InstallmentBanks  []interface{} `json:"installment_banks"`
				Country           string        `json:"country"`
				DisabledReason    string        `json:"disabled_reason"`
				ExtraData         struct {
					IncludeShippingAndDiscounts bool   `json:"include_shipping_and_discounts"`
					SubDescription1             string `json:"sub_description_1"`
					VoucherPaymentType          int    `json:"voucher_payment_type"`
					V1ChannelGroupingInfo       struct {
						ChannelIdsToCombine []interface{} `json:"channel_ids_to_combine"`
						ChannelName         string        `json:"channel_name"`
						IconPath            string        `json:"icon_path"`
						IconBackground      string        `json:"icon_background"`
						GroupID             int           `json:"group_id"`
						Version             int           `json:"version"`
					} `json:"v1_channel_grouping_info"`
					BannedCategories    []interface{} `json:"banned_categories"`
					MinAmount           int           `json:"min_amount"`
					ExpiryDuration      int           `json:"expiry_duration"`
					ExpiryExtension     int           `json:"expiry_extension"`
					SubDescription2     string        `json:"sub_description_2"`
					AlternateChannelIds []interface{} `json:"alternate_channel_ids"`
					MaxCheckoutPrice    int64         `json:"max_checkout_price"`
					PriceLimit          int           `json:"price_limit"`
					EligibleTransfer    string        `json:"eligible_transfer"`
				} `json:"extra_data"`
				HasWallet           bool          `json:"has_wallet"`
				ChannelID           int           `json:"channel_id"`
				BlockedByPlp        bool          `json:"blocked_by_plp"`
				IsNew               bool          `json:"is_new"`
				InstallmentPlans    []interface{} `json:"installment_plans"`
				LinkBankAccountInfo struct {
					LinkIcon string `json:"link_icon"`
					LinkText string `json:"link_text"`
				} `json:"link_bank_account_info"`
				Category         int    `json:"category"`
				IconBackground   string `json:"icon_background"`
				ActivationNeeded bool   `json:"activation_needed"`
			} `json:"subchannels,omitempty"`
			ChannelBehavior3 struct {
			} `json:"channel_behavior,omitempty"`
		} `json:"channels"`
		GroupingInfo struct {
			Groups []struct {
				GroupID     int   `json:"group_id"`
				ToCombine   []int `json:"to_combine"`
				DisplayInfo struct {
					PromotionInfo struct {
						Text string `json:"text"`
					} `json:"promotion_info"`
					IsNew          bool   `json:"is_new"`
					IconBackground string `json:"icon_background"`
					InfoLink       string `json:"info_link"`
					Name           string `json:"name"`
					Version        int    `json:"version"`
					ChannelID      int    `json:"channel_id"`
					Icon           string `json:"icon"`
					InfoText       string `json:"info_text"`
				} `json:"display_info"`
			} `json:"groups"`
		} `json:"grouping_info"`
		PromotionInfo struct {
			Text string `json:"text"`
		} `json:"promotion_info"`
	} `json:"payment_channel_info"`
	BuyerTxnFeeInfo struct {
		Title        string `json:"title"`
		Description  string `json:"description"`
		LearnMoreURL string `json:"learn_more_url"`
	} `json:"buyer_txn_fee_info"`
	DisabledCheckoutInfo struct {
		Description string `json:"description"`
		AutoPopup   bool   `json:"auto_popup"`
		ErrorInfos  []struct {
			ErrorAction string      `json:"error_action"`
			ErrorType   string      `json:"error_type"`
			Message     interface{} `json:"message"`
		} `json:"error_infos"`
	} `json:"disabled_checkout_info"`
	CanCheckout bool `json:"can_checkout"`
}

type PayloadGetCheckOut struct {
	Cft        []int `json:"_cft"`
	Timestamp  int   `json:"timestamp"`
	Shoporders []struct {
		Shop struct {
			Shopid int `json:"shopid"`
		} `json:"shop"`
		Items []struct {
			Itemid         int64       `json:"itemid"`
			Modelid        int64       `json:"modelid"`
			Quantity       int         `json:"quantity"`
			AddOnDealID    int         `json:"add_on_deal_id"`
			IsAddOnSubItem bool        `json:"is_add_on_sub_item"`
			ItemGroupID    interface{} `json:"item_group_id"`
			Insurances     []struct {
				InsuranceProductID    string `json:"insurance_product_id"`
				Name                  string `json:"name"`
				Description           string `json:"description"`
				ProductDetailPageURL  string `json:"product_detail_page_url"`
				Premium               int    `json:"premium"`
				PremiumBeforeDiscount int    `json:"premium_before_discount"`
				InsuranceQuantity     int    `json:"insurance_quantity"`
				Selected              bool   `json:"selected"`
			} `json:"insurances"`
		} `json:"items"`
		ShippingID int `json:"shipping_id"`
	} `json:"shoporders"`
	SelectedPaymentChannelData struct {
		ChannelID             int `json:"channel_id"`
		ChannelItemOptionInfo struct {
			OptionInfo string `json:"option_info"`
		} `json:"channel_item_option_info"`
		Version int `json:"version"`
	} `json:"selected_payment_channel_data"`
	PromotionData struct {
		UseCoins                bool `json:"use_coins"`
		FreeShippingVoucherInfo struct {
			FreeShippingVoucherID   int         `json:"free_shipping_voucher_id"`
			FreeShippingVoucherCode string      `json:"free_shipping_voucher_code"`
			DisabledReason          interface{} `json:"disabled_reason"`
			BannerInfo              struct {
				Msg          string `json:"msg"`
				LearnMoreMsg string `json:"learn_more_msg"`
			} `json:"banner_info"`
		} `json:"free_shipping_voucher_info"`
		PlatformVouchers          []interface{} `json:"platform_vouchers"`
		ShopVouchers              []interface{} `json:"shop_vouchers"`
		CheckShopVoucherEntrances bool          `json:"check_shop_voucher_entrances"`
		AutoApplyShopVoucher      bool          `json:"auto_apply_shop_voucher"`
	} `json:"promotion_data"`
	FsvSelectionInfos []interface{} `json:"fsv_selection_infos"`
	DeviceInfo        struct {
		DeviceID          string `json:"device_id"`
		DeviceFingerprint string `json:"device_fingerprint"`
		TongdunBlackbox   string `json:"tongdun_blackbox"`
		BuyerPaymentInfo  struct {
		} `json:"buyer_payment_info"`
	} `json:"device_info"`
	BuyerInfo struct {
		ShareToFriendsInfo struct {
			DisplayToggle bool `json:"display_toggle"`
			EnableToggle  bool `json:"enable_toggle"`
			AllowToShare  bool `json:"allow_to_share"`
		} `json:"share_to_friends_info"`
		KycInfo       interface{} `json:"kyc_info"`
		CheckoutEmail string      `json:"checkout_email"`
	} `json:"buyer_info"`
	CartType int `json:"cart_type"`
	ClientID int `json:"client_id"`
	TaxInfo  struct {
		TaxID string `json:"tax_id"`
	} `json:"tax_info"`
	DropshippingInfo struct {
		Enabled     bool   `json:"enabled"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	} `json:"dropshipping_info"`
	ShippingOrders []struct {
		Sync             bool `json:"sync"`
		BuyerAddressData struct {
			Addressid   int    `json:"addressid"`
			AddressType int    `json:"address_type"`
			TaxAddress  string `json:"tax_address"`
		} `json:"buyer_address_data"`
		SelectedLogisticChannelid             int         `json:"selected_logistic_channelid"`
		ShippingID                            int         `json:"shipping_id"`
		ShoporderIndexes                      []int       `json:"shoporder_indexes"`
		SelectedPreferredDeliveryTimeOptionID int         `json:"selected_preferred_delivery_time_option_id"`
		SelectedPreferredDeliveryTimeSlotID   interface{} `json:"selected_preferred_delivery_time_slot_id"`
	} `json:"shipping_orders"`
	OrderUpdateInfo struct {
	} `json:"order_update_info"`
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
		TotalPayable                    int64       `json:"total_payable"`
	} `json:"checkout_price_data"`
}

type PayloadPlaceHolder struct {
	Status  int `json:"status"`
	Headers struct {
	} `json:"headers"`
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
		TotalPayable                    int64       `json:"total_payable"`
	} `json:"checkout_price_data"`
	OrderUpdateInfo struct {
	} `json:"order_update_info"`
	DropshippingInfo struct {
		Enabled     bool   `json:"enabled"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	} `json:"dropshipping_info"`
	PromotionData struct {
		CanUseCoins             bool          `json:"can_use_coins"`
		UseCoins                bool          `json:"use_coins"`
		PlatformVouchers        []interface{} `json:"platform_vouchers"`
		FreeShippingVoucherInfo struct {
			FreeShippingVoucherID   int         `json:"free_shipping_voucher_id"`
			FreeShippingVoucherCode string      `json:"free_shipping_voucher_code"`
			DisabledReason          interface{} `json:"disabled_reason"`
			BannerInfo              struct {
				Msg          string `json:"msg"`
				LearnMoreMsg string `json:"learn_more_msg"`
			} `json:"banner_info"`
		} `json:"free_shipping_voucher_info"`
		ShopVoucherEntrances []struct {
			Shopid int  `json:"shopid"`
			Status bool `json:"status"`
		} `json:"shop_voucher_entrances"`
		AppliedVoucherCode interface{} `json:"applied_voucher_code"`
		VoucherCode        interface{} `json:"voucher_code"`
		VoucherInfo        struct {
			CoinEarned         int         `json:"coin_earned"`
			VoucherCode        interface{} `json:"voucher_code"`
			CoinPercentage     int         `json:"coin_percentage"`
			DiscountPercentage int         `json:"discount_percentage"`
			DiscountValue      int         `json:"discount_value"`
			Promotionid        int         `json:"promotionid"`
			RewardType         int         `json:"reward_type"`
			UsedPrice          int         `json:"used_price"`
		} `json:"voucher_info"`
		InvalidMessage string `json:"invalid_message"`
		PriceDiscount  int    `json:"price_discount"`
		CoinInfo       struct {
			CoinOffset        int `json:"coin_offset"`
			CoinUsed          int `json:"coin_used"`
			CoinEarnByVoucher int `json:"coin_earn_by_voucher"`
			CoinEarn          int `json:"coin_earn"`
		} `json:"coin_info"`
		CardPromotionID      interface{} `json:"card_promotion_id"`
		CardPromotionEnabled bool        `json:"card_promotion_enabled"`
		PromotionMsg         string      `json:"promotion_msg"`
	} `json:"promotion_data"`
	SelectedPaymentChannelData struct {
		Version               int    `json:"version"`
		OptionInfo            string `json:"option_info"`
		ChannelID             int    `json:"channel_id"`
		ChannelItemOptionInfo struct {
			OptionInfo string `json:"option_info"`
		} `json:"channel_item_option_info"`
		AdditionalInfo struct {
			Reason          string `json:"reason"`
			ChannelBlackbox string `json:"channel_blackbox"`
		} `json:"additional_info"`
		TextInfo struct {
		} `json:"text_info"`
	} `json:"selected_payment_channel_data"`
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
			Itemid      int64       `json:"itemid"`
			Modelid     int64       `json:"modelid"`
			Quantity    int         `json:"quantity"`
			ItemGroupID interface{} `json:"item_group_id"`
			Insurances  []struct {
				InsuranceProductID    string `json:"insurance_product_id"`
				Name                  string `json:"name"`
				Description           string `json:"description"`
				ProductDetailPageURL  string `json:"product_detail_page_url"`
				Premium               int    `json:"premium"`
				PremiumBeforeDiscount int    `json:"premium_before_discount"`
				InsuranceQuantity     int    `json:"insurance_quantity"`
				Selected              bool   `json:"selected"`
			} `json:"insurances"`
			Shopid                  int    `json:"shopid"`
			Shippable               bool   `json:"shippable"`
			NonShippableErr         string `json:"non_shippable_err"`
			NoneShippableReason     string `json:"none_shippable_reason"`
			NoneShippableFullReason string `json:"none_shippable_full_reason"`
			Price                   int    `json:"price"`
			Name                    string `json:"name"`
			ModelName               string `json:"model_name"`
			AddOnDealID             int    `json:"add_on_deal_id"`
			IsAddOnSubItem          bool   `json:"is_add_on_sub_item"`
			IsPreOrder              bool   `json:"is_pre_order"`
			IsStreamingPrice        bool   `json:"is_streaming_price"`
			Image                   string `json:"image"`
			Checkout                bool   `json:"checkout"`
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
		TaxPayable                int           `json:"tax_payable"`
		ShippingID                int           `json:"shipping_id"`
		ShippingFeeDiscount       int           `json:"shipping_fee_discount"`
		ShippingFee               int           `json:"shipping_fee"`
		OrderTotalWithoutShipping int           `json:"order_total_without_shipping"`
		OrderTotal                int64         `json:"order_total"`
		BuyerRemark               string        `json:"buyer_remark"`
		ExtAdInfoMappings         []interface{} `json:"ext_ad_info_mappings"`
	} `json:"shoporders"`
	ShippingOrders []struct {
		ShippingID                            int    `json:"shipping_id"`
		ShoporderIndexes                      []int  `json:"shoporder_indexes"`
		SelectedLogisticChannelid             int    `json:"selected_logistic_channelid"`
		SelectedPreferredDeliveryTimeOptionID int    `json:"selected_preferred_delivery_time_option_id"`
		BuyerRemark                           string `json:"buyer_remark"`
		BuyerAddressData                      struct {
			Addressid   int    `json:"addressid"`
			AddressType int    `json:"address_type"`
			TaxAddress  string `json:"tax_address"`
		} `json:"buyer_address_data"`
		FulfillmentInfo struct {
			FulfillmentFlag      int    `json:"fulfillment_flag"`
			FulfillmentSource    string `json:"fulfillment_source"`
			ManagedBySbs         bool   `json:"managed_by_sbs"`
			OrderFulfillmentType int    `json:"order_fulfillment_type"`
			WarehouseAddressID   int    `json:"warehouse_address_id"`
			IsFromOverseas       bool   `json:"is_from_overseas"`
		} `json:"fulfillment_info"`
		OrderTotal                           int64       `json:"order_total"`
		OrderTotalWithoutShipping            int         `json:"order_total_without_shipping"`
		SelectedLogisticChannelidWithWarning interface{} `json:"selected_logistic_channelid_with_warning"`
		ShippingFee                          int         `json:"shipping_fee"`
		ShippingFeeDiscount                  int         `json:"shipping_fee_discount"`
		ShippingGroupDescription             string      `json:"shipping_group_description"`
		ShippingGroupIcon                    string      `json:"shipping_group_icon"`
		TaxPayable                           int         `json:"tax_payable"`
		IsFsvApplied                         bool        `json:"is_fsv_applied"`
	} `json:"shipping_orders"`
	FsvSelectionInfos []interface{} `json:"fsv_selection_infos"`
	BuyerInfo         struct {
		ShareToFriendsInfo struct {
			DisplayToggle bool `json:"display_toggle"`
			EnableToggle  bool `json:"enable_toggle"`
			AllowToShare  bool `json:"allow_to_share"`
		} `json:"share_to_friends_info"`
		KycInfo       interface{} `json:"kyc_info"`
		CheckoutEmail string      `json:"checkout_email"`
	} `json:"buyer_info"`
	BuyerTxnFeeInfo struct {
		Title        string `json:"title"`
		Description  string `json:"description"`
		LearnMoreURL string `json:"learn_more_url"`
	} `json:"buyer_txn_fee_info"`
	DisabledCheckoutInfo struct {
		Description string        `json:"description"`
		AutoPopup   bool          `json:"auto_popup"`
		ErrorInfos  []interface{} `json:"error_infos"`
	} `json:"disabled_checkout_info"`
	CanCheckout bool  `json:"can_checkout"`
	Cft         []int `json:"_cft"`
	DeviceInfo  struct {
		DeviceSzFingerprint string `json:"device_sz_fingerprint"`
	} `json:"device_info"`
	CaptchaVersion int `json:"captcha_version"`
}
