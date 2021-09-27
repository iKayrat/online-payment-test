package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/fatih/structs"
	"github.com/google/uuid"
)

const (
	checkoutUrl      = "https://pay.fondy.eu/api/checkout/url/"
	merchantId       = "1396424"
	merchantPassword = "test"
	currencyUSD      = "USD"
	currencyEUR      = "EUR"
	languageRU       = "ru"
)

type APIRequest struct {
	Requset interface{} `json:"request"`
}

type APIResponse struct {
	Response interface{} `json:"response"`
}

type CheckouRequest struct {
	OrderID          string `json:"order_id"`
	MerchantId       string `json:"merchant_id"`
	OrderDesc        string `json:"orderDesc"`
	Signature        string `json:"signature"`
	Amount           string `json:"amount"`
	Currency         string `json:"currency"`
	ResponseUrl      string `json:"response_url"`
	ServeCallbackUrl string `json:"serve_callback_url"`
	SenderEmail      string `json:"sender_email"`
	Language         string `json:"language"`
	ProductID        string `json:"produc_id"`
}

type InterimResponse struct {
	Status      string `json:"response_status"`
	CheckoutUrl string `json:"checkout_url"`
	PaymentId   string `json:"payment_id"`
}

type CheckouResponse struct {
	OrderID        string `json:"order_id"`
	MerchantId     string `json:"merchant_id"`
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	Signature      string `json:"signature"`
	OrderStatus    string `json:"order_status"`
	ResponseStatus string `json:"response_status"`
	MaskedCard     string `json:"masked_card"`
	CardType       string `json:"card_type"`
	Fee            string `json:"fee"`
	PaymentSystem  string `json:"payment_system"`
	ProductID      string `json:"produc_id"`
	AdditionalInfo string `json:"additional_info"`
}

func (c *CheckouRequest) SetSignature(password string) {
	params := structs.Map(c)
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	values := []string{}
	for _, key := range keys {
		value := params[key].(string)
		if value == "" {
			continue
		}
		values = append(values, value)
	}

	c.Signature = generateSignature(values, password)
}

func generateSignature(values []string, password string) string {
	newValues := []string{password}
	newValues = append(newValues, values...)

	signatureString := strings.Join(newValues, "|")

	fmt.Println(signatureString)

	hash := sha1.New()
	hash.Write([]byte(signatureString))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func main() {
	id := uuid.New()
	checkouReq := &CheckouRequest{
		OrderID:    id.String(),
		MerchantId: merchantId,
		// tourism tickets
		OrderDesc: "Expedition Tour",
		Amount:    "5000",
		Currency:  currencyEUR,
		// https://
		ServeCallbackUrl: "",
	}

	checkouReq.SetSignature(merchantPassword)

	request := APIRequest{Requset: checkouReq}
	requestBody, err := json.Marshal(request)
	if err != nil {
		fmt.Println("marshal err: ", err)
	}

	resp, err := http.Post(checkoutUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	apiResp := APIResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		panic(err)
	}
	fmt.Println(apiResp)
}
