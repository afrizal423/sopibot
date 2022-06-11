package pengguna

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// respone waktu login
type ResponseLogin struct {
	IsInstallmentPaymentWhitelist   bool        `json:"is_installment_payment_whitelist"`
	Cover                           string      `json:"cover"`
	IsCcInstallmentPaymentEligible  bool        `json:"is_cc_installment_payment_eligible"`
	Shopid                          int         `json:"shopid"`
	HasPassword                     bool        `json:"has_password"`
	PaymentPassword                 interface{} `json:"payment_password"`
	Portrait                        string      `json:"portrait"`
	IsSeller                        bool        `json:"is_seller"`
	IsCcInstallmentPaymentWhitelist bool        `json:"is_cc_installment_payment_whitelist"`
	DisallowDataProcessing          interface{} `json:"disallow_data_processing"`
	Access                          struct {
		WalletSetting   int  `json:"wallet_setting"`
		HasLegacyWallet bool `json:"has_legacy_wallet"`
		WalletProvider  int  `json:"wallet_provider"`
	} `json:"access"`
	RecentCartItems []interface{} `json:"recent_cart_items"`
	DefaultAddress  struct {
		LogisticsStatus  int    `json:"logistics_status"`
		Mtime            int    `json:"mtime"`
		Icno             string `json:"icno"`
		ID               int    `json:"id"`
		City             string `json:"city"`
		District         string `json:"district"`
		Zipcode          string `json:"zipcode"`
		State            string `json:"state"`
		GeoString        string `json:"geoString"`
		Status           int    `json:"status"`
		Deftime          int    `json:"deftime"`
		FullAddress      string `json:"full_address"`
		Phone            string `json:"phone"`
		FormattedAddress string `json:"formattedAddress"`
		Address          string `json:"address"`
		Extinfo          struct {
			PlaceID             string `json:"place_id"`
			DeliveryInstruction struct {
			} `json:"delivery_instruction"`
			LabelID                 int    `json:"label_id"`
			PreferredDeliveryOption int    `json:"preferred_delivery_option"`
			ClientID                int    `json:"client_id"`
			Geoinfo                 string `json:"geoinfo"`
		} `json:"extinfo"`
		Ctime   int    `json:"ctime"`
		Town    string `json:"town"`
		Name    string `json:"name"`
		Country string `json:"country"`
		Userid  int    `json:"userid"`
		Geoinfo struct {
			Country string `json:"country"`
			Region  struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"region"`
			Admin1           string `json:"admin1"`
			Admin3           string `json:"admin3"`
			Admin2           string `json:"admin2"`
			FormattedAddress string `json:"formattedAddress"`
			Admin4           string `json:"admin4"`
			PostalCode       string `json:"postalCode"`
		} `json:"geoinfo"`
	} `json:"default_address"`
	FeedAccountInfo struct {
		CanPostFeed bool `json:"can_post_feed"`
	} `json:"feed_account_info"`
	RecentCartItemDetails              []interface{} `json:"recent_cart_item_details"`
	CbOption                           int           `json:"cb_option"`
	Email                              string        `json:"email"`
	Username                           string        `json:"username"`
	TosAcceptedTime                    interface{}   `json:"tos_accepted_time"`
	HolidayModeMtime                   interface{}   `json:"holiday_mode_mtime"`
	Phone                              string        `json:"phone"`
	CreditCardChannel                  int           `json:"credit_card_channel"`
	IsNonCcInstallmentPaymentEligible  bool          `json:"is_non_cc_installment_payment_eligible"`
	PhoneVerified                      bool          `json:"phone_verified"`
	WebOption                          interface{}   `json:"web_option"`
	HolidayModeOn                      interface{}   `json:"holiday_mode_on"`
	Ctime                              int           `json:"ctime"`
	BirthTimestamp                     int           `json:"birth_timestamp"`
	AdultConsent                       interface{}   `json:"adult_consent"`
	IsNonCcInstallmentPaymentWhitelist bool          `json:"is_non_cc_installment_payment_whitelist"`
	Userid                             int           `json:"userid"`
	CookiesAcceptedTime                interface{}   `json:"cookies_accepted_time"`
	NotNewUser                         bool          `json:"not_new_user"`
	CartItemCount                      int           `json:"cart_item_count"`
}

type Address struct {
	Address           string
	City              string
	Country           string
	District          string
	Formatted_address string
	Full_address      string
	Geo_string        string
	Id                int
	Name              string
	Phone             string
	State             string
	Town              string
	Zipcode           string
}

type User struct {
	Userid          int
	Shopid          int
	Name            string
	Email           string
	Phone           string
	Phone_verified  bool
	Default_address Address
	Cookie          string
	Csrf_token      string
}

type AlamatCheckout struct {
	Error int `json:"error"`
	Data  struct {
		Addresses []struct {
			ID                 int    `json:"id"`
			Userid             int    `json:"userid"`
			Name               string `json:"name"`
			Phone              string `json:"phone"`
			Country            string `json:"country"`
			State              string `json:"state"`
			City               string `json:"city"`
			Address            string `json:"address"`
			Status             int    `json:"status"`
			Ctime              int    `json:"ctime"`
			Mtime              int    `json:"mtime"`
			Zipcode            string `json:"zipcode"`
			Deftime            int    `json:"deftime"`
			FullAddress        string `json:"full_address"`
			District           string `json:"district"`
			Town               string `json:"town"`
			LogisticsStatus    int    `json:"logistics_status"`
			Icno               string `json:"icno"`
			Label              string `json:"label"`
			AddressInstruction string `json:"address_instruction"`
			Geoinfo            struct {
				Region struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"region"`
			} `json:"geoinfo"`
			LabelID           int  `json:"label_id"`
			IsPickupAddress   bool `json:"is_pickup_address"`
			IsDeliveryAddress bool `json:"is_delivery_address"`
		} `json:"addresses"`
	} `json:"data"`
}

func Login() User {
	var csrf_token string

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://shopee.co.id/api/v1/account_info", nil)
	if err != nil {
		fmt.Println("No response from request")
	}
	req.Header = http.Header{
		"Accept":          []string{"/"},
		"Accept-Encoding": []string{"zip, deflate, br"},
		"Referer":         []string{"https://shopee.co.id/"},
		"User-Agent":      []string{"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1"},
		"Cookie":          []string{os.Getenv("COOKIE")},
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("No response from request")
	}
	body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))
	var result ResponseLogin
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	// fmt.Println(prettyPrint(result))
	for _, cookie1 := range strings.Split(os.Getenv("COOKIE"), ";") {
		key_value := strings.Split(cookie1, "=")
		if strings.TrimSpace(key_value[0]) == "csrftoken" {
			csrf_token = key_value[1]
		}
	}
	// fmt.Print("csrfnya=", csrf_token)
	u := User{
		Userid:         result.Userid,
		Shopid:         result.Shopid,
		Name:           result.Username,
		Email:          result.Email,
		Phone:          result.Phone,
		Phone_verified: result.PhoneVerified,
		Default_address: Address{
			Address:           result.DefaultAddress.Address,
			City:              result.DefaultAddress.City,
			Country:           result.DefaultAddress.Country,
			District:          result.DefaultAddress.District,
			Formatted_address: result.DefaultAddress.FormattedAddress,
			Full_address:      result.DefaultAddress.FullAddress,
			Geo_string:        result.DefaultAddress.GeoString,
			Id:                result.DefaultAddress.ID,
			Name:              result.DefaultAddress.Name,
			Phone:             result.DefaultAddress.Phone,
			State:             result.DefaultAddress.State,
			Town:              result.DefaultAddress.Town,
			Zipcode:           result.DefaultAddress.Zipcode,
		},
		Cookie:     os.Getenv("COOKIE"),
		Csrf_token: csrf_token,
	}
	// fmt.Println(u)

	return u
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
